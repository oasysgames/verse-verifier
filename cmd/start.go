package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/beacon"
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/collector"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/debug"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/metrics"
	"github.com/oasysgames/oasys-optimism-verifier/p2p"
	"github.com/oasysgames/oasys-optimism-verifier/submitter"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verifier"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/oasysgames/oasys-optimism-verifier/version"
	"github.com/oasysgames/oasys-optimism-verifier/wallet"
	"github.com/spf13/cobra"
)

const (
	StakeManagerAddress = "0x0000000000000000000000000000000000001001"
	SCCName             = "StateCommitmentChain"
	L2OOName            = "L2OutputOracle"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Verifier",
	Long:  "Start the Verifier",
	Run:   runStartCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runStartCmd(cmd *cobra.Command, args []string) {
	log.Info(fmt.Sprintf("Start %s", commandName), "version", version.SemVer())

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := mustNewServer(cmd)

	// start metrics server
	s.mustStartMetrics(ctx)

	// start pprof server
	s.mustStartPprof(ctx)

	// start the ipc server and services dependent on ipc
	s.mustStartIPC(ctx, []func(context.Context, *ipc.IPCServer){
		// unlock walelts(wait forever)
		s.mustUnlockWallets,
		// start p2p (Note: must start the P2P before setup beacon worker)
		s.mustStartP2P,
	})

	// setup workers
	s.setupCollector()
	s.mustSetupVerifier()
	s.setupSubmitter()
	s.setupBeacon()

	// start cache updater
	s.smcache.Refresh(ctx) // first time synchronous
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.smcache.RefreshLoop(ctx, time.Hour)
	}()

	// start workers
	s.startCollector(ctx)
	s.startVerifier(ctx)
	s.startSubmitter(ctx)
	s.startVerseDiscovery(ctx)
	s.startBeacon(ctx)
	log.Info("Start all workers")

	// wait for signal
	s.wg.Wait()
	log.Info("Stopped all workers")
}

type server struct {
	wg               sync.WaitGroup
	conf             *config.Config
	db               *database.Database
	ks               *wallet.KeyStore
	hub              ethutil.Client
	smcache          *stakemanager.Cache
	p2p              *p2p.Node
	blockCollector   *collector.BlockCollector
	eventCollector   *collector.EventCollector
	verifier         *verifier.Verifier
	submitter        *submitter.Submitter
	bw               *beacon.BeaconWorker
	discoveredVerses chan []*config.Verse
}

func mustNewServer(cmd *cobra.Command) *server {
	var err error

	s := &server{discoveredVerses: make(chan []*config.Verse)}

	// load configuration file
	if s.conf, err = loadConfig(cmd); err != nil {
		log.Crit("Failed to load configuration file", "err", err)
	}

	// setup database
	if s.conf.Database.Path == "" {
		s.conf.Database.Path = s.conf.DatabasePath()
	}
	if s.db, err = database.NewDatabase(&s.conf.Database); err != nil {
		log.Crit("Failed to open database", "err", err)
	}

	// open geth keystore
	s.ks = wallet.NewKeyStore(s.conf.KeyStore)

	// construct hub-layer client
	if s.hub, err = ethutil.NewClient(s.conf.HubLayer.RPC); err != nil {
		log.Crit("Failed to construct hub-layer client", "err", err)
	}

	// construct stakemanager cache
	sm, err := stakemanager.NewStakemanagerCaller(
		common.HexToAddress(StakeManagerAddress), s.hub)
	if err != nil {
		log.Crit("Failed to construct StakeManager", "err", err)
	}
	s.smcache = stakemanager.NewCache(sm)

	return s
}

func (s *server) mustStartMetrics(ctx context.Context) {
	if !s.conf.Metrics.Enable {
		return
	}

	metrics.Initialize(&s.conf.Metrics)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := metrics.ListenAndServe(ctx); err != nil {
			log.Crit("Failed to start metrics server", "err", err)
		}
	}()
}

func (s *server) mustStartPprof(ctx context.Context) {
	if !s.conf.Debug.Pprof.Enable {
		return
	}

	ps := debug.NewPprofServer(&s.conf.Debug.Pprof)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := ps.ListenAndServe(ctx); err != nil {
			log.Crit("Failed to start pprof server", "err", err)
		}
	}()
}

func (s *server) mustStartIPC(ctx context.Context, depends []func(context.Context, *ipc.IPCServer)) {
	if s.conf.IPC.Sockname == "" {
		log.Crit("IPC socket name is required")
	}

	ipc, err := ipc.NewIPCServer(s.conf.IPC.Sockname)
	if err != nil {
		log.Crit("Failed to start ipc server", "err", err)
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		ipc.Start(ctx)
	}()

	for _, dep := range depends {
		dep(ctx, ipc)
	}
}

func (s *server) mustStartP2P(ctx context.Context, ipc *ipc.IPCServer) {
	// get p2p private key
	p2pKey, err := getOrCreateP2PKey(s.conf.P2PKeyPath())
	if err != nil {
		log.Crit("Failed to get(or create) p2p key", "err", err)
	}

	// construct libp2p host
	host, dht, bwm, hpHelper, err := p2p.NewHost(ctx, &s.conf.P2P, p2pKey)
	if err != nil {
		log.Crit("Failed to construct libp2p host", "err", err)
	}

	// ignore self-signed signatures
	ignoreSigners := []common.Address{}
	if s.conf.Verifier.Enable {
		_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
		ignoreSigners = append(ignoreSigners, account.Address)
	}

	s.p2p, err = p2p.NewNode(&s.conf.P2P, s.db, host, dht, bwm,
		hpHelper, s.conf.HubLayer.ChainId, ignoreSigners, s.smcache)
	if err != nil {
		log.Crit("Failed to construct p2p node", "err", err)
	}

	ipc.SetHandler(ipccmd.PingCmd.NewHandler(ctx, s.p2p.Host(), s.p2p.HolePunchHelper()))
	ipc.SetHandler(ipccmd.StatusCmd.NewHandler(s.p2p.Host()))

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.p2p.Start(ctx)
	}()
}

func (s *server) setupCollector() {
	if !s.conf.Verifier.Enable {
		return
	}

	s.blockCollector = collector.NewBlockCollector(&s.conf.Verifier, s.db, s.hub)

	_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
	s.eventCollector = collector.NewEventCollector(
		&s.conf.Verifier, s.db, s.hub, account.Address)
}

func (s *server) startCollector(ctx context.Context) {
	// start block collector
	if s.blockCollector != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.blockCollector.Start(ctx)
		}()
	}

	// start event collector
	if s.eventCollector != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.eventCollector.Start(ctx)
		}()
	}
}

func (s *server) mustSetupVerifier() {
	if !s.conf.Verifier.Enable {
		return
	}

	wallet, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
	l1Signer, err := ethutil.NewSignableClient(
		new(big.Int).SetUint64(s.conf.HubLayer.ChainId), s.conf.HubLayer.RPC, wallet, account)
	if err != nil {
		log.Crit("Failed to construct verifier", "err", err)
	}

	s.verifier = verifier.NewVerifier(&s.conf.Verifier, s.db, l1Signer)
}

func (s *server) startVerifier(ctx context.Context) {
	if s.verifier == nil {
		return
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.verifier.Start(ctx)
	}()

	// start database optimizer
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		// optimize database every hour
		tick := util.NewTicker(s.conf.Verifier.OptimizeInterval, 1)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				s.db.OPSignature.RepairPreviousID(s.verifier.L1Signer().Signer())
			}
		}
	}()

	// publish new signature via p2p
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		sub := s.verifier.SubscribeNewSignature(ctx)
		defer sub.Cancel()

		debounce := time.NewTicker(time.Second * 5)
		defer debounce.Stop()

		subscribes := map[common.Address]*database.OptimismSignature{}
		for {
			select {
			case <-ctx.Done():
				return
			case sig := <-sub.Next():
				subscribes[sig.Signer.Address] = sig
			case <-debounce.C:
				if len(subscribes) == 0 {
					continue
				}
				var publishes []*database.OptimismSignature
				for _, sig := range subscribes {
					publishes = append(publishes, sig)
				}
				s.p2p.PublishSignatures(ctx, publishes)
				subscribes = map[common.Address]*database.OptimismSignature{}
			}
		}
	}()
}

func (s *server) setupSubmitter() {
	if !s.conf.Submitter.Enable {
		return
	}

	s.submitter = submitter.NewSubmitter(&s.conf.Submitter, s.db, s.smcache)
}

func (s *server) startSubmitter(ctx context.Context) {
	if s.submitter == nil {
		return
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.submitter.Start(ctx)
	}()
}

func (s *server) startVerseDiscovery(ctx context.Context) {
	// run discovery handler
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		// read verses from the configuration
		go func() {
			s.discoveredVerses <- s.conf.VerseLayer.Directs
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case verses := <-s.discoveredVerses:
				s.verseDiscoveryHandler(verses)
			}
		}
	}()

	// start dynamic discovery
	if s.conf.VerseLayer.Discovery.Endpoint == "" {
		return
	}

	disc := config.NewVerseDiscovery(
		http.DefaultClient,
		s.conf.VerseLayer.Discovery.Endpoint,
		s.conf.VerseLayer.Discovery.RefreshInterval)

	// run worker
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		time.Sleep(time.Second)
		disc.Start(ctx)
	}()

	// publish subscribed verses to verifier and submitter
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		sub := disc.Subscribe(ctx)
		defer sub.Cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case s.discoveredVerses <- <-sub.Next():
			}
		}
	}()
}

func (s *server) verseDiscoveryHandler(discovers []*config.Verse) {
	if s.verifier == nil && s.submitter == nil {
		log.Warn("Both Verifier and Submitter are disabled")
		return
	}

	verseFactories := map[string]verse.VerseFactory{
		SCCName:  verse.NewOPLegacy,
		L2OOName: verse.NewOPStack,
	}
	verifyContracts := map[string]common.Address{
		SCCName:  common.HexToAddress(s.conf.Submitter.SCCVerifierAddress),
		L2OOName: common.HexToAddress(s.conf.Submitter.L2OOVerifierAddress),
	}

	type verse_ struct {
		cfg    *config.Verse
		verse  verse.Verse
		verify common.Address
	}
	var verses []*verse_
	for _, cfg := range discovers {
		for name, addr := range cfg.L1Contracts {
			if factory, ok := verseFactories[name]; ok {
				verses = append(verses, &verse_{
					cfg:    cfg,
					verse:  factory(s.db, s.hub, common.HexToAddress(addr)),
					verify: verifyContracts[name],
				})
			}
		}
	}

	for _, x := range verses {
		// add verse to Verifier
		if s.verifier != nil && !s.verifier.HasTask(x.verse.RollupContract(), x.cfg.RPC) {
			l2Client, err := ethutil.NewClient(x.cfg.RPC)
			if err != nil {
				log.Error("Failed to construct verse-layer client", "err", err)
			} else {
				s.verifier.AddTask(x.verse.WithVerifiable(l2Client))
			}
		}

		// add verse to Submitter
		if s.submitter != nil {
			for _, tg := range s.conf.Submitter.Targets {
				if tg.ChainID != x.cfg.ChainID || s.submitter.HasTask(x.verse.RollupContract()) {
					continue
				}

				wallet, account := findWallet(s.conf, s.ks, tg.Wallet)
				l1Signer, err := ethutil.NewSignableClient(
					new(big.Int).SetUint64(s.conf.HubLayer.ChainId),
					s.conf.HubLayer.RPC, wallet, account)
				if err != nil {
					log.Error("Failed to construct hub-layer client", "err", err)
				} else {
					s.submitter.AddTask(x.verse.WithTransactable(l1Signer, x.verify))
				}
			}
		}
	}
}

func (s *server) setupBeacon() {
	if !s.conf.Beacon.Enable || !s.conf.Verifier.Enable {
		return
	}

	_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
	s.bw = beacon.NewBeaconWorker(
		&s.conf.Beacon,
		http.DefaultClient,
		beacon.Beacon{
			Signer:  account.Address.Hex(),
			Version: version.SemVer(),
			PeerID:  s.p2p.PeerID().String(),
		},
	)
}

func (s *server) startBeacon(ctx context.Context) {
	if s.bw == nil {
		return
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.bw.Start(ctx)
	}()
}

func getOrCreateP2PKey(filename string) (crypto.PrivKey, error) {
	data, err := ioutil.ReadFile(filename)

	if err == nil {
		dec, peerID, err := p2p.DecodePrivateKey(string(data))
		if err != nil {
			log.Error("Failed to decode p2p private key", "err", err)
			return nil, err
		}

		log.Info("Loaded p2p private key", "file", filename, "id", peerID)
		return dec, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		log.Error("Failed to load p2p private key", "err", err)
		return nil, err
	}

	priv, _, peerID, err := p2p.GenerateKeyPair()
	if err != nil {
		log.Error("Failed to generate p2p private key", "err", err)
		return nil, err
	}

	enc, err := p2p.EncodePrivateKey(priv)
	if err != nil {
		log.Error("Failed to encode p2p private key", "err", err)
		return nil, err
	}

	err = ioutil.WriteFile(filename, []byte(enc), 0644)
	if err != nil {
		log.Error("Failed to write p2p private key", "err", err)
		return nil, err
	}

	log.Info("Generated and saved to p2p private key", "file", filename, "id", peerID)
	return priv, nil
}

func (s *server) mustUnlockWallets(ctx context.Context, ipc *ipc.IPCServer) {
	ipc.SetHandler(ipccmd.WalletUnlockCmd.NewHandler(s.ks))

	var wg sync.WaitGroup
	wg.Add(len(s.conf.Wallets))

	for name, wallet := range s.conf.Wallets {
		go func(name string, wallet config.Wallet) {
			defer wg.Done()

			address := common.HexToAddress(wallet.Address)
			_wallet, account, err := s.ks.FindWallet(address)
			if err != nil {
				log.Crit("Failed to find a wallet",
					"name", name, "address", wallet.Address, "err", err)
			}

			if wallet.Password != "" {
				pw, err := ioutil.ReadFile(wallet.Password)
				if err != nil {
					log.Crit("Failed to read password file",
						"name", name, "address", address, "err", err)
				}

				if err := s.ks.Unlock(*account, strings.Trim(string(pw), "\r\n\t ")); err != nil {
					log.Crit("Failed to unlock wallet using password file",
						"name", name, "address", address, "err", err)
				}
				log.Info("Wallet unlocked using password file", "name", name, "address", address)
			} else if s.ks.Unlock(*account, "") == nil {
				log.Info("Wallet unlocked", "name", name, "address", address)
			} else {
				log.Info("Waiting for wallet unlock via IPC", "name", name, "address", address)
				if err := s.ks.WaitForUnlock(ctx, _wallet); err != nil {
					log.Crit("Wallet was not unlocked",
						"name", name, "address", address, "err", err)
				}
				log.Info("Wallet unlocked via IPC", "name", name, "address", address)
			}
		}(name, wallet)
	}

	wg.Wait()
}

func findWallet(
	c *config.Config,
	ks *wallet.KeyStore,
	name string,
) (accounts.Wallet, *accounts.Account) {
	wallet, account, err := ks.FindWallet(common.HexToAddress(c.Wallets[name].Address))
	if err != nil {
		log.Crit("Wallet not found", "name", name)
	}
	return wallet, account
}
