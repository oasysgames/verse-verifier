package cmd

import (
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status",
	Long:  "Show status",
	Run: func(cmd *cobra.Command, args []string) {
		ipccmd.StatusCmd.Run(commandName)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
