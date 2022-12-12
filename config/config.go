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
		"p2p.publish_interval":                   5 * time.Minute,
		"verifier.interval":                      15 * time.Second,
		"verifier.concurrency":                   50,
		"verifier.block_limit":                   1000,
		"submitter.interval":                     15 * time.Second,
		"submitter.concurrency":                  50,
		"submitter.confirmations":                6,
		"submitter.gas_multiplier":               1.0,
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
	P2P p2p `json:"p2p"`

	// IPC configuration.
	IPC ipc `json:"ipc"`

	// Verifier  configuration.
	Verifier verifier `json:"verifier" mapstructure:"verifier"`

	// Submitter  configuration.
	Submitter submitter `json:"submitter" mapstructure:"submitter"`
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

type p2p struct {
	// Address and port to listen.
	Listen string `json:"listen" validate:"hostname_port"`

	// Interval to publish own signature status.
	PublishInterval time.Duration `json:"publish_interval" mapstructure:"publish_interval"`

	// Initial node list.
	Bootnodes []string `json:"bootnodes"`
}

type ipc struct {
	// Whether to enable worker.
	Enable bool `json:"enable"`
}

type verifier struct {
	// Enable to verifier.
	Enable bool `json:"enable"`

	// Name of the wallet to create signature.
	Wallet string `json:"wallet" validate:"required_if=Enable true"`

	// Interval for get block data.
	Interval time.Duration `json:"interval" mapstructure:"interval"`

	// Number of concurrent executions.
	Concurrency int `json:"concurrency"`

	// Number of blocks to event filter.
	BlockLimit int `json:"block_limit" mapstructure:"block_limit"`
}

type submitter struct {
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

	Targets []struct {
		// Chain ID of the Verse-Layer.
		ChainID uint64 `json:"chain_id"     mapstructure:"chain_id"     validate:"required"`

		// Name of the wallet to send transaction.
		Wallet string `json:"wallet" validate:"required"`
	} `json:"targets" validate:"dive"`
}
