package config

import (
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestNewConfig() {
	input := (`
	datastore: /tmp
	keystore: /tmp

	wallets:
		wallet1:
			address: '0xBA3186c30Bb0d9e8c7924147238F82617C3fE729'
			password: /etc/passwd
			plain: '0x70ce1ba0e76547883c0999662d093dd3426d550ec783a6c775b0060bf4ee6d0f'

	hub_layer:
		chain_id: 12345
		rpc: http://127.0.0.1:8545/

	verse_layer:
		discovery:
			endpoint: http://127.0.0.1/api/v1/verse-layers.json
			refresh_interval: 5s

		directs:
			- chain_id: 12345
			  rpc: http://127.0.0.1:8545/
			  l1_contracts:
			    StateCommitmentChain: '0x62b105FD57A11819f9E50892E18a354bd7c89937'

	p2p:
		listens:
			- listen0
		no_announce:
			- noann0
		connection_filter:
			- connfil0
		bootnodes:
			- /ip4/127.0.0.1/tcp/20002/p2p/12D3KooWCNqRgVdwAhGrurCc8XE4RsWB8S2T83yMZR9R7Gdtf899
		relay_service:
			enable: true
			duration_limit: 1m
			data_limit: 2
			reservation_ttl: 3m
			max_reservations: 4
			max_circuits: 5
			buffer_size: 6
			max_reservations_per_peer: 7
			max_reservations_per_ip: 8
			max_reservations_per_asn: 9
		relay_client:
			relay_nodes: ["relay-0", "relay-1"]

	ipc:
		sockname: testsock

	verifier:
		enable: true
		wallet: wallet1
		interval: 5s
		state_collect_limit: 5
		state_collect_timeout: 1s
		confirmations: 4
		start_block_offset: 5760
		max_retry_backoff: 1m
		retry_timeout: 2m

	submitter:
		enable: true
		interval: 5s
		concurrency: 10
		confirmations: 4
		gas_multiplier: 1.5
		batch_size: 100
		max_gas: 1_000
		scc_verifier_address: '0xC79800039e6c4d6C29E10F2aCf2158516Fe686AA'
		l2oo_verifier_address: '0x67a16865f03F6d46a206EF894F7A56597E0152b7'
		use_multicall: true
		multicall_address: '0x74746c14ABD3b4e8B6317e279E8C9e27D9dA56E5'
		targets:
			- chain_id: 12345
			  wallet: wallet1

	beacon:
		enable: true
		endpoint: http://127.0.0.1/beacon
		interval: 1s

	database:
		long_query_time: 1s
		min_examined_row_limit: 100

	metrics:
		enable: true
		type: testcollector
		prefix: testprefix
		listen: 127.0.0.1:3030
		endpoint: /testmetrics

	debug:
		pprof:
			enable: true
			listen: 0.0.0.0:12345
			basic_auth:
				username: my-username
				password: my-password
			block_profile_rate: 1
			mem_profile_rate: 2
	`)

	want := &Config{
		Datastore: "/tmp",
		Keystore:  "/tmp",
		Wallets: map[string]*Wallet{
			"wallet1": {
				Address:  "0xBA3186c30Bb0d9e8c7924147238F82617C3fE729",
				Password: "/etc/passwd",
				Plain:    "0x70ce1ba0e76547883c0999662d093dd3426d550ec783a6c775b0060bf4ee6d0f",
			},
		},
		HubLayer: HubLayer{
			ChainID: 12345,
			RPC:     "http://127.0.0.1:8545/",
		},
		VerseLayer: VerseLayer{
			Discovery: struct {
				Endpoint        string        "validate:\"omitempty,url\""
				RefreshInterval time.Duration "koanf:\"refresh_interval\""
			}{
				Endpoint:        "http://127.0.0.1/api/v1/verse-layers.json",
				RefreshInterval: 5 * time.Second,
			},
			Directs: []*Verse{
				0: {
					ChainID: 12345,
					RPC:     "http://127.0.0.1:8545/",
					L1Contracts: map[string]string{
						"StateCommitmentChain": "0x62b105FD57A11819f9E50892E18a354bd7c89937",
					},
				},
			},
		},
		P2P: P2P{
			Listens:          []string{"listen0"},
			NoAnnounce:       []string{"noann0"},
			ConnectionFilter: []string{"connfil0"},
			Transports: struct {
				TCP  bool
				QUIC bool
			}{
				TCP:  true,
				QUIC: true,
			},
			Listen: "",
			Bootnodes: []string{
				"/ip4/127.0.0.1/tcp/20002/p2p/12D3KooWCNqRgVdwAhGrurCc8XE4RsWB8S2T83yMZR9R7Gdtf899",
			},
			NAT: struct {
				UPnP      bool "koanf:\"upnp\""
				AutoNAT   bool "koanf:\"autonat\""
				HolePunch bool "koanf:\"holepunch\""
			}{
				UPnP:      true,
				AutoNAT:   true,
				HolePunch: true,
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
				Enable:                 true,
				DurationLimit:          testhelper.Pointer(time.Minute),
				DataLimit:              testhelper.Pointer(int64(2)),
				ReservationTTL:         testhelper.Pointer(3 * time.Minute),
				MaxReservations:        testhelper.Pointer(int(4)),
				MaxCircuits:            testhelper.Pointer(int(5)),
				BufferSize:             testhelper.Pointer(int(6)),
				MaxReservationsPerPeer: testhelper.Pointer(int(7)),
				MaxReservationsPerIP:   testhelper.Pointer(int(8)),
				MaxReservationsPerASN:  testhelper.Pointer(int(9)),
			},
			RelayClient: struct {
				Enable     bool
				RelayNodes []string "koanf:\"relay_nodes\""
			}{
				Enable:     true,
				RelayNodes: []string{"relay-0", "relay-1"},
			},
			PublishInterval: 5 * time.Minute,
			StreamTimeout:   10 * time.Second,
			OutboundLimits: struct {
				Concurrency int
				Throttling  int
			}{
				Concurrency: 10,
				Throttling:  500,
			},
			InboundLimits: struct {
				Concurrency int
				Throttling  int
				MaxSendTime time.Duration "koanf:\"max_send_time\""
			}{
				Concurrency: 10,
				Throttling:  500,
				MaxSendTime: 30 * time.Second,
			},
		},
		IPC: IPC{Sockname: "testsock"},
		Verifier: Verifier{
			Enable:              true,
			Wallet:              "wallet1",
			Interval:            5 * time.Second,
			StateCollectLimit:   5,
			StateCollectTimeout: time.Second,
			Confirmations:       4,
			StartBlockOffset:    5760,
			MaxRetryBackoff:     time.Minute,
			RetryTimeout:        time.Minute * 2,
		},
		Submitter: Submitter{
			Enable:              true,
			Concurrency:         10,
			Interval:            5 * time.Second,
			Confirmations:       4,
			GasMultiplier:       1.5,
			BatchSize:           100,
			MaxGas:              1_000,
			SCCVerifierAddress:  "0xC79800039e6c4d6C29E10F2aCf2158516Fe686AA",
			L2OOVerifierAddress: "0x67a16865f03F6d46a206EF894F7A56597E0152b7",
			UseMulticall:        true,
			MulticallAddress:    "0x74746c14ABD3b4e8B6317e279E8C9e27D9dA56E5",
			Targets: []*SubmitterTarget{
				{
					ChainID: 12345,
					Wallet:  "wallet1",
				},
			},
		},
		Beacon: Beacon{
			Enable:   true,
			Endpoint: "http://127.0.0.1/beacon",
			Interval: time.Second,
		},
		Database: Database{
			LongQueryTime:       time.Second,
			MinExaminedRowLimit: 100,
		},
		Metrics: Metrics{
			Enable:   true,
			Type:     "testcollector",
			Prefix:   "testprefix",
			Listen:   "127.0.0.1:3030",
			Endpoint: "/testmetrics",
		},
		Debug: Debug{
			Pprof: Pprof{
				Enable: true,
				Listen: "0.0.0.0:12345",
				BasicAuth: struct {
					Username string
					Password string
				}{
					Username: "my-username",
					Password: "my-password",
				},
				BlockProfileRate: 1,
				MemProfileRate:   2,
			},
		},
	}

	got, _ := NewConfig(s.toBytes(input), false)

	s.Equal(want, got)
}

func (s *ConfigTestSuite) TestValidate() {
	input := (`
	keystore: /xxx
	verse_layer:
		discovery:
			endpoint: xxx
		directs:
			- rpc: xxx
			  l1_contracts:
			    test: xxx
	wallets:
		wallet1:
			address: xxx
			password: passw0rd
	p2p:
		listen: xxx
	verifier:
		enable: true
	submitter:
		targets:
			- {}
	metrics:
		listen: xxx
	`)

	wants := map[string]string{
		"Config.datastore":                                 "dir",
		"Config.keystore":                                  "dir",
		"Config.wallets[wallet1].address":                  "hexadecimal",
		"Config.wallets[wallet1].password":                 "file",
		"Config.hub_layer.chain_id":                        "required",
		"Config.hub_layer.rpc":                             "url",
		"Config.verse_layer.discovery.endpoint":            "url",
		"Config.verse_layer.directs[0].chain_id":           "required",
		"Config.verse_layer.directs[0].rpc":                "url",
		"Config.verse_layer.directs[0].l1_contracts[test]": "hexadecimal",
		"Config.p2p.listen":                                "hostname_port",
		"Config.verifier.wallet":                           "required_if",
		"Config.submitter.targets[0].chain_id":             "required",
		"Config.submitter.targets[0].wallet":               "required",
		"Config.metrics.listen":                            "hostname_port",
	}

	_, err := NewConfig(s.toBytes(input), false)

	gots := map[string]string{}
	for _, e := range err.(validator.ValidationErrors) {
		gots[e.Namespace()] = e.Tag()
	}

	s.Len(gots, len(wants))
	for field := range wants {
		s.Equal(wants[field], gots[field])
	}
}

func (s *ConfigTestSuite) TestDefaultValues() {
	input := (`
	datastore: /tmp
	keystore: /tmp

	hub_layer:
		chain_id: 12345
		rpc: http://127.0.0.1:8545/

	verse_layer:
		discovery:
			endpoint: http://127.0.0.1/

	p2p:
		listen: 127.0.0.1:20001
	`)

	got, err := NewConfig(s.toBytes(input), false)
	s.NoError(err)

	s.Equal(time.Hour, got.VerseLayer.Discovery.RefreshInterval)

	s.Equal([]string{
		"/ip4/127.0.0.1/ipcidr/8",
		"/ip4/10.0.0.0/ipcidr/8",
		"/ip4/172.16.0.0/ipcidr/12",
		"/ip4/192.168.0.0/ipcidr/16",
	}, got.P2P.NoAnnounce)
	s.Equal([]string{
		"/ip4/127.0.0.1/ipcidr/8",
		"/ip4/10.0.0.0/ipcidr/8",
		"/ip4/172.16.0.0/ipcidr/12",
		"/ip4/192.168.0.0/ipcidr/16",
	}, got.P2P.ConnectionFilter)
	s.Equal(true, got.P2P.Transports.TCP)
	s.Equal(true, got.P2P.Transports.QUIC)
	s.Equal(true, got.P2P.NAT.UPnP)
	s.Equal(true, got.P2P.NAT.AutoNAT)
	s.Equal(true, got.P2P.NAT.HolePunch)
	s.Equal(5*time.Minute, got.P2P.PublishInterval)
	s.Equal(10*time.Second, got.P2P.StreamTimeout)
	s.Equal(10, got.P2P.OutboundLimits.Concurrency)
	s.Equal(500, got.P2P.OutboundLimits.Throttling)
	s.Equal(10, got.P2P.InboundLimits.Concurrency)
	s.Equal(500, got.P2P.InboundLimits.Throttling)
	s.Equal(30*time.Second, got.P2P.InboundLimits.MaxSendTime)

	s.Equal("oasvlfy", got.IPC.Sockname)

	s.Equal(6*time.Second, got.Verifier.Interval)
	s.Equal(1000, got.Verifier.StateCollectLimit)
	s.Equal(15*time.Second, got.Verifier.StateCollectTimeout)
	s.Equal(3, got.Verifier.Confirmations)
	s.Equal(time.Hour, got.Verifier.MaxRetryBackoff)
	s.Equal(time.Hour*24, got.Verifier.RetryTimeout)

	s.Equal(30*time.Second, got.Submitter.Interval)
	s.Equal(50, got.Submitter.Concurrency)
	s.Equal(3, got.Submitter.Confirmations)
	s.Equal(1.1, got.Submitter.GasMultiplier)
	s.Equal(20, got.Submitter.BatchSize)
	s.Equal(uint64(5_000_000), got.Submitter.MaxGas)
	s.Equal("0x5200000000000000000000000000000000000014", got.Submitter.SCCVerifierAddress)
	s.Equal("0xF62fD2d4ef5a99C5bAa1effd0dc20889c5021E1c", got.Submitter.L2OOVerifierAddress)
	s.Equal(true, got.Submitter.UseMulticall)
	s.Equal("0x5200000000000000000000000000000000000022", got.Submitter.MulticallAddress)

	s.True(got.Beacon.Enable)
	s.Equal(
		"https://script.google.com/macros/s/AKfycbzJpDKyn271jbm5otk_BxGkrS2b1YdMQerVq2-XxLdTOdhUPKCZICqvagvGgByxx_nq0Q/exec",
		got.Beacon.Endpoint,
	)
	s.Equal(15*time.Minute, got.Beacon.Interval)

	s.Equal(200*time.Millisecond, got.Database.LongQueryTime)
	s.Equal(10000, got.Database.MinExaminedRowLimit)

	s.Equal("prometheus", got.Metrics.Type)
	s.Equal("oasvlfy", got.Metrics.Prefix)
	s.Equal("127.0.0.1:9200", got.Metrics.Listen)
	s.Equal("/metrics", got.Metrics.Endpoint)

	s.Equal("127.0.0.1:6060", got.Debug.Pprof.Listen)
	s.Equal("username", got.Debug.Pprof.BasicAuth.Username)
	s.Equal("password", got.Debug.Pprof.BasicAuth.Password)
	s.Equal(0, got.Debug.Pprof.BlockProfileRate)
	s.Equal(524288, got.Debug.Pprof.MemProfileRate)
}

func (s *ConfigTestSuite) toBytes(yaml string) []byte {
	return []byte(strings.ReplaceAll(yaml, "\t", "  "))
}
