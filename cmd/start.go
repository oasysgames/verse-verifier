package cmd

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer cancel()
		sig := <-sigC
		log.Info("Received signal, stopping...", "signal", sig)
	}()

	s := mustNewServer(ctx)

	// start metrics server
	s.mustStartMetrics(ctx)

	// start pprof server
	s.mustStartPprof(ctx)

	// start the ipc server and services dependent on ipc
	s.mustStartIPC(ctx, []func(context.Context, *ipc.IPCServer){
		// unlock walelts(wait forever)
		s.mustLoadSigners,
		// start p2p (Note: must start the P2P before setup beacon worker)
		s.mustStartP2P,
	})

	// setup workers
	s.mustSetupCollector()
	s.mustSetupVerifier()
	s.setupSubmitter()
	s.mustSetupBeacon()

	// start cache updater
	go func() {
		// NOTE: Don't add wait group, as no need to guarantee the completion
		s.smcache.RefreshLoop(ctx, time.Hour)
	}()

	// start workers
	s.startCollector(ctx)
	s.startVerifier(ctx)
	s.startSubmitter(ctx)
	s.startVerseDiscovery(ctx)
	s.startBeacon(ctx)
	log.Info("All workers started")

	// wait for signal
	<-ctx.Done()
	log.Info("Shutting down all workers")

	// Shutdown metrics server
	if s.msvr != nil {
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.msvr.Shutdown(c)
	}
	// Shutdown pprof server
	if s.psvr != nil {
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.psvr.Shutdown(c)
	}
	// Shutdown ipc server
	if s.ipc != nil {
		s.ipc.Close()
	}

	var (
		// time limit until all worker stop
		limit       = time.Second * 60
		wc, wcancel = context.WithTimeout(context.Background(), limit)
		isTimeout   = true
	)
	go func() {
		s.wg.Wait()
		isTimeout = false
		wcancel()
	}()

	// all worker stopped or timeout
	<-wc.Done()
	if isTimeout {
		log.Crit("Worker stopping time limit has elapsed", "limit", limit)
	}
	log.Info("All workers stopped")
}

type server struct {
	wg               sync.WaitGroup
	conf             *config.Config
	db               *database.Database
	signers          map[string]ethutil.Signer
	hub              ethutil.Client
	smcache          *stakemanager.Cache
	p2p              *p2p.Node
	blockCollector   *collector.BlockCollector
	eventCollector   *collector.EventCollector
	verifier         *verifier.Verifier
	submitter        *submitter.Submitter
	bw               *beacon.BeaconWorker
	discoveredVerses chan []*config.Verse
	msvr             *http.Server
	psvr             *http.Server
	ipc              *ipc.IPCServer
}

func mustNewServer(ctx context.Context) *server {
	var err error

	s := &server{
		signers:          map[string]ethutil.Signer{},
		discoveredVerses: make(chan []*config.Verse),
	}

	if s.conf, err = globalConfigLoader.load(); err != nil {
		log.Crit("Failed to load configuration", "err", err)
	}

	// setup database
	if s.conf.Database.Path == "" {
		s.conf.Database.Path = s.conf.DatabasePath()
	}
	if s.db, err = database.NewDatabase(&s.conf.Database); err != nil {
		log.Crit("Failed to open database", "err", err)
	}

	// construct hub-layer client
	if s.hub, err = ethutil.NewClient(s.conf.HubLayer.RPC); err != nil {
		log.Crit("Failed to construct hub-layer client", "err", err)
	}

	// Make sue the s.hub can connect to the chain
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if _, err := s.hub.BlockNumber(ctx); err != nil {
		log.Crit("Failed to connect to the hub-layer chain", "err", err)
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

	s.msvr = metrics.Initialize(&s.conf.Metrics)
	go func() {
		// NOTE: Don't add wait group, as no need to guarantee the completion
		metrics.ListenAndServe(ctx, s.msvr)
	}()
}

func (s *server) mustStartPprof(ctx context.Context) {
	if !s.conf.Debug.Pprof.Enable {
		return
	}

	var ps *debug.PprofServer
	ps, s.psvr = debug.NewPprofServer(&s.conf.Debug.Pprof)

	go func() {
		// NOTE: Don't add wait group, as no need to guarantee the completion
		ps.ListenAndServe(ctx, s.psvr)
	}()
}

func (s *server) mustStartIPC(ctx context.Context, depends []func(context.Context, *ipc.IPCServer)) {
	if s.conf.IPC.Sockname == "" {
		log.Crit("IPC socket name is required")
	}

	var err error
	if s.ipc, err = ipc.NewIPCServer(s.conf.IPC.Sockname); err != nil {
		log.Crit("Failed to start ipc server", "err", err)
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.ipc.Start()
	}()

	for _, dep := range depends {
		dep(ctx, s.ipc)
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
	if signer, ok := s.signers[s.conf.Verifier.Wallet]; ok {
		ignoreSigners = append(ignoreSigners, signer.From())
	}

	s.p2p, err = p2p.NewNode(&s.conf.P2P, s.db, host, dht, bwm,
		hpHelper, s.conf.HubLayer.ChainID, ignoreSigners, s.smcache)
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

func (s *server) mustSetupCollector() {
	if !s.conf.Verifier.Enable {
		return
	}

	signer, ok := s.signers[s.conf.Verifier.Wallet]
	if !ok {
		log.Crit("Wallet for the Verifier not found", "wallet", s.conf.Verifier.Wallet)
	}

	s.blockCollector = collector.NewBlockCollector(&s.conf.Verifier, s.db, s.hub)
	s.eventCollector = collector.NewEventCollector(&s.conf.Verifier, s.db, s.hub, signer.From())
}

func (s *server) startCollector(ctx context.Context) {
	if s.blockCollector == nil || s.eventCollector == nil {
		return
	}

	log.Info("Block collector started", "interval", s.conf.Verifier.Interval, "block-limit", s.conf.Verifier.BlockLimit)
	log.Info("Event collector started", "interval", s.conf.Verifier.Interval, "block-limit", s.conf.Verifier.BlockLimit)

	ticker := time.NewTicker(s.conf.Verifier.Interval)
	defer ticker.Stop()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		for {
			select {
			case <-ctx.Done():
				log.Info("Block collector stopped")
				return
			case <-ticker.C:
				s.blockCollector.Work(ctx)
				s.eventCollector.Work(ctx)
			}
		}
	}()
}

func (s *server) mustSetupVerifier() {
	if !s.conf.Verifier.Enable {
		return
	}

	signer, ok := s.signers[s.conf.Verifier.Wallet]
	if !ok {
		log.Crit("Wallet for the Verifier not found", "wallet", s.conf.Verifier.Wallet)
	}
	l1Signer := ethutil.NewSignableClient(new(big.Int).SetUint64(s.conf.HubLayer.ChainID), s.hub, signer)
	s.verifier = verifier.NewVerifier(&s.conf.Verifier, s.db, l1Signer)
}

func (s *server) startVerifier(ctx context.Context) {
	if s.verifier == nil {
		return
	}

	log.Info("Verifier started", "signer", s.verifier.L1Signer().Signer())

	// Start verifier ticker
	vTick := time.NewTicker(s.conf.Verifier.Interval)
	defer vTick.Stop()

	// Subscribe new signature from validators
	var (
		sub        = s.verifier.SubscribeNewSignature(ctx)
		subscribes = map[common.Address]*database.OptimismSignature{}
	)
	defer sub.Cancel()

	// Publish new signature via p2p
	debounce := time.NewTicker(time.Second * 5)
	defer debounce.Stop()

	// Optimize database every hour
	dbTick := util.NewTicker(s.conf.Verifier.OptimizeInterval, 1)
	defer dbTick.Stop()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		for {
			select {
			case <-ctx.Done():
				log.Info("Verifier stopped")
				return
			case <-vTick.C:
				s.verifier.Work(ctx)
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
			case <-dbTick.C:
				log.Info("Optimize database")
				s.db.OPSignature.RepairPreviousID(s.verifier.L1Signer().Signer())
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
	if s.conf.VerseLayer.Discovery.Endpoint == "" {
		// read verses from the configuration only, if the discovery endpoint is not set
		s.verseDiscoveryHandler(s.conf.VerseLayer.Directs)
		return
	}

	// dinamically discovered verses
	disc := config.NewVerseDiscovery(
		http.DefaultClient,
		s.conf.VerseLayer.Discovery.Endpoint,
		s.conf.VerseLayer.Discovery.RefreshInterval)

	discTick := time.NewTicker(s.conf.VerseLayer.Discovery.RefreshInterval)
	defer discTick.Stop()

	// Subscribed verses to verifier and submitter
	sub := disc.Subscribe(ctx)
	defer sub.Cancel()

	s.wg.Add(1)
	go func() {
		for {
			defer s.wg.Done()

			select {
			case <-ctx.Done():
				log.Info("Verse discovery stopped")
				return
			case verses := <-s.discoveredVerses:
				s.verseDiscoveryHandler(verses)
			case <-discTick.C:
				disc.Work(ctx)
			case s.discoveredVerses <- <-sub.Next():
				// publish the subscribed verses
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

				signer, ok := s.signers[tg.Wallet]
				if !ok {
					log.Error("Wallet for the Submitter not found", "wallet", tg.Wallet)
					continue
				}
				l1Signer := ethutil.NewSignableClient(new(big.Int).SetUint64(s.conf.HubLayer.ChainID), s.hub, signer)
				s.submitter.AddTask(x.verse.WithTransactable(l1Signer, x.verify))
			}
		}
	}
}

func (s *server) mustSetupBeacon() {
	if !s.conf.Beacon.Enable || !s.conf.Verifier.Enable {
		return
	}

	signer, ok := s.signers[s.conf.Verifier.Wallet]
	if !ok {
		log.Crit("Wallet for the Verifier not found", "wallet", s.conf.Verifier.Wallet)
	}

	// TODO: make sure the endpoint(s.conf.Beacon.Endpoint) is reachable here

	s.bw = beacon.NewBeaconWorker(
		&s.conf.Beacon,
		http.DefaultClient,
		beacon.Beacon{
			Signer:  signer.From().Hex(),
			Version: version.SemVer(),
			PeerID:  s.p2p.PeerID().String(),
		},
	)
}

func (s *server) startBeacon(ctx context.Context) {
	if s.bw == nil {
		return
	}
	s.bw.Start(ctx)
}

func getOrCreateP2PKey(filename string) (crypto.PrivKey, error) {
	data, err := os.ReadFile(filename)

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

	err = os.WriteFile(filename, []byte(enc), 0644)
	if err != nil {
		log.Error("Failed to write p2p private key", "err", err)
		return nil, err
	}

	log.Info("Generated and saved to p2p private key", "file", filename, "id", peerID)
	return priv, nil
}

func (s *server) mustLoadSigners(ctx context.Context, ipc *ipc.IPCServer) {
	// open geth keystore
	var ks *wallet.KeyStore
	if s.conf.Keystore != "" {
		ks = wallet.NewKeyStore(s.conf.Keystore)
		ipc.SetHandler(ipccmd.WalletUnlockCmd.NewHandler(ks))
	}

	var wg sync.WaitGroup
	wg.Add(len(s.conf.Wallets))

	for n, w := range s.conf.Wallets {
		go func(name string, wallet *config.Wallet) {
			defer wg.Done()
			address := common.HexToAddress(wallet.Address)

			// Plain text private key.
			if wallet.Plain != "" {
				priv, err := ethcrypto.HexToECDSA(strings.TrimPrefix(wallet.Plain, "0x"))
				if err != nil {
					log.Crit("Failed to decode private key",
						"name", name, "address", wallet.Address, "err", err)
				}

				signer := ethutil.NewPrivateKeySigner(priv)
				if signer.From() != address {
					log.Crit("Decrypted private key address does not "+
						"match the wallet address in the config",
						"name", name, "want", address, "got", signer.From())
				}

				s.signers[name] = signer
				log.Info("Loaded plaintext private key wallet", "name", name, "address", address)
				return
			}

			// go-ethereum's private key.
			if ks == nil {
				log.Crit("Keystore directory is not specified")
			}

			_wallet, account, err := ks.FindWallet(address)
			if err != nil {
				log.Crit("Failed to find the wallet",
					"name", name, "address", wallet.Address, "err", err)
			}

			if wallet.Password != "" {
				pw, err := os.ReadFile(wallet.Password)
				if err != nil {
					log.Crit("Failed to read password file",
						"name", name, "address", address, "err", err)
				}

				if err := ks.Unlock(*account, strings.Trim(string(pw), "\r\n\t ")); err != nil {
					log.Crit("Failed to unlock wallet using password file",
						"name", name, "address", address, "err", err)
				}
				log.Info("Wallet unlocked using password file", "name", name, "address", address)
			} else if ks.Unlock(*account, "") == nil {
				log.Info("Wallet unlocked using empty password", "name", name, "address", address)
			} else {
				log.Info("Waiting for wallet unlock via IPC", "name", name, "address", address)
				if err := ks.WaitForUnlock(ctx, _wallet); err != nil {
					log.Crit("Wallet was not unlocked",
						"name", name, "address", address, "err", err)
				}
				log.Info("Wallet unlocked via IPC", "name", name, "address", address)
			}

			s.signers[name] = ethutil.NewKeystoreSigner(_wallet, account)
		}(n, w)
	}

	wg.Wait()
}
