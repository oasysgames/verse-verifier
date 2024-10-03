package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

const (
	l1BlockTime = time.Second * 6
)

var (
	validate = validator.New()
)

func init() {
	// Convert error message to use the field names from
	// configuration file instead of the struct field names.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("koanf"), ",", 2)[0]
		switch name {
		case "":
			return strings.ToLower(fld.Name)
		case "-":
			return ""
		default:
			return name
		}
	})
}

func Defaults() map[string]interface{} {
	return map[string]interface{}{
		"hub_layer.block_time": l1BlockTime,

		"verse_layer.discovery.refresh_interval": time.Hour,

		"p2p.no_announce": []string{
			"/ip4/127.0.0.1/ipcidr/8",
			"/ip4/10.0.0.0/ipcidr/8",
			"/ip4/172.16.0.0/ipcidr/12",
			"/ip4/192.168.0.0/ipcidr/16",
		},
		"p2p.connection_filter": []string{
			"/ip4/127.0.0.1/ipcidr/8",
			"/ip4/10.0.0.0/ipcidr/8",
			"/ip4/172.16.0.0/ipcidr/12",
			"/ip4/192.168.0.0/ipcidr/16",
		},
		"p2p.transports.tcp":               true,
		"p2p.transports.quic":              true,
		"p2p.nat.upnp":                     true,
		"p2p.nat.autonat":                  true,
		"p2p.nat.holepunch":                true,
		"p2p.relay_client.enable":          true,
		"p2p.publish_interval":             5 * time.Minute,
		"p2p.stream_timeout":               10 * time.Second,
		"p2p.outbound_limits.concurrency":  10,
		"p2p.outbound_limits.throttling":   500,
		"p2p.inbound_limits.concurrency":   10,
		"p2p.inbound_limits.throttling":    500,
		"p2p.inbound_limits.max_send_time": 30 * time.Second,

		"ipc.sockname": "oasvlfy",

		"verifier.max_workers":               10,
		"verifier.interval":                  l1BlockTime,
		"verifier.state_collect_limit":       1000,
		"verifier.state_collect_timeout":     15 * time.Second,
		"verifier.confirmations":             3,               // 3 confirmations are enough for later than v1.3.0 L1.
		"verifier.max_log_fetch_block_range": 14400,           // 1 day in case of 6s block time
		"verifier.max_index_diff":            86400 * 2 / 120, // Number of rollups for 2days(L2BlockTime=1s,RollupInterval=120s)
		"verifier.max_retry_backoff":         time.Minute * 5,
		"verifier.retry_timeout":             time.Hour,

		// The minimum interval for Verse v0 is 15 seconds.
		// On the other hand, the minimum interval for Verse v1 is 80 seconds.
		// Balance the two by setting the default to 30 seconds.
		"submitter.max_workers":           5,
		"submitter.interval":              30 * time.Second,
		"submitter.confirmations":         3, // 3 confirmations are enough for later than v1.3.0 L1.
		"submitter.gas_multiplier":        1.1,
		"submitter.batch_size":            20,
		"submitter.max_gas":               5_000_000,
		"submitter.scc_verifier_address":  "0x5200000000000000000000000000000000000014",
		"submitter.l2oo_verifier_address": "0xF62fD2d4ef5a99C5bAa1effd0dc20889c5021E1c",
		"submitter.use_multicall":         true,
		"submitter.multicall_address":     "0x5200000000000000000000000000000000000022",

		"beacon.enable":   true,
		"beacon.endpoint": "https://script.google.com/macros/s/AKfycbzJpDKyn271jbm5otk_BxGkrS2b1YdMQerVq2-XxLdTOdhUPKCZICqvagvGgByxx_nq0Q/exec",
		"beacon.interval": 15 * time.Minute,

		"database.long_query_time":        200 * time.Millisecond,
		"database.min_examined_row_limit": 10000,

		"metrics.type":     "prometheus",
		"metrics.prefix":   "oasvlfy",
		"metrics.listen":   "127.0.0.1:9200",
		"metrics.endpoint": "/metrics",

		"debug.pprof.listen":              "127.0.0.1:6060",
		"debug.pprof.basic_auth.username": "username",
		"debug.pprof.basic_auth.password": "password",
		"debug.pprof.block_profile_rate":  0,
		"debug.pprof.mem_profile_rate":    524288,
	}
}

// Build configuration.
func NewConfig(input []byte, enableStrictValidation bool) (*Config, error) {
	k := koanf.New(".")

	// load default values
	if err := k.Load(confmap.Provider(Defaults(), "."), nil); err != nil {
		return nil, err
	}

	// load yaml configuration file
	err := k.Load(rawbytes.Provider(input), yaml.Parser())
	if err != nil {
		return nil, err
	}

	// assign values to `Config` struct
	var conf Config
	if err := k.Unmarshal("", &conf); err != nil {
		return nil, err
	}

	// run validation
	if err := Validate(&conf, enableStrictValidation); err != nil {
		return nil, err
	}

	return &conf, nil
}

// Returns a Config with default values set.
// However, the returned config is missing required fields, so it cannot be used as is.
func MustNewDefaultConfig() *Config {
	k := koanf.New(".")

	if err := k.Load(confmap.Provider(Defaults(), "."), nil); err != nil {
		panic(err)
	}

	var conf Config
	if err := k.Unmarshal("", &conf); err != nil {
		panic(err)
	}

	return &conf
}

func Validate(conf *Config, strict bool) error {
	if err := validate.Struct(conf); err != nil {
		return err
	}
	if strict {
		// validate verse discovery configuration
		if conf.VerseLayer.Discovery.Endpoint == "" && len(conf.VerseLayer.Directs) == 0 {
			return errors.New("either verse.discovery or verse.directs must be set")
		}
		// NOTE: Commented out because bootnode disable verifier and submitter
		// validate verifier and submitter configuration
		// if !conf.Verifier.Enable && !conf.Submitter.Enable {
		// 	return errors.New("either verifier.enable or submitter.enable must be set")
		// }
	}
	return nil
}

// App configuration.
type Config struct {
	// Datastore directory path.
	Datastore string `validate:"dir"`

	// Validator keystore directory path.
	Keystore string `validate:"omitempty,dir"`

	// Address used to create signatures and send transactions.
	Wallets map[string]*Wallet `validate:"dive"`

	// Configuration of the Hub-Layer.
	HubLayer HubLayer `koanf:"hub_layer"`

	// Configuration of Verse-Layer discovery.
	VerseLayer VerseLayer `koanf:"verse_layer"`

	// P2P worker configuration.
	P2P P2P

	// IPC configuration.
	IPC IPC

	// Verifier configuration.
	Verifier Verifier

	// Submitter configuration.
	Submitter Submitter

	// Beacon worker configuration.
	Beacon Beacon

	// Database configuration.
	Database Database

	// Metrics configuration
	Metrics Metrics

	// Debug configuration.
	Debug Debug
}

func (c *Config) DatabasePath() string {
	return filepath.Join(c.Datastore, "db.sqlite")
}

func (c *Config) P2PKeyPath() string {
	return filepath.Join(c.Datastore, "p2p.key")
}

type Wallet struct {
	// Address of the wallet.
	Address string `validate:"hexadecimal"`

	// Password file of the wallet.
	Password string `validate:"omitempty,file"`

	// Hex-encoded plaintext private key.
	Plain string `validate:"omitempty,hexadecimal"`
}

type HubLayer struct {
	// Chain ID of the Hub-Layer.
	ChainID uint64 `koanf:"chain_id" validate:"required"`

	// RPC of the Hub-Layer(HTTP or WebSocket).
	RPC string `validate:"url"`

	// Block interval of the Hub-Layer.
	BlockTime time.Duration `koanf:"block_time"`
}

type Verse struct {
	// Chain ID of the Verse-Layer.
	ChainID uint64 `json:"chain_id" koanf:"chain_id" validate:"required"`

	// RPC of the Verse-Layer(HTTP or WebSocket).
	RPC string `json:"rpc" validate:"url"`

	// Contract addresses on the Hub-Layer.
	L1Contracts map[string]string `json:"l1_contracts" koanf:"l1_contracts" validate:"required,dive,hexadecimal"`
}

type VerseLayer struct {
	// Discover from API.
	Discovery struct {
		Endpoint        string        `validate:"omitempty,url"`
		RefreshInterval time.Duration `koanf:"refresh_interval"`
	}

	// List of Verse-Layer.
	Directs []*Verse `validate:"dive"`
}

type P2P struct {
	// libp2p multi-addresses to listen.
	Listens []string

	// Additional multi-addresses to advertise.
	AppendAnnounce []string `koanf:"append_announce"`

	// Multi-addresses not advertised.
	NoAnnounce []string `koanf:"no_announce"`

	// Multi-addresses that filter dial or receive connections.
	ConnectionFilter []string `koanf:"connection_filter"`

	// Enabled transport protocols.
	Transports struct {
		TCP  bool
		QUIC bool
	}

	// Deprecated: Address and port to listen.
	Listen string `validate:"omitempty,hostname_port"`

	// Initial node list.
	Bootnodes []string

	// Enabled NAT Travasal features.
	NAT struct {
		UPnP      bool `koanf:"upnp"`
		AutoNAT   bool `koanf:"autonat"`
		HolePunch bool `koanf:"holepunch"`
	}

	// Enable Circuit Relay(v2) service.
	// Note: Public connectivity is required.
	RelayService struct {
		Enable bool

		// DurationLimit is the limit of data relayed (on each direction) before resetting the connection.
		DurationLimit *time.Duration `koanf:"duration_limit"`
		// DataLimit is the time limit before resetting a relayed connection.
		DataLimit *int64 `koanf:"data_limit"`

		// ReservationTTL is the duration of a new (or refreshed reservation).
		ReservationTTL *time.Duration `koanf:"reservation_ttl"`

		// MaxReservations is the maximum number of active relay slots.
		MaxReservations *int `koanf:"max_reservations"`
		// MaxCircuits is the maximum number of open relay connections for each peer; defaults to 16.
		MaxCircuits *int `koanf:"max_circuits"`
		// BufferSize is the size of the relayed connection buffers.
		BufferSize *int `koanf:"buffer_size"`

		// MaxReservationsPerPeer is the maximum number of reservations originating from the same peer.
		MaxReservationsPerPeer *int `koanf:"max_reservations_per_peer"`
		// MaxReservationsPerIP is the maximum number of reservations originating from the same IP address.
		MaxReservationsPerIP *int `koanf:"max_reservations_per_ip"`
		// MaxReservationsPerASN is the maximum number of reservations origination from the same ASN.
		MaxReservationsPerASN *int `koanf:"max_reservations_per_asn"`
	} `koanf:"relay_service"`

	// Enable Circuit Relay(v2) client.
	RelayClient struct {
		Enable     bool
		RelayNodes []string `koanf:"relay_nodes"`
	} `koanf:"relay_client"`

	// Interval to publish own signature status.
	PublishInterval time.Duration `koanf:"publish_interval"`

	// Timeout for P2P stream communication.
	StreamTimeout time.Duration `koanf:"stream_timeout"`

	OutboundLimits struct {
		// Maximum number of concurrent signature requests from oneself to peers.
		Concurrency int

		// The number of signatures that can be sent to peers per second.
		Throttling int
	} `koanf:"outbound_limits"`

	InboundLimits struct {
		// Maximum number of concurrent signature requests from peers to oneself.
		Concurrency int

		// The number of signatures that can be sent to peers per second.
		Throttling int

		// Maximum time to send signatures to a peer.
		MaxSendTime time.Duration `koanf:"max_send_time"`
	} `koanf:"inbound_limits"`

	// Options for go-libp2p-kad-dht/LanDHT
	ExperimentalLanDHT struct {
		Loopback  bool
		Bootnodes []string
	} `koanf:"experimental_lan_dht"`
}

type IPC struct {
	// Socket file name, In UNIX-based OS, it is created as /tmp/{sockname}.sock.
	Sockname string
}

type Verifier struct {
	// Enable to verifier.
	Enable bool

	// Name of the wallet to create signature.
	Wallet string `validate:"required_if=Enable true"`

	// Maximum number of concurrent workers.
	MaxWorkers int `koanf:"max_workers"`

	// Interval for get block data.
	Interval time.Duration

	// Number of state root to collect at a time.
	StateCollectLimit int `koanf:"state_collect_limit"`

	// Timeout for state root collection.
	StateCollectTimeout time.Duration `koanf:"state_collect_timeout"`

	// Number of confirmation blocks for transaction receipt.
	Confirmations int

	// The max block range to fetch events.
	MaxLogFetchBlockRange int `koanf:"max_log_fetch_block_range"`

	// Do not verify if `rollup index - next index` is greater than this value.
	MaxIndexDiff int `koanf:"max_index_diff"`

	// The maximum exponential backoff time for retries.
	MaxRetryBackoff time.Duration `koanf:"max_retry_backoff"`

	// The maximum duration to attempt retries.
	RetryTimeout time.Duration `koanf:"retry_timeout"`
}

func (c *Verifier) String() string {
	return fmt.Sprintf(
		"wallet:%s max_workers:%d interval:%s state_collect_limit:%d state_collect_timeout:%s"+
			" confirmations:%d max_log_fetch_block_range:%d max_index_diff:%d max_retry_backoff:%s"+
			" retry_timeout:%s",
		c.Wallet, c.MaxWorkers, c.Interval, c.StateCollectLimit, c.StateCollectTimeout,
		c.Confirmations, c.MaxLogFetchBlockRange, c.MaxIndexDiff, c.MaxRetryBackoff,
		c.RetryTimeout)
}

type Submitter struct {
	// Whether to enable worker.
	Enable bool `koanf:"enable"`

	// Maximum number of concurrent workers.
	MaxWorkers int `koanf:"max_workers"`

	// Interval for send transaction.
	Interval time.Duration

	// Number of confirmation blocks for transaction receipt.
	Confirmations int

	// How much to increase the estimated gas limit.
	GasMultiplier float64 `koanf:"gas_multiplier"`

	// Maximum number of calls for Multicall2.
	BatchSize int `koanf:"batch_size"`

	// Maximum gas of calls for Multicall2.
	MaxGas uint64 `koanf:"max_gas"`

	// Address of the OasysStateCommitmentChainVerifier contract.
	SCCVerifierAddress string `koanf:"scc_verifier_address"`

	// Address of the OasysL2OutputOracleVerifier contract.
	L2OOVerifierAddress string `koanf:"l2oo_verifier_address"`

	// Address of the Multicall2 contract.
	UseMulticall     bool   `koanf:"use_multicall"`
	MulticallAddress string `koanf:"multicall_address"`

	// List of verses to submit signatures
	Targets []*SubmitterTarget `validate:"dive"`
}

func (c *Submitter) String() string {
	var targets []string
	for _, tg := range c.Targets {
		targets = append(targets, fmt.Sprintf("{chain_id=%d wallet=%s}", tg.ChainID, tg.Wallet))
	}

	return fmt.Sprintf("max_workers:%d interval:%s confirmations:%d gas_multiplier:%f"+
		" max_gas:%d batch_size:%d scc_verifier_address:%s l2oo_verifier_address:%s"+
		" use_multicall:%v multicall_address:%s targets:[%s]",
		c.MaxWorkers, c.Interval, c.Confirmations, c.GasMultiplier,
		c.MaxGas, c.BatchSize, c.SCCVerifierAddress, c.L2OOVerifierAddress,
		c.UseMulticall, c.MulticallAddress, strings.Join(targets, ","))
}

type SubmitterTarget struct {
	// Chain ID of the Verse-Layer.
	ChainID uint64 `koanf:"chain_id" validate:"required"`

	// Name of the wallet to send transaction.
	Wallet string `validate:"required"`
}

func (c *Submitter) MultiplyGas(base uint64) uint64 {
	return uint64(float64(base) * c.GasMultiplier)
}

type Beacon struct {
	// Whether to enable worker.
	Enable bool

	// URL of beacon.
	Endpoint string `validate:"omitempty,url"`

	// Interval for send beacon.
	Interval time.Duration
}

type Database struct {
	// File path of the SQLite database.
	Path string

	// Slow query log configurations.
	LongQueryTime       time.Duration `koanf:"long_query_time"`
	MinExaminedRowLimit int           `koanf:"min_examined_row_limit"`
}

type Metrics struct {
	// Whether to pprof server.
	Enable bool

	// Address and port to listen.
	Listen string `validate:"hostname_port"`

	// The URL used to retrieve metrics.
	Endpoint string

	// The type of metrics collector.
	Type string

	// Metric name prefix.
	Prefix string
}

type Debug struct {
	Pprof Pprof
}

type Pprof struct {
	// Whether to pprof server.
	Enable bool

	// Address and port to listen.
	Listen string `validate:"hostname_port"`

	BasicAuth struct {
		Username string
		Password string
	} `koanf:"basic_auth"`

	// Turn on block profiling with the given rate.
	BlockProfileRate int `koanf:"block_profile_rate"`

	// Turn on memory profiling with the given rate.
	MemProfileRate int `koanf:"mem_profile_rate"`
}
