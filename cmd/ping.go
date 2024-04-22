package cmd

import (
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/spf13/cobra"
)

const (
	peerFlag           = "peer"
	forceHolePunchFlag = "force-holepunch"
)

var pingCmd = &cobra.Command{
	Use:   "p2p:ping",
	Short: "Send ping via P2P to specified peer",
	Long:  "Send ping via P2P to specified peer",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := globalConfigLoader.load()
		if err != nil {
			util.Exit(1, "Failed to load configuration: %s\n", err)
		}

		peerID, err := cmd.Flags().GetString(peerFlag)
		if err != nil {
			util.Exit(1, "Failed to read '%s' argument: %s\n", peerFlag, err)
		}

		holePunch, err := cmd.Flags().GetBool(forceHolePunchFlag)
		if err != nil {
			util.Exit(1, "Failed to read '%s' argument: %s\n", forceHolePunchFlag, err)
		}

		ipccmd.PingCmd.Run(cmd.Context(), conf.IPC.Sockname, peerID, holePunch)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	pingCmd.Flags().String(peerFlag, "", "Target peer id")
	pingCmd.MarkFlagRequired(peerFlag)

	pingCmd.Flags().Bool(forceHolePunchFlag, false, "Enforce hole punching")
}
