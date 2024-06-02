package main

import (
	"os"

	"log/slog"

	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/cmd"
	"github.com/oasysgames/oasys-optimism-verifier/logger"
)

func main() {
	setupLogger()
	cmd.Execute()
}

func setupLogger() {
	level := slog.LevelInfo
	if os.Getenv("DEBUG") != "" {
		level = slog.LevelDebug
	}
	output := os.Stdout
	useColor := true

	var handler slog.Handler
	handler = log.NewTerminalHandlerWithLevel(output, level, useColor)
	handler = logger.NewCallerHandler(handler)

	log.SetDefault(log.NewLogger(handler))
}
