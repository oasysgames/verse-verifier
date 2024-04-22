package cmd

import (
	"fmt"
	"os"

	"github.com/oasysgames/oasys-optimism-verifier/version"
	"github.com/spf13/cobra"
)

const (
	commandName = "oasvlfy"
)

var rootCmd = &cobra.Command{
	Use: commandName,
	Long: fmt.Sprintf(`Name:
  %s - The verifier of the Verse-Layer.

  Copyright 2022 Oasys | Blockchain for Games All Rights Reserved.
  
Version:
  %s`, commandName, version.SemVer()),
}

var globalConfigLoader *configLoader

func init() {
	globalConfigLoader = mustNewConfigLoader(rootCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
