package cmd

import (
	"fmt"
	"syscall"

	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	nameFlag     = "name"
	passwordFlag = "password"
)

var unlockCmd = &cobra.Command{
	Use:   "wallet:unlock",
	Short: "Unlock the Wallet",
	Long:  "Unlock the Wallet",
	Run:   runUnlockCmd,
}

func init() {
	rootCmd.AddCommand(unlockCmd)

	unlockCmd.Flags().String(nameFlag, "", "wallet name")
	unlockCmd.MarkFlagRequired(nameFlag)

	unlockCmd.Flags().String(passwordFlag, "", "wallet password")
}

func runUnlockCmd(cmd *cobra.Command, args []string) {
	conf, err := globalConfigLoader.load()
	if err != nil {
		util.Exit(1, "Failed to load configuration: %s\n", err)
	}

	name, err := cmd.Flags().GetString(nameFlag)
	if err != nil {
		util.Exit(1, "Failed to read '%s' argument: %s\n", nameFlag, err)
	}

	wallet, ok := conf.Wallets[name]
	if !ok {
		util.Exit(1, "unknown wallet\n")
	}

	password, err := cmd.Flags().GetString(passwordFlag)
	if err != nil {
		util.Exit(1, "Failed to read '%s' argument: %s\n", passwordFlag, err)
	}

	if password == "" {
		fmt.Print("Password: ")
		input, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			util.Exit(1, "Failed to read password: %s\n", err)
		}

		fmt.Print("\n")
		password = string(input)
	}

	ipccmd.WalletUnlockCmd.Run(conf.IPC.Sockname, wallet.Address, password)
}
