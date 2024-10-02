package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

type ConfigLoaderTestSuite struct {
	testhelper.Suite

	datastoreDir,
	keystoreDir string

	confFile,
	passwdFile1,
	passwdFile2 *os.File
}

func TestConfigLoader(t *testing.T) {
	suite.Run(t, new(ConfigLoaderTestSuite))
}

func (s *ConfigLoaderTestSuite) SetupTest() {
	s.datastoreDir = s.T().TempDir()
	s.keystoreDir = s.T().TempDir()
	s.confFile, _ = ioutil.TempFile(s.datastoreDir, "")
	s.passwdFile1, _ = ioutil.TempFile(s.datastoreDir, "")
	s.passwdFile2, _ = ioutil.TempFile(s.datastoreDir, "")
}

func (s *ConfigLoaderTestSuite) TestLoadConfigFromYAML() {
	want := s.configWithMinCliArgs()
	s.applyP2PCliArgs(want)
	s.applyVerseCliArgs(want)
	s.applyVerifierCliArgs(want)
	s.applySubmitterCliArgs(want)

	yaml := fmt.Sprintf(`
	datastore: %s
	keystore: %s

	wallets:
		verifier:
			address: '0x08E9441C28c9f34dcB1fa06f773a0450f15B6F43'
			password: %s
			plain: '0x5ea366a14e0bd46e7da7e894c8cc896ebecd1f6452b674aaa41688878f45ff73'
		submitter:
			address: '0xD244F03CA3e99C6093f6cBEFBD2f4508244C59D4'
			password: %s
			plain: '0xebf3a7f5f805e02c0bbbd599acd5c881f40db22caa95127d4bf48e2dde5fd7bb'

	hub_layer:
		chain_id: 1
		rpc: https://rpc.hub.example.com/

	verse_layer:
		discovery:
			endpoint: https://discovery.example.com/
		directs:
			- chain_id: 2
			  rpc: https://rpc.verse.example.com/
			  l1_contracts:
			    StateCommitmentChain: '0x01E901F3c65fA7CBd4505F5eF3A88e4ce432e4B5'
			    L2OutputOracle: '0x2489317FA6e003550111D5D196302Ba0879354e2'

	p2p:
		listens:
			- listen0
			- listen1
		bootnodes:
			- bootnode0
			- bootnode1
		append_announce:
			- appendann0
			- appendann1
		no_announce:
			- noann0
			- noann1
		connection_filter:
			- connfil0
			- connfil1

	verifier:
		enable: true
		wallet: verifier
		max_retry_backoff: 1m
		retry_timeout: 2m

	submitter:
		enable: true
		confirmations: 10
		scc_verifier_address: '0x239eD34cE5d21afD99e11b9B8e1Ea6067981DE9a'
		l2oo_verifier_address: '0xD05dDB4b9f736530367AE984dE37877245EC05b8'
		multicall_address: '0x0664C632576A4CA04166D585c2f3620aBc0c65D9'
		targets:
			- chain_id: 2
			  wallet: submitter
			- chain_id: 3
			  wallet: submitter
	`,
		s.datastoreDir, s.keystoreDir,
		s.passwdFile1.Name(), s.passwdFile2.Name(),
	)

	// write yaml to tempfile
	s.confFile.WriteString(strings.ReplaceAll(yaml, "\t", "  "))

	// add dummy command
	subCmd := &cobra.Command{Use: "test:load-config"}
	rootCmd.AddCommand(subCmd)

	// run dummy command
	rootCmd.SetArgs([]string{
		subCmd.Use,
		"--config", s.confFile.Name(),
	})
	rootCmd.Execute()

	got, _ := globalConfigLoader.load(false)
	s.Equal(want, got)
}

func (s *ConfigLoaderTestSuite) TestLoadConfigWithMinCliArgs() {
	want := s.configWithMinCliArgs()
	got := s.executeWithCliArgs(nil)
	s.Equal(want, got)
}

func (s *ConfigLoaderTestSuite) TestLoadConfigWithP2PCliArgs() {
	want := s.configWithMinCliArgs()
	s.applyP2PCliArgs(want)

	got := s.executeWithCliArgs([]string{
		"--config.p2p.bootnodes", "bootnode0",
		"--config.p2p.bootnodes", "bootnode1",
		"--config.p2p.append_announce", "appendann0",
		"--config.p2p.append_announce", "appendann1",
		"--config.p2p.no_announce", "noann0",
		"--config.p2p.no_announce", "noann1",
		"--config.p2p.connection_filter", "connfil0",
		"--config.p2p.connection_filter", "connfil1",
	})
	s.Equal(want, got)

	// test if default values can be removed
	want = s.configWithMinCliArgs()
	want.P2P.NoAnnounce = []string{}
	want.P2P.ConnectionFilter = []string{}

	got = s.executeWithCliArgs([]string{
		"--config.p2p.no_announce", "",
		"--config.p2p.connection_filter", "",
	})
	s.Equal(want, got)
}

func (s *ConfigLoaderTestSuite) TestLoadConfigWithVerseArgs() {
	want := s.configWithMinCliArgs()
	s.applyVerseCliArgs(want)

	got := s.executeWithCliArgs([]string{
		"--config.verse",
		"--config.verse.chain_id", "2",
		"--config.verse.rpc", "https://rpc.verse.example.com/",
		"--config.verse.scc", "0x01E901F3c65fA7CBd4505F5eF3A88e4ce432e4B5",
		"--config.verse.l2oo", "0x2489317FA6e003550111D5D196302Ba0879354e2",
		"--config.verse.discovery", "https://discovery.example.com/",
	})

	s.Equal(want, got)
}

func (s *ConfigLoaderTestSuite) TestLoadConfigWithVerifierArgs() {
	want := s.configWithMinCliArgs()
	s.applyVerifierCliArgs(want)

	got := s.executeWithCliArgs([]string{
		"--config.verifier",
		"--config.verifier.wallet.address", "0x08E9441C28c9f34dcB1fa06f773a0450f15B6F43",
		"--config.verifier.wallet.password", s.passwdFile1.Name(),
		"--config.verifier.wallet.plain", "0x5ea366a14e0bd46e7da7e894c8cc896ebecd1f6452b674aaa41688878f45ff73",
		"--config.verifier.max-retry-backoff", "1m",
		"--config.verifier.retry-timeout", "2m",
	})

	s.Equal(want, got)
}

func (s *ConfigLoaderTestSuite) TestLoadConfigWithSubmitterArgs() {
	want := s.configWithMinCliArgs()
	s.applySubmitterCliArgs(want)

	got := s.executeWithCliArgs([]string{
		"--config.submitter",
		"--config.submitter.confirmations", "10",
		"--config.submitter.scc-verifier-address", "0x239eD34cE5d21afD99e11b9B8e1Ea6067981DE9a",
		"--config.submitter.l2oo-verifier-address", "0xD05dDB4b9f736530367AE984dE37877245EC05b8",
		"--config.submitter.multicall-address", "0x0664C632576A4CA04166D585c2f3620aBc0c65D9",
		"--config.submitter.targets", "2",
		"--config.submitter.targets", "3",
		"--config.submitter.wallet.address", "0xD244F03CA3e99C6093f6cBEFBD2f4508244C59D4",
		"--config.submitter.wallet.password", s.passwdFile2.Name(),
		"--config.submitter.wallet.plain", "0xebf3a7f5f805e02c0bbbd599acd5c881f40db22caa95127d4bf48e2dde5fd7bb",
	})

	s.Equal(want, got)
}

func (s *ConfigLoaderTestSuite) executeWithCliArgs(appendArgs []string) *config.Config {
	var cmd cobra.Command
	opts := mustNewConfigLoader(&cmd)

	cmd.SetArgs(append([]string{
		// set required flags
		"--config.cli",
		"--config.datastore", s.datastoreDir,
		"--config.keystore", s.keystoreDir,
		"--config.hub.chain_id", "1",
		"--config.hub.rpc", "https://rpc.hub.example.com/",
		"--config.p2p.listens", "listen0",
		"--config.p2p.listens", "listen1",
	}, appendArgs...))
	cmd.Execute()

	conf, err := opts.load(false)
	s.Require().NoError(err)
	return conf
}

func (s *ConfigLoaderTestSuite) configWithMinCliArgs() *config.Config {
	defaults := config.Defaults()

	return &config.Config{
		Datastore: s.datastoreDir,
		Keystore:  s.keystoreDir,
		Wallets:   map[string]*config.Wallet{},
		HubLayer: config.HubLayer{
			ChainID:   1,
			RPC:       "https://rpc.hub.example.com/",
			BlockTime: time.Second * 6,
		},
		VerseLayer: config.VerseLayer{
			Discovery: struct {
				Endpoint        string        "validate:\"omitempty,url\""
				RefreshInterval time.Duration "koanf:\"refresh_interval\""
			}{
				Endpoint:        "",
				RefreshInterval: defaults["verse_layer.discovery.refresh_interval"].(time.Duration),
			},
			Directs: nil,
		},
		P2P: config.P2P{
			Listens:          []string{"listen0", "listen1"},
			NoAnnounce:       defaults["p2p.no_announce"].([]string),
			ConnectionFilter: defaults["p2p.connection_filter"].([]string),
			Transports: struct {
				TCP  bool
				QUIC bool
			}{
				TCP:  defaults["p2p.transports.tcp"].(bool),
				QUIC: defaults["p2p.transports.quic"].(bool),
			},
			Listen:    "",
			Bootnodes: nil,
			NAT: struct {
				UPnP      bool "koanf:\"upnp\""
				AutoNAT   bool "koanf:\"autonat\""
				HolePunch bool "koanf:\"holepunch\""
			}{
				UPnP:      defaults["p2p.nat.upnp"].(bool),
				AutoNAT:   defaults["p2p.nat.autonat"].(bool),
				HolePunch: defaults["p2p.nat.holepunch"].(bool),
			},
			RelayService: struct {
				Enable                 bool
				DurationLimit          *time.Duration "koanf:\"duration_limit\""
				DataLimit              *int64         "koanf:\"data_limit\""
				ReservationTTL         *time.Duration "koanf:\"reservation_ttl\""
				MaxReservations        *int           "koanf:\"max_reservations\""
				MaxCircuits            *int           "koanf:\"max_circuits\""
				BufferSize             *int           "koanf:\"buffer_size\""
				MaxReservationsPerPeer *int           "koanf:\"max_reservations_per_peer\""
				MaxReservationsPerIP   *int           "koanf:\"max_reservations_per_ip\""
				MaxReservationsPerASN  *int           "koanf:\"max_reservations_per_asn\""
			}{
				Enable:                 false,
				DurationLimit:          nil,
				DataLimit:              nil,
				ReservationTTL:         nil,
				MaxReservations:        nil,
				MaxCircuits:            nil,
				BufferSize:             nil,
				MaxReservationsPerPeer: nil,
				MaxReservationsPerIP:   nil,
				MaxReservationsPerASN:  nil,
			},
			RelayClient: struct {
				Enable     bool
				RelayNodes []string "koanf:\"relay_nodes\""
			}{
				Enable:     defaults["p2p.relay_client.enable"].(bool),
				RelayNodes: nil,
			},
			PublishInterval: defaults["p2p.publish_interval"].(time.Duration),
			StreamTimeout:   defaults["p2p.stream_timeout"].(time.Duration),
			OutboundLimits: struct {
				Concurrency int
				Throttling  int
			}{
				Concurrency: defaults["p2p.outbound_limits.concurrency"].(int),
				Throttling:  defaults["p2p.outbound_limits.throttling"].(int),
			},
			InboundLimits: struct {
				Concurrency int
				Throttling  int
				MaxSendTime time.Duration "koanf:\"max_send_time\""
			}{
				Concurrency: defaults["p2p.inbound_limits.concurrency"].(int),
				Throttling:  defaults["p2p.inbound_limits.throttling"].(int),
				MaxSendTime: defaults["p2p.inbound_limits.max_send_time"].(time.Duration),
			},
		},
		IPC: config.IPC{
			Sockname: defaults["ipc.sockname"].(string),
		},
		Verifier: config.Verifier{
			Enable:                false,
			Wallet:                "",
			Interval:              defaults["verifier.interval"].(time.Duration),
			StateCollectLimit:     defaults["verifier.state_collect_limit"].(int),
			StateCollectTimeout:   defaults["verifier.state_collect_timeout"].(time.Duration),
			Confirmations:         defaults["verifier.confirmations"].(int),
			MaxLogFetchBlockRange: defaults["verifier.max_log_fetch_block_range"].(int),
			MaxIndexDiff:          defaults["verifier.max_index_diff"].(int),
			MaxRetryBackoff:       defaults["verifier.max_retry_backoff"].(time.Duration),
			RetryTimeout:          defaults["verifier.retry_timeout"].(time.Duration),
		},
		Submitter: config.Submitter{
			Enable:              false,
			Confirmations:       defaults["submitter.confirmations"].(int),
			Concurrency:         defaults["submitter.concurrency"].(int),
			Interval:            defaults["submitter.interval"].(time.Duration),
			GasMultiplier:       defaults["submitter.gas_multiplier"].(float64),
			BatchSize:           defaults["submitter.batch_size"].(int),
			MaxGas:              uint64(defaults["submitter.max_gas"].(int)),
			SCCVerifierAddress:  defaults["submitter.scc_verifier_address"].(string),
			L2OOVerifierAddress: defaults["submitter.l2oo_verifier_address"].(string),
			UseMulticall:        true,
			MulticallAddress:    defaults["submitter.multicall_address"].(string),
			Targets:             nil,
		},
		Beacon: config.Beacon{
			Enable:   defaults["beacon.enable"].(bool),
			Endpoint: defaults["beacon.endpoint"].(string),
			Interval: defaults["beacon.interval"].(time.Duration),
		},
		Database: config.Database{
			LongQueryTime:       defaults["database.long_query_time"].(time.Duration),
			MinExaminedRowLimit: defaults["database.min_examined_row_limit"].(int),
		},
		Metrics: config.Metrics{
			Enable:   false,
			Type:     defaults["metrics.type"].(string),
			Prefix:   defaults["metrics.prefix"].(string),
			Listen:   defaults["metrics.listen"].(string),
			Endpoint: defaults["metrics.endpoint"].(string),
		},
		Debug: config.Debug{
			Pprof: config.Pprof{
				Enable: false,
				Listen: defaults["debug.pprof.listen"].(string),
				BasicAuth: struct {
					Username string
					Password string
				}{
					Username: defaults["debug.pprof.basic_auth.username"].(string),
					Password: defaults["debug.pprof.basic_auth.password"].(string),
				},
				BlockProfileRate: defaults["debug.pprof.block_profile_rate"].(int),
				MemProfileRate:   defaults["debug.pprof.mem_profile_rate"].(int),
			},
		},
	}
}

func (s *ConfigLoaderTestSuite) applyP2PCliArgs(c *config.Config) {
	c.P2P.Bootnodes = []string{"bootnode0", "bootnode1"}
	c.P2P.AppendAnnounce = []string{"appendann0", "appendann1"}
	c.P2P.NoAnnounce = []string{"noann0", "noann1"}
	c.P2P.ConnectionFilter = []string{"connfil0", "connfil1"}
}

func (s *ConfigLoaderTestSuite) applyVerseCliArgs(c *config.Config) {
	c.VerseLayer.Discovery.Endpoint = "https://discovery.example.com/"
	c.VerseLayer.Directs = []*config.Verse{
		{
			ChainID: 2,
			RPC:     "https://rpc.verse.example.com/",
			L1Contracts: map[string]string{
				"StateCommitmentChain": "0x01E901F3c65fA7CBd4505F5eF3A88e4ce432e4B5",
				"L2OutputOracle":       "0x2489317FA6e003550111D5D196302Ba0879354e2",
			},
		},
	}
}

func (s *ConfigLoaderTestSuite) applyVerifierCliArgs(c *config.Config) {
	c.Wallets["verifier"] = &config.Wallet{
		Address:  "0x08E9441C28c9f34dcB1fa06f773a0450f15B6F43",
		Password: s.passwdFile1.Name(),
		Plain:    "0x5ea366a14e0bd46e7da7e894c8cc896ebecd1f6452b674aaa41688878f45ff73",
	}
	c.Verifier.Enable = true
	c.Verifier.Wallet = "verifier"
	c.Verifier.MaxRetryBackoff = time.Minute
	c.Verifier.RetryTimeout = time.Minute * 2
}

func (s *ConfigLoaderTestSuite) applySubmitterCliArgs(c *config.Config) {
	c.Wallets["submitter"] = &config.Wallet{
		Address:  "0xD244F03CA3e99C6093f6cBEFBD2f4508244C59D4",
		Password: s.passwdFile2.Name(),
		Plain:    "0xebf3a7f5f805e02c0bbbd599acd5c881f40db22caa95127d4bf48e2dde5fd7bb",
	}
	c.Submitter.Enable = true
	c.Submitter.Confirmations = 10
	c.Submitter.SCCVerifierAddress = "0x239eD34cE5d21afD99e11b9B8e1Ea6067981DE9a"
	c.Submitter.L2OOVerifierAddress = "0xD05dDB4b9f736530367AE984dE37877245EC05b8"
	c.Submitter.MulticallAddress = "0x0664C632576A4CA04166D585c2f3620aBc0c65D9"
	c.Submitter.Targets = []*config.SubmitterTarget{
		{
			ChainID: 2,
			Wallet:  "submitter",
		},
		{
			ChainID: 3,
			Wallet:  "submitter",
		},
	}
}
