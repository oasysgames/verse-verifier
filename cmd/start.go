package cmd

import (
	"context"
	"errors"
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
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/p2p"
	"github.com/oasysgames/oasys-optimism-verifier/verselayer"
	"github.com/oasysgames/oasys-optimism-verifier/wallet"
	"github.com/spf13/cobra"
)

const (
	SccName             = "StateCommitmentChain"
	StakeManagerAddress = "0x0000000000000000000000000000000000001001"
	VerifySccAddress    = "0x5200000000000000000000000000000000000014"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Verifier",
	Long:  "Start the Verifier",
	Run:   runStartCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().String(configFlag, "", "configuration file")
	startCmd.MarkFlagRequired(configFlag)
}

func runStartCmd(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	wg := &sync.WaitGroup{}

	// load configuration file
	conf, err := loadConfig(cmd)
	if err != nil {
		log.Crit("Failed to load configuration file", "err", err)
	}

	// setup database
	db, err := database.NewDatabase(conf.DatabasePath())
	if err != nil {
		log.Crit("Failed to open database", "err", err)
	}

	// open geth keystore
	ks := wallet.NewKeyStore(conf.KeyStore)

	// start ipc server
	// note: start ipc server before unlocking wallet
	ipc := newIPC(conf, ks)
	if ipc != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ipc.Start(ctx)
		}()
	}

	// unlock walelts(wait forever)
	waitForUnlockWallets(ctx, conf, ks)

	// create hub-layer client
	hub, err := ethutil.NewReadOnlyClient(conf.HubLayer.RPC)
	if err != nil {
		log.Crit("Failed to create hub-layer client", "err", err)
	}

	//  start p2p server
	p2p := newP2P(ctx, conf, db)
	wg.Add(1)
	go func() {
		defer wg.Done()
		p2p.Start(ctx)
	}()

	// start block collector
	bkCollector := newBlockCollector(ctx, conf, db, hub)
	if bkCollector != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bkCollector.Start(ctx)
		}()
	}

	// start event collector
	evCollector := newEventCollector(ctx, conf, db, hub)
	if evCollector != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			evCollector.Start(ctx)
		}()
	}

	// start state verifier
	sccVerifier := newSccVerifier(ctx, conf, ks, db)
	if sccVerifier != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sccVerifier.Start(ctx)
		}()

		go func() {
			sub := sccVerifier.SubscribeNewSignature(ctx)
			defer sub.Cancel()

			for {
				select {
				case <-ctx.Done():
					return
				case sig := <-sub.Next():
					p2p.PublishSignatures(ctx, []*database.OptimismSignature{sig})
				}
			}
		}()
	}

	// start signature submitter
	sccSubmitter := newSccSubmitter(ctx, conf, ks, db, hub)
	if sccSubmitter != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sccSubmitter.Start(ctx)
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		startVerseDiscovery(ctx, conf, ks, sccVerifier, sccSubmitter)
	}()

	wg.Wait()
	log.Info("Stopped all workers")
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

func waitForUnlockWallets(ctx context.Context, c *config.Config, ks *wallet.KeyStore) {
	wg := &sync.WaitGroup{}
	wg.Add(len(c.Wallets))

	for name, wallet := range c.Wallets {
		go func(name string, wallet config.Wallet) {
			defer wg.Done()

			_wallet, account, err := ks.FindWallet(common.HexToAddress(wallet.Address))
			if err != nil {
				log.Crit("Failed to find a wallet",
					"name", name, "address", wallet.Address, "err", err)
			}

			if wallet.Password != "" {
				b, err := ioutil.ReadFile(wallet.Password)
				if err != nil {
					log.Crit(
						"Failed to read password file",
						"name", name, "address", wallet.Address, "err", err)
				}

				if err := ks.Unlock(*account, strings.Trim(string(b), "\r\n\t ")); err != nil {
					log.Crit("Failed to unlock wallet using password file",
						"name", name, "address", wallet.Address, "err", err)
				}
			} else if ks.Unlock(*account, "") != nil {
				log.Info("Waiting for wallet unlock", "name", name, "address", wallet.Address)
				if err := ks.WaitForUnlock(ctx, _wallet); err != nil {
					log.Crit("Wallet was not unlocked",
						"name", name, "address", wallet.Address, "err", err)
				}
			}

			log.Info("Wallet unlocked", "name", name, "address", wallet.Address)
		}(name, wallet)
	}

	wg.Wait()
}

func newIPC(c *config.Config, ks *wallet.KeyStore) *ipc.IPCServer {
	if !c.IPC.Enable {
		return nil
	}

	ipc, err := ipc.NewIPCServer(commandName)
	if err != nil {
		log.Crit("Failed to create ipc server", "err", err)
	}

	ipc.SetHandler(ipccmd.WalletUnlockCmd.NewHandler(ks))
	return ipc
}

func newP2P(ctx context.Context, c *config.Config, db *database.Database) *p2p.Node {
	// get p2p private key
	p2pKey, err := getOrCreateP2PKey(c.P2PKeyPath())
	if err != nil {
		log.Crit(err.Error())
	}

	// setup p2p node
	listens := strings.Split(c.P2P.Listen, ":")
	host, dht, bwm, err := p2p.NewHost(ctx, listens[0], listens[1], p2pKey)
	if err != nil {
		log.Crit(err.Error())
	}

	// connect to bootstrap peers and setup peer discovery
	p2p.Bootstrap(ctx, host, dht, p2p.ConvertPeers(c.P2P.Bootnodes))

	node, err := p2p.NewNode(db, host, dht, bwm, c.P2P.PublishInterval)
	if err != nil {
		log.Crit("Failed to create p2p server", "err", err)
	}

	return node
}

func newBlockCollector(
	ctx context.Context,
	c *config.Config,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *hublayer.BlockCollector {
	if !c.Verifier.Enable {
		return nil
	}

	return hublayer.NewBlockCollector(db, hub, c.Verifier.Interval)
}

func newEventCollector(
	ctx context.Context,
	c *config.Config,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *hublayer.EventCollector {
	if !c.Verifier.Enable {
		return nil
	}

	return hublayer.NewEventCollector(
		db,
		hub,
		common.HexToAddress(c.Wallets[c.Verifier.Wallet].Address),
		c.Verifier.Interval,
		c.Verifier.BlockLimit,
	)
}

func newSccVerifier(
	ctx context.Context,
	c *config.Config,
	ks *wallet.KeyStore,
	db *database.Database,
) *verselayer.SccVerifier {
	if !c.Verifier.Enable {
		return nil
	}

	wallet, account := findWallet(c, ks, c.Verifier.Wallet)
	signer, err := ethutil.NewWritableClient(
		new(big.Int).SetUint64(c.HubLayer.ChainId),
		c.HubLayer.RPC,
		wallet,
		account,
	)
	if err != nil {
		log.Crit("Failed to create hub-layer clinet", "err", err)
	}

	return verselayer.NewSccVerifier(db, signer, c.Verifier.Interval, c.Verifier.Concurrency)
}

func newSccSubmitter(
	ctx context.Context,
	c *config.Config,
	ks *wallet.KeyStore,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *hublayer.SccSubmitter {
	if !c.Submitter.Enable {
		return nil
	}

	sm, err := stakemanager.NewStakemanager(common.HexToAddress(StakeManagerAddress), hub)
	if err != nil {
		log.Crit("Failed to create StakeManager", "err", err)
	}

	return hublayer.NewSccSubmitter(
		db,
		sm,
		common.HexToAddress(VerifySccAddress),
		c.Submitter.Interval,
		c.Submitter.Concurrency,
		c.Submitter.Confirmations,
		c.Submitter.GasMultiplier,
	)
}

func startVerseDiscovery(
	ctx context.Context,
	c *config.Config,
	ks *wallet.KeyStore,
	verifier *verselayer.SccVerifier,
	submitter *hublayer.SccSubmitter,
) {
	notify := make(chan struct{}, 1)
	verses := &sync.Map{}
	for _, v := range c.VerseLayer.Directs {
		verses.Store(v.ChainID, v)
	}
	notify <- struct{}{}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-notify:
				verses.Range(func(key, value any) bool {
					verse, ok := value.(*config.Verse)
					if !ok {
						return true
					}

					// get contract address
					var scc common.Address
					if s, ok := verse.L1Contracts[SccName]; !ok {
						return true
					} else {
						scc = common.HexToAddress(s)
					}

					// add verse to SccVerifier
					if c.Verifier.Enable {
						if client, err := ethutil.NewReadOnlyClient(verse.RPC); err != nil {
							log.Error("Failed to create verse-layer client", "err", err)
						} else if !verifier.HasVerse(scc, client) {
							verifier.AddVerse(scc, client)
						}
					}

					// add verse to SccSubmitter
					if c.Submitter.Enable && !submitter.HasVerse(scc) {
						for _, t := range c.Submitter.Targets {
							if t.ChainID != verse.ChainID {
								continue
							}

							wallet, account := findWallet(c, ks, t.Wallet)
							hubClient, err := ethutil.NewWritableClient(
								new(big.Int).SetUint64(c.HubLayer.ChainId),
								c.HubLayer.RPC,
								wallet,
								account,
							)
							if err != nil {
								log.Error("Failed to create hub-layer client", "err", err)
							} else {
								submitter.AddVerse(scc, hubClient)
							}

							break
						}
					}

					return true
				})
			}
		}
	}()

	if c.VerseLayer.Discovery.Endpoint == "" {
		return
	}

	// start verse discovery
	discv := config.NewVerseDiscovery(
		http.DefaultClient,
		c.VerseLayer.Discovery.Endpoint,
		c.VerseLayer.Discovery.RefreshInterval,
	)

	go func() {
		sub := discv.Subscribe(ctx)
		defer sub.Cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case verse := <-sub.Next():
				verses.Store(verse.ChainID, verse)
				notify <- struct{}{}
			}
		}
	}()

	time.Sleep(1 * time.Second)
	discv.Start(ctx)
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
