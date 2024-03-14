package cmd

import (
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status",
	Long:  "Show status",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := globalConfigLoader.load()
		if err != nil {
			util.Exit(1, "Failed to load configuration: %s\n", err)
		}
		ipccmd.StatusCmd.Run(conf.IPC.Sockname)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
