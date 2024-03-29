package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/cmd"
)

func main() {
	setupLogger()
	cmd.Execute()
}

func setupLogger() {
	level := log.LvlInfo
	if os.Getenv("DEBUG") != "" {
		level = log.LvlDebug
	}

	logHandler := log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	logHandler = log.LvlFilterHandler(level, logHandler)
	logHandler = log.CallerFileHandler(logHandler)
	log.Root().SetHandler(logHandler)
}
