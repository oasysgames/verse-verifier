package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/version"
	"github.com/spf13/cobra"
)

const (
	commandName = "oasvlfy"
	configFlag  = "config"
)

var rootCmd = &cobra.Command{
	Use: commandName,
	Long: fmt.Sprintf(`Name:
  %s - The verifier of the Verse-Layer.

  Copyright 2022 Oasys | Blockchain for Games All Rights Reserved.
  
Version:
  %s`, commandName, version.SemVer()),
}

func init() {
	rootCmd.PersistentFlags().String(configFlag, "", "configuration file")
	rootCmd.MarkFlagRequired(configFlag)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func loadConfig(cmd *cobra.Command) (*config.Config, error) {
	path, err := cmd.Flags().GetString(configFlag)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c, err := config.NewConfig(buf)
	if err != nil {
		return nil, err
	}

	return c, nil
}
