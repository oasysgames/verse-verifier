package cmd

import (
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/spf13/cobra"
)

const (
	pingFlag = "peer"
)

var pingCmd = &cobra.Command{
	Use:   "p2p:ping",
	Short: "Send ping via P2P to specified peer",
	Long:  "Send ping via P2P to specified peer",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := loadConfig(cmd)
		if err != nil {
			util.Exit(1, "Failed to load configuration file: %s\n", err)
		}

		peerID, err := cmd.Flags().GetString(pingFlag)
		if err != nil {
			util.Exit(1, "Failed to read '%s' argument: %s\n", pingFlag, err)
		}
		ipccmd.PingCmd.Run(cmd.Context(), conf.IPC.Sockname, peerID)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	pingCmd.Flags().String(pingFlag, "", "Target peer id")
	pingCmd.MarkFlagRequired(pingFlag)
}
