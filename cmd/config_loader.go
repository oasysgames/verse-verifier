package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/spf13/cobra"
)

const (
	fileConfigFlag = "config"
)

type configLoader struct {
	cmd *cobra.Command

	// from file
	fromFile string

	// from cli
	fromCli bool
	cfg     *config.Config
	verse   struct {
		use       bool
		verse     config.Verse
		scc, l2oo string
	}
	verifier struct {
		use    bool
		wallet config.Wallet
	}
	submitter struct {
		use     bool
		wallet  config.Wallet
		targets []uint
	}
}

type argConfigFlagGroup struct {
	name  string
	flags map[string]func(name string)
}

func argConfigFlag[T any](
	ptr *T,
	flagSetFn func(ptr *T, name string, defVal T, help string),
	help string,
) func(name string) {
	return func(name string) {
		flagSetFn(ptr, name, *ptr, help)
	}
}

func mustNewConfigLoader(cmd *cobra.Command) *configLoader {
	cfg := config.MustNewDefaultConfig()
	opts := &configLoader{cmd: cmd, cfg: cfg}
	f := cmd.PersistentFlags()

	// add flag for load configuration from the file
	f.StringVar(&opts.fromFile, fileConfigFlag, "", "Load config from the file")

	// add flags for load configuration from command line arguments
	argGroups := []*argConfigFlagGroup{
		{
			name: "",
			flags: map[string]func(name string){
				"cli":       argConfigFlag(&opts.fromCli, f.BoolVar, "Load config from command line arguments"),
				"datastore": argConfigFlag(&cfg.Datastore, f.StringVar, "Datastore directory path"),
				"keystore":  argConfigFlag(&cfg.Keystore, f.StringVar, "Private keys directory path"),
			},
		},
		{
			name: "hub",
			flags: map[string]func(name string){
				"chain_id": argConfigFlag(&cfg.HubLayer.ChainID, f.Uint64Var, "Chain ID of the Hub-Layer"),
				"rpc":      argConfigFlag(&cfg.HubLayer.RPC, f.StringVar, "RPC of the Hub-Layer(HTTP or WebSocket)"),
			},
		},
		{
			name: "p2p",
			flags: map[string]func(name string){
				"listens":   argConfigFlag(&cfg.P2P.Listens, f.StringSliceVar, "libp2p multi-addresses to listen"),
				"bootnodes": argConfigFlag(&cfg.P2P.Bootnodes, f.StringSliceVar, "Initial node list"),
			},
		},
		{
			name: "verse",
			flags: map[string]func(name string){
				"":          argConfigFlag(&opts.verse.use, f.BoolVar, "Use static verse setting"),
				"chain_id":  argConfigFlag(&opts.verse.verse.ChainID, f.Uint64Var, "Chain ID of the Verse-Layer"),
				"rpc":       argConfigFlag(&opts.verse.verse.RPC, f.StringVar, "RPC of the Verse-Layer(HTTP or WebSocket)"),
				"scc":       argConfigFlag(&opts.verse.scc, f.StringVar, "Address of the StateCommitmentChain"),
				"l2oo":      argConfigFlag(&opts.verse.l2oo, f.StringVar, "Address of the L2OutputOracle"),
				"discovery": argConfigFlag(&cfg.VerseLayer.Discovery.Endpoint, f.StringVar, "URL of the Verse-Layer list json"),
			},
		},
		{
			name: "verifier",
			flags: map[string]func(name string){
				"":                argConfigFlag(&opts.verifier.use, f.BoolVar, "Enable the verifier feature"),
				"wallet.address":  argConfigFlag(&opts.verifier.wallet.Address, f.StringVar, "Address of the verifier wallet"),
				"wallet.password": argConfigFlag(&opts.verifier.wallet.Password, f.StringVar, "Password file of the verifier wallet"),
				"wallet.plain":    argConfigFlag(&opts.verifier.wallet.Plain, f.StringVar, "Plaintext private key of the verifier wallet"),
			},
		},
		{
			name: "submitter",
			flags: map[string]func(name string){
				"":                      argConfigFlag(&opts.submitter.use, f.BoolVar, "Enable the submitter feature"),
				"confirmations":         argConfigFlag(&cfg.Submitter.Confirmations, f.IntVar, "Number of confirmation blocks for transaction receipt"),
				"scc-verifier-address":  argConfigFlag(&cfg.Submitter.SCCVerifierAddress, f.StringVar, "Address of the OasysStateCommitmentChainVerifier contract.."),
				"l2oo-verifier-address": argConfigFlag(&cfg.Submitter.L2OOVerifierAddress, f.StringVar, "Address of the OasysL2OutputOracleVerifier contract"),
				"multicall-address":     argConfigFlag(&cfg.Submitter.MulticallAddress, f.StringVar, "Address of the Multicall contract"),
				"targets":               argConfigFlag(&opts.submitter.targets, f.UintSliceVar, "List of Chain IDs to submit"),
				"wallet.address":        argConfigFlag(&opts.submitter.wallet.Address, f.StringVar, "Address of the submitter wallet"),
				"wallet.password":       argConfigFlag(&opts.submitter.wallet.Password, f.StringVar, "Password file of the submitter wallet"),
				"wallet.plain":          argConfigFlag(&opts.submitter.wallet.Plain, f.StringVar, "Plaintext private key of the submitter wallet"),
			},
		},
	}
	for _, ag := range argGroups {
		parts := []string{fileConfigFlag}
		if ag.name != "" {
			parts = append(parts, ag.name)
		}
		for name, flagSetFn := range ag.flags {
			parts := parts
			if name != "" {
				parts = append(parts, name)
			}
			flagSetFn(strings.Join(parts, "."))
		}
	}

	return opts
}

func (opts *configLoader) load() (*config.Config, error) {
	// load config from the file
	if !opts.fromCli {
		path, err := opts.cmd.Flags().GetString(fileConfigFlag)
		if err != nil {
			return nil, err
		}

		input, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		conf, err := config.NewConfig(input)
		if err != nil {
			return nil, err
		}

		return conf, nil
	}

	// load config from command line arguments
	opts.cfg.Wallets = map[string]*config.Wallet{}
	if opts.verse.use {
		opts.verse.verse.L1Contracts = map[string]string{}
		if opts.verse.scc != "" {
			opts.verse.verse.L1Contracts[SCCName] = opts.verse.scc
		}
		if opts.verse.l2oo != "" {
			opts.verse.verse.L1Contracts[L2OOName] = opts.verse.l2oo
		}
		opts.cfg.VerseLayer.Directs = []*config.Verse{&opts.verse.verse}
	}

	if opts.verifier.use {
		opts.cfg.Wallets["verifier"] = &opts.verifier.wallet
		opts.cfg.Verifier.Enable = true
		opts.cfg.Verifier.Wallet = "verifier"
	}

	if opts.submitter.use {
		opts.cfg.Wallets["submitter"] = &opts.submitter.wallet
		opts.cfg.Submitter.Enable = true
		for _, chainID := range opts.submitter.targets {
			target := config.SubmitterTarget{ChainID: uint64(chainID), Wallet: "submitter"}
			opts.cfg.Submitter.Targets = append(opts.cfg.Submitter.Targets, &target)
		}
	}

	if err := config.Validate(opts.cfg); err != nil {
		return nil, err
	}

	return opts.cfg, nil
}
