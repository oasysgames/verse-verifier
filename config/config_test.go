package config

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestParseConfig() {
	input := (`
	datastore: /tmp

	keystore: /tmp

	wallets:
		wallet1:
			address: '0xBA3186c30Bb0d9e8c7924147238F82617C3fE729'
			password: /etc/passwd
	
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
		concurrency: 10
		block_limit: 500
		event_filter_limit: 50
		state_collect_limit: 5
		state_collect_timeout: 1s
		db_optimize_interval: 2s
	
	submitter:
		enable: true
		interval: 5s
		concurrency: 10
		confirmations: 4
		gas_multiplier: 1.5
		batch_size: 100
		max_gas: 1_000
		verifier_address: '0xC79800039e6c4d6C29E10F2aCf2158516Fe686AA'
		multicall2_address: '0x74746c14ABD3b4e8B6317e279E8C9e27D9dA56E5'
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

	got, _ := NewConfig([]byte(strings.ReplaceAll(input, "\t", "  ")))

	s.Equal("/tmp", got.DataStore)

	s.Equal("/tmp", got.KeyStore)

	s.Equal(map[string]Wallet{
		"wallet1": {
			Address:  "0xBA3186c30Bb0d9e8c7924147238F82617C3fE729",
			Password: "/etc/passwd",
		},
	}, got.Wallets)

	s.Equal(hubLayer{
		ChainId: 12345,
		RPC:     "http://127.0.0.1:8545/",
	}, got.HubLayer)

	s.Equal(verseLayer{
		Discovery: struct {
			Endpoint        string        "json:\"endpoint\" validate:\"omitempty,url\""
			RefreshInterval time.Duration "json:\"refresh_interval\" mapstructure:\"refresh_interval\""
		}{
			Endpoint:        "http://127.0.0.1/api/v1/verse-layers.json",
			RefreshInterval: 5 * time.Second,
		},
		Directs: []*Verse{
			{
				ChainID: 12345,
				RPC:     "http://127.0.0.1:8545/",
				L1Contracts: map[string]string{
					"StateCommitmentChain": "0x62b105FD57A11819f9E50892E18a354bd7c89937",
				},
			},
		},
	}, got.VerseLayer)

	durationPtr := func(d time.Duration) *time.Duration { return &d }
	intPtr := func(d int) *int { return &d }
	int64Ptr := func(d int64) *int64 { return &d }
	s.Equal(P2P{
		Listens:          []string{"listen0"},
		NoAnnounce:       []string{"noann0"},
		ConnectionFilter: []string{"connfil0"},
		Transports: struct {
			TCP  bool "json:\"tcp\""
			QUIC bool "json:\"quic\""
		}{
			TCP:  true,
			QUIC: true,
		},
		Listen: "",
		Bootnodes: []string{
			"/ip4/127.0.0.1/tcp/20002/p2p/12D3KooWCNqRgVdwAhGrurCc8XE4RsWB8S2T83yMZR9R7Gdtf899",
		},
		NAT: struct {
			UPnP      bool "json:\"upnp\" mapstructure:\"upnp\""
			AutoNAT   bool "json:\"autonat\" mapstructure:\"autonat\""
			HolePunch bool "json:\"holepunch\" mapstructure:\"holepunch\""
		}{
			UPnP:      true,
			AutoNAT:   true,
			HolePunch: true,
		},
		RelayService: struct {
			Enable                 bool           "json:\"enable\""
			DurationLimit          *time.Duration "json:\"duration_limit,omitempty\" mapstructure:\"duration_limit\""
			DataLimit              *int64         "json:\"data_limit,omitempty\" mapstructure:\"data_limit\""
			ReservationTTL         *time.Duration "json:\"reservation_ttl,omitempty\" mapstructure:\"reservation_ttl\""
			MaxReservations        *int           "json:\"max_reservations,omitempty\" mapstructure:\"max_reservations\""
			MaxCircuits            *int           "json:\"max_circuits,omitempty\" mapstructure:\"max_circuits\""
			BufferSize             *int           "json:\"buffer_size,omitempty\" mapstructure:\"buffer_size\""
			MaxReservationsPerPeer *int           "json:\"max_reservations_per_peer,omitempty\" mapstructure:\"max_reservations_per_peer\""
			MaxReservationsPerIP   *int           "json:\"max_reservations_per_ip,omitempty\" mapstructure:\"max_reservations_per_ip\""
			MaxReservationsPerASN  *int           "json:\"max_reservations_per_asn,omitempty\" mapstructure:\"max_reservations_per_asn\""
		}{
			Enable:                 true,
			DurationLimit:          durationPtr(time.Minute),
			DataLimit:              int64Ptr(2),
			ReservationTTL:         durationPtr(3 * time.Minute),
			MaxReservations:        intPtr(4),
			MaxCircuits:            intPtr(5),
			BufferSize:             intPtr(6),
			MaxReservationsPerPeer: intPtr(7),
			MaxReservationsPerIP:   intPtr(8),
			MaxReservationsPerASN:  intPtr(9),
		},
		RelayClient: struct {
			Enable     bool     "json:\"enable\""
			RelayNodes []string "json:\"relay_nodes\" mapstructure:\"relay_nodes\""
		}{
			Enable:     true,
			RelayNodes: []string{"relay-0", "relay-1"},
		},
		PublishInterval: 5 * time.Minute,
		StreamTimeout:   10 * time.Second,
		OutboundLimits: struct {
			Concurrency int "json:\"concurrency\""
			Throttling  int "json:\"throttling\""
		}{
			Concurrency: 10,
			Throttling:  500,
		},
		InboundLimits: struct {
			Concurrency int           "json:\"concurrency\""
			Throttling  int           "json:\"throttling\""
			MaxSendTime time.Duration "json:\"max_send_time\" mapstructure:\"max_send_time\""
		}{
			Concurrency: 10,
			Throttling:  500,
			MaxSendTime: 30 * time.Second,
		},
	}, got.P2P)

	s.Equal(IPC{Sockname: "testsock"}, got.IPC)

	s.Equal(Verifier{
		Enable:              true,
		Wallet:              "wallet1",
		Interval:            5 * time.Second,
		Concurrency:         10,
		BlockLimit:          500,
		EventFilterLimit:    50,
		StateCollectLimit:   5,
		StateCollectTimeout: time.Second,
		OptimizeInterval:    time.Second * 2,
	}, got.Verifier)

	s.Equal(Submitter{
		Enable:            true,
		Concurrency:       10,
		Interval:          5 * time.Second,
		Confirmations:     4,
		GasMultiplier:     1.5,
		BatchSize:         100,
		MaxGas:            1_000,
		VerifierAddress:   "0xC79800039e6c4d6C29E10F2aCf2158516Fe686AA",
		Multicall2Address: "0x74746c14ABD3b4e8B6317e279E8C9e27D9dA56E5",
		Targets: []struct {
			ChainID uint64 "json:\"chain_id\"     mapstructure:\"chain_id\"     validate:\"required\""
			Wallet  string "json:\"wallet\" validate:\"required\""
		}{
			{
				ChainID: 12345,
				Wallet:  "wallet1",
			},
		},
	}, got.Submitter)

	s.Equal(Beacon{
		Enable:   true,
		Endpoint: "http://127.0.0.1/beacon",
		Interval: time.Second,
	}, got.Beacon)

	s.Equal(Database{
		LongQueryTime:       time.Second,
		MinExaminedRowLimit: 100,
	}, got.Database)

	s.Equal(Metrics{
		Enable:   true,
		Type:     "testcollector",
		Prefix:   "testprefix",
		Listen:   "127.0.0.1:3030",
		Endpoint: "/testmetrics",
	}, got.Metrics)

	s.Equal(Debug{
		Pprof: Pprof{
			Enable: true,
			Listen: "0.0.0.0:12345",
			BasicAuth: struct {
				Username string "json:\"username\""
				Password string "json:\"password\""
			}{
				Username: "my-username",
				Password: "my-password",
			},
			BlockProfileRate: 1,
			MemProfileRate:   2,
		},
	}, got.Debug)
}

func (s *ConfigTestSuite) TestValidate() {
	input := (`
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
			- xxx
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

	// parse config
	var config Config
	viper.ReadConfig(bytes.NewBuffer([]byte(strings.ReplaceAll(input, "\t", "  "))))
	viper.Unmarshal(&config)

	// do validation
	err := validate.Struct(&config)

	// assert
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

	got, _ := NewConfig([]byte(strings.ReplaceAll(input, "\t", "  ")))

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

	s.Equal(15*time.Second, got.Verifier.Interval)
	s.Equal(50, got.Verifier.Concurrency)
	s.Equal(1000, got.Verifier.BlockLimit)
	s.Equal(1000, got.Verifier.EventFilterLimit)
	s.Equal(1000, got.Verifier.StateCollectLimit)
	s.Equal(15*time.Second, got.Verifier.StateCollectTimeout)
	s.Equal(time.Hour, got.Verifier.OptimizeInterval)

	s.Equal(15*time.Second, got.Submitter.Interval)
	s.Equal(50, got.Submitter.Concurrency)
	s.Equal(6, got.Submitter.Confirmations)
	s.Equal(1.1, got.Submitter.GasMultiplier)
	s.Equal(20, got.Submitter.BatchSize)
	s.Equal(5_000_000, got.Submitter.MaxGas)
	s.Equal("0x5200000000000000000000000000000000000014", got.Submitter.VerifierAddress)
	s.Equal("0x5200000000000000000000000000000000000022", got.Submitter.Multicall2Address)

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
