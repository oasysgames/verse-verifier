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
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/beacon"
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/debug"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/metrics"
	"github.com/oasysgames/oasys-optimism-verifier/p2p"
	"github.com/oasysgames/oasys-optimism-verifier/submitter"
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
	s.mustSetupVerifier()
	s.setupSubmitter()
	s.mustSetupBeacon()

	// Fetch the total stake and the stakes synchronously
	if err := s.smcache.Refresh(ctx); err != nil {
		// Exit if the first refresh faild, because the following refresh higly likely fail
		log.Crit("Failed to refresh stake cache", "err", err)
	}
	// start cache updater
	go func() {
		// NOTE: Don't add wait group, as no need to guarantee the completion
		s.smcache.RefreshLoop(ctx, time.Hour)
	}()

	s.startVerseDiscovery(ctx)
	s.startBeacon(ctx)
	s.startVerifier(ctx)
	s.startSubmitter(ctx)
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
	wg        sync.WaitGroup
	conf      *config.Config
	db        *database.Database
	signers   map[string]ethutil.Signer
	hub       ethutil.Client
	smcache   *stakemanager.Cache
	p2p       *p2p.Node
	versepool verse.VersePool
	verifier  *verifier.Verifier
	submitter *submitter.Submitter
	bw        *beacon.BeaconWorker
	msvr      *http.Server
	psvr      *http.Server
	ipc       *ipc.IPCServer
}

func mustNewServer(ctx context.Context) *server {
	var err error

	s := &server{
		signers: map[string]ethutil.Signer{},
	}

	if s.conf, err = globalConfigLoader.load(true); err != nil {
		log.Crit("Failed to load configuration", "err", err)
	}

	log.Info("Loaded configuration", "conf", s.conf)

	// setup database
	if s.conf.Database.Path == "" {
		s.conf.Database.Path = s.conf.DatabasePath()
	}
	if s.db, err = database.NewDatabase(&s.conf.Database); err != nil {
		log.Crit("Failed to open database", "err", err)
	}

	// construct hub-layer client
	if s.hub, err = ethutil.NewClient(s.conf.HubLayer.RPC, s.conf.HubLayer.BlockTime); err != nil {
		log.Crit("Failed to construct hub-layer client", "err", err)
	}

	// Make sue the s.hub can connect to the chain
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if _, err := s.hub.HeaderWithCache(ctx); err != nil {
		log.Crit("Failed to connect to the hub-layer chain", "err", err)
	}

	// construct global verse pool
	s.versepool = verse.NewVersePool(s.hub)

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
		if err := metrics.ListenAndServe(ctx, s.msvr); err != nil {
			// `ErrServerClosed` is thrown when `Shutdown` is intentionally called
			if !errors.Is(err, http.ErrServerClosed) {
				log.Crit("Failed to start metrics server", "err", err)
			}
		}
		log.Info("Metrics server have exited listening", "addr", s.conf.Metrics.Listen)
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
		if err := ps.ListenAndServe(ctx, s.psvr); err != nil {
			// `ErrServerClosed` is thrown when `Shutdown` is intentionally called
			if !errors.Is(err, http.ErrServerClosed) {
				log.Crit("Failed to start pprof server", "err", err)
			}
		}
		log.Info("pprof server have exited listening", "addr", s.conf.Debug.Pprof.Listen)
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
		log.Info("IPC server has stopped, decrement wait group")
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

		enableSubscriber := s.conf.Submitter.Enable
		s.p2p.Start(ctx, enableSubscriber)
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
	s.verifier = verifier.NewVerifier(
		&s.conf.Verifier, s.db, s.p2p, l1Signer, ethutil.NewClient, s.versepool)
}

func (s *server) setupSubmitter() {
	if !s.conf.Submitter.Enable {
		return
	}

	var newSignerFn submitter.L1SignerFn = func(chainID uint64) ethutil.SignableClient {
		for _, cfg := range s.conf.Submitter.Targets {
			if cfg.ChainID == chainID {
				if signer, ok := s.signers[cfg.Wallet]; ok {
					return ethutil.NewSignableClient(
						new(big.Int).SetUint64(s.conf.HubLayer.ChainID), s.hub, signer)
				}
			}
		}
		return nil
	}
	s.submitter = submitter.NewSubmitter(&s.conf.Submitter, s.db, newSignerFn, s.smcache, s.versepool)
}

func (s *server) startVerseDiscovery(ctx context.Context) {
	if len(s.conf.VerseLayer.Directs) != 0 {
		// read verses from the configuration
		s.verseDiscoveryHandler(ctx, s.conf.VerseLayer.Directs)
	}

	if s.conf.VerseLayer.Discovery.Endpoint == "" {
		// Disable dinamically discovered verses, if the endpoint is not set
		return
	}

	// dinamically discovered verses
	disc, err := config.NewVerseDiscovery(
		ctx,
		http.DefaultClient,
		s.conf.VerseLayer.Discovery.Endpoint,
		s.conf.VerseLayer.Discovery.RefreshInterval,
	)
	if err != nil {
		log.Crit("Failed to construct verse discovery", "err", err)
	}

	// Subscribed verses to verifier and submitter
	discSub := disc.Subscribe(ctx)

	// synchronously try the first discovery
	if err := disc.Work(ctx); err != nil {
		// exit if the first discovery faild, because the following discovery highly likely fail
		log.Crit("Failed to work verse discovery", "err", err)
	}

	s.wg.Add(1)
	go func() {
		defer func() {
			defer s.wg.Done()
			discSub.Cancel()
			log.Info("Verse discovery has stopped, decrement wait group")
		}()

		discTick := time.NewTicker(s.conf.VerseLayer.Discovery.RefreshInterval)
		defer discTick.Stop()

		log.Info("Verse discovery started", "endpoint", s.conf.VerseLayer.Discovery.Endpoint, "interval", s.conf.VerseLayer.Discovery.RefreshInterval)

		for {
			select {
			case <-ctx.Done():
				log.Info("Verse discovery stopped")
				return
			case verses := <-discSub.Next():
				s.verseDiscoveryHandler(ctx, verses)
			case <-discTick.C:
				if err := disc.Work(ctx); err != nil {
					log.Error("Failed to work verse discovery", "err", err)
				}
			}
		}
	}()
}

func (s *server) verseDiscoveryHandler(ctx context.Context, discovers []*config.Verse) {
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

	// Delete erased Verse-Layer from the discovery JSON from the pool.
	erased := make(map[common.Address]bool)
	s.versepool.Range(func(item *verse.VersePoolItem) bool {
		erased[item.Verse().RollupContract()] = true
		return true
	})

	// Marking the Verse-Layer to be processed by the Submitter.
	canSubmits := make(map[uint64]bool)
	for _, cfg := range s.conf.Submitter.Targets {
		canSubmits[cfg.ChainID] = true
	}

	// Create a new Verse instance and add it to the pool.
	var chainIDs []uint64
	for _, cfg := range discovers {
		for name, addr := range cfg.L1Contracts {
			if factory, ok := verseFactories[name]; ok {
				verse := factory(s.db, s.hub, cfg.ChainID,
					cfg.RPC, common.HexToAddress(addr), verifyContracts[name])
				if s.versepool.Add(verse, canSubmits[cfg.ChainID]) {
					log.Info("Add verse to verse pool",
						"chain-id", cfg.ChainID, "rpc", cfg.RPC)
				}

				delete(erased, verse.RollupContract())
				chainIDs = append(chainIDs, cfg.ChainID)
			}
		}
	}

	// Delete erased verses from the pool.
	for contract := range erased {
		if verse, ok := s.versepool.Get(contract); ok {
			log.Info("Delete verse from verse pool", "chain-id", verse.Verse().ChainID())
			s.versepool.Delete(contract)
		}
	}

	log.Info("Discovered verses", "count", len(chainIDs), "chain-ids", chainIDs)
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
	go func() {
		s.bw.Start(ctx)
	}()
}

func (s *server) startVerifier(ctx context.Context) {
	if s.verifier == nil {
		return
	}
	s.wg.Add(1)
	go func() {
		s.verifier.Start(ctx)
		s.wg.Done()
	}()
}

func (s *server) startSubmitter(ctx context.Context) {
	if s.submitter == nil {
		return
	}
	s.wg.Add(1)
	go func() {
		s.submitter.Start(ctx)
		s.wg.Done()
	}()
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
