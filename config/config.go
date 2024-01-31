package config

import (
	"bytes"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var (
	validate = validator.New()

	defaults = map[string]interface{}{
		"verse_layer.discovery.refresh_interval": time.Hour,

		"p2p.publish_interval": 5 * time.Minute,
		"p2p.stream_timeout":   15 * time.Second,
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
		"p2p.enable_upnp":          true,
		"p2p.enable_auto_nat":      true,
		"p2p.enable_hole_punching": true,

		"ipc.sockname": "oasvlfy",

		"verifier.interval":              15 * time.Second,
		"verifier.concurrency":           50,
		"verifier.block_limit":           1000,
		"verifier.event_filter_limit":    1000,
		"verifier.state_collect_limit":   1000,
		"verifier.state_collect_timeout": 15 * time.Second,
		"verifier.db_optimize_interval":  time.Hour,

		"submitter.interval":           15 * time.Second,
		"submitter.concurrency":        50,
		"submitter.confirmations":      6,
		"submitter.gas_multiplier":     1.1,
		"submitter.batch_size":         20,
		"submitter.max_gas":            5_000_000,
		"submitter.verifier_address":   "0x5200000000000000000000000000000000000014",
		"submitter.multicall2_address": "0x5200000000000000000000000000000000000022",

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
)

func init() {
	initViper()
	initValidate()
}

// Read the configuration file.
func NewConfig(input []byte) (*Config, error) {
	if err := viper.ReadConfig(bytes.NewBuffer(input)); err != nil {
		return nil, err
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	if err := validate.Struct(conf); err != nil {
		return nil, err
	}

	return &conf, nil

}

func initViper() {
	// set config types
	viper.SetConfigType("json")
	viper.SetConfigType("yaml")

	// set default values
	for k, b := range defaults {
		viper.SetDefault(k, b)
	}
}

func initValidate() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// use `json` tag
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// App configuration.
type Config struct {
	// Datastore directory path.
	DataStore string `json:"datastore" validate:"dir"`

	// Validator keystore directory path.
	KeyStore string `json:"keystore" validate:"dir"`

	// Address used to create signatures and send transactions.
	Wallets map[string]Wallet `json:"wallets" validate:"dive"`

	// Configuration of the Hub-Layer.
	HubLayer hubLayer `json:"hub_layer" mapstructure:"hub_layer"`

	// Configuration of Verse-Layer discovery.
	VerseLayer verseLayer `json:"verse_layer" mapstructure:"verse_layer"`

	// P2P worker configuration.
	P2P P2P `json:"p2p"`

	// IPC configuration.
	IPC IPC `json:"ipc"`

	// Verifier configuration.
	Verifier Verifier `json:"verifier" mapstructure:"verifier"`

	// Submitter configuration.
	Submitter Submitter `json:"submitter" mapstructure:"submitter"`

	// Beacon worker configuration.
	Beacon Beacon `json:"beacon" mapstructure:"beacon"`

	// Database configuration.
	Database Database `json:"database" mapstructure:"database"`

	// Metrics configuration
	Metrics Metrics `json:"metrics"`

	// Debug configuration.
	Debug Debug `json:"debug" mapstructure:"debug"`
}

func (c *Config) DatabasePath() string {
	return filepath.Join(c.DataStore, "db.sqlite")
}

func (c *Config) P2PKeyPath() string {
	return filepath.Join(c.DataStore, "p2p.key")
}

func (c *Config) OpenKeyStore() *keystore.KeyStore {
	return keystore.NewKeyStore(c.KeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
}

type Wallet struct {
	// Address of the wallet.
	Address string `json:"address" validate:"hexadecimal"`

	// Password file of the wallet.
	Password string `json:"password" validate:"omitempty,file"`
}

type hubLayer struct {
	// Chain ID of the Hub-Layer.
	ChainId uint64 `json:"chain_id" mapstructure:"chain_id" validate:"required"`

	// RPC of the Hub-Layer(HTTP or WebSocket).
	RPC string `json:"rpc" validate:"url"`
}

type Verse struct {
	// Chain ID of the Verse-Layer.
	ChainID uint64 `json:"chain_id" mapstructure:"chain_id" validate:"required"`

	// RPC of the Verse-Layer(HTTP or WebSocket).
	RPC string `json:"rpc" validate:"url"`

	// Contract addresses on the Hub-Layer.
	L1Contracts map[string]string `json:"l1_contracts" mapstructure:"l1_contracts" validate:"required,dive,hexadecimal"`
}

type verseLayer struct {
	// Discover from API.
	Discovery struct {
		Endpoint        string        `json:"endpoint" validate:"omitempty,url"`
		RefreshInterval time.Duration `json:"refresh_interval" mapstructure:"refresh_interval"`
	} `json:"discovery" validate:"dive"`

	// List of Verse-Layer.
	Directs []*Verse `json:"directs" validate:"dive"`
}

type P2P struct {
	// libp2p multi-addresses to listen.
	Listens []string `json:"listens" mapstructure:"listens"`

	// Additional multi-addresses to advertise.
	AppendAnnounce []string `json:"append_announce" mapstructure:"append_announce"`

	// Multi-addresses not advertised.
	NoAnnounce []string `json:"no_announce" mapstructure:"no_announce"`

	// Multi-addresses that filter dial or receive connections.
	ConnectionFilter []string `json:"connection_filter" mapstructure:"connection_filter"`

	// Deprecated: Address and port to listen.
	Listen string `json:"listen" validate:"omitempty,hostname_port"`

	// Interval to publish own signature status.
	PublishInterval time.Duration `json:"publish_interval" mapstructure:"publish_interval"`

	// Timeout for P2P stream communication.
	StreamTimeout time.Duration `json:"stream_timeout" mapstructure:"stream_timeout"`

	// Initial node list.
	Bootnodes []string `json:"bootnodes"`

	// Enable NAT traversal using UPnP.
	EnableUPnP bool `json:"enable_upnp" mapstructure:"enable_upnp"`

	// Enable AutoNAT service.
	EnableAutoNAT bool `json:"enable_auto_nat" mapstructure:"enable_auto_nat"`

	// Enable Circuit Relay(v2) service.
	// Note: Public connectivity is required.
	RelayService struct {
		Enable bool `json:"enable"`

		// DurationLimit is the limit of data relayed (on each direction) before resetting the connection.
		DurationLimit *time.Duration `json:"duration_limit,omitempty" mapstructure:"duration_limit"`
		// DataLimit is the time limit before resetting a relayed connection.
		DataLimit *int64 `json:"data_limit,omitempty" mapstructure:"data_limit"`

		// ReservationTTL is the duration of a new (or refreshed reservation).
		ReservationTTL *time.Duration `json:"reservation_ttl,omitempty" mapstructure:"reservation_ttl"`

		// MaxReservations is the maximum number of active relay slots.
		MaxReservations *int `json:"max_reservations,omitempty" mapstructure:"max_reservations"`
		// MaxCircuits is the maximum number of open relay connections for each peer; defaults to 16.
		MaxCircuits *int `json:"max_circuits,omitempty" mapstructure:"max_circuits"`
		// BufferSize is the size of the relayed connection buffers.
		BufferSize *int `json:"buffer_size,omitempty" mapstructure:"buffer_size"`

		// MaxReservationsPerPeer is the maximum number of reservations originating from the same peer.
		MaxReservationsPerPeer *int `json:"max_reservations_per_peer,omitempty" mapstructure:"max_reservations_per_peer"`
		// MaxReservationsPerIP is the maximum number of reservations originating from the same IP address.
		MaxReservationsPerIP *int `json:"max_reservations_per_ip,omitempty" mapstructure:"max_reservations_per_ip"`
		// MaxReservationsPerASN is the maximum number of reservations origination from the same ASN.
		MaxReservationsPerASN *int `json:"max_reservations_per_asn,omitempty" mapstructure:"max_reservations_per_asn"`
	} `json:"relay_service" mapstructure:"relay_service"`

	// Enable Circuit Relay(v2) client.
	RelayClient struct {
		Enable     bool     `json:"enable"`
		RelayNodes []string `json:"relay_nodes" mapstructure:"relay_nodes"`
	} `json:"relay_client" mapstructure:"relay_client"`

	// Enable NAT traversal using UDP hole punching.
	EnableHolePunching bool `json:"enable_hole_punching" mapstructure:"enable_hole_punching"`
}

type IPC struct {
	// Socket file name, In UNIX-based OS, it is created as /tmp/{sockname}.sock.
	Sockname string `json:"sockname"`
}

type Verifier struct {
	// Enable to verifier.
	Enable bool `json:"enable"`

	// Name of the wallet to create signature.
	Wallet string `json:"wallet" validate:"required_if=Enable true"`

	// Interval for get block data.
	Interval time.Duration `json:"interval" mapstructure:"interval"`

	// Number of concurrent executions.
	Concurrency int `json:"concurrency"`

	// Number of block headers to collect at a time.
	BlockLimit int `json:"block_limit" mapstructure:"block_limit"`

	// Number of blocks to event filter.
	EventFilterLimit int `json:"event_filter_limit" mapstructure:"event_filter_limit"`

	// Number of state root to collect at a time.
	StateCollectLimit int `json:"state_collect_limit" mapstructure:"state_collect_limit"`

	// Timeout for state root collection.
	StateCollectTimeout time.Duration `json:"state_collect_timeout" mapstructure:"state_collect_timeout"`

	// Interval to optimize database.
	OptimizeInterval time.Duration `json:"db_optimize_interval" mapstructure:"db_optimize_interval"`
}

type Submitter struct {
	// Whether to enable worker.
	Enable bool `json:"enable"`

	// Interval for send transaction.
	Interval time.Duration `json:"interval" mapstructure:"interval"`

	// Number of concurrent executions.
	Concurrency int `json:"concurrency"`

	// Number of confirmation blocks for transaction receipt.
	Confirmations int `json:"confirmations"`

	// How much to increase the estimated gas limit.
	GasMultiplier float64 `json:"gas_multiplier" mapstructure:"gas_multiplier"`

	// Maximum number of calls for Multicall2.
	BatchSize int `json:"batch_size" mapstructure:"batch_size"`

	// Maximum gas of calls for Multicall2.
	MaxGas int `json:"max_gas" mapstructure:"max_gas"`

	// Address of the OasysStateCommitmentChain contract.
	VerifierAddress string `json:"verifier_address" mapstructure:"verifier_address"`

	// Address of the Multicall2 contract.
	Multicall2Address string `json:"multicall2_address" mapstructure:"multicall2_address"`

	Targets []struct {
		// Chain ID of the Verse-Layer.
		ChainID uint64 `json:"chain_id"     mapstructure:"chain_id"     validate:"required"`

		// Name of the wallet to send transaction.
		Wallet string `json:"wallet" validate:"required"`
	} `json:"targets" validate:"dive"`
}

type Beacon struct {
	// Whether to enable worker.
	Enable bool `json:"enable"`

	// URL of beacon.
	Endpoint string `json:"endpoint" validate:"omitempty,url"`

	// Interval for send beacon.
	Interval time.Duration `json:"interval"`
}

type Database struct {
	// File path of the SQLite database.
	Path string `json:"path" mapstructure:"path"`

	// Slow query log configurations.
	LongQueryTime       time.Duration `json:"long_query_time"        mapstructure:"long_query_time"`
	MinExaminedRowLimit int           `json:"min_examined_row_limit" mapstructure:"min_examined_row_limit"`
}

type Metrics struct {
	// Whether to pprof server.
	Enable bool `json:"enable"`

	// Address and port to listen.
	Listen string `json:"listen" validate:"hostname_port"`

	// The URL used to retrieve metrics.
	Endpoint string `json:"endpoint"`

	// The type of metrics collector.
	Type string `json:"type"`

	// Metric name prefix.
	Prefix string `json:"prefix"`
}

type Debug struct {
	Pprof Pprof `json:"pprof" validate:"dive"`
}

type Pprof struct {
	// Whether to pprof server.
	Enable bool `json:"enable"`

	// Address and port to listen.
	Listen string `json:"listen" validate:"hostname_port"`

	BasicAuth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"basic_auth" mapstructure:"basic_auth"`

	// Turn on block profiling with the given rate.
	BlockProfileRate int `json:"block_profile_rate" mapstructure:"block_profile_rate"`

	// Turn on memory profiling with the given rate.
	MemProfileRate int `json:"mem_profile_rate" mapstructure:"mem_profile_rate"`
}
