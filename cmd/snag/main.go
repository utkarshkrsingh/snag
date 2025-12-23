package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/utkarhskrsingh/snag/internal/config"
	"github.com/utkarhskrsingh/snag/internal/logger"
)

type application struct {
	logger  logger.Logger
	rootCmd *cobra.Command
	verbose bool
	cfgMgr  config.CfgMgr
}

func main() {
	logger := logger.NewConsoleUI(false)
	cfgMgr := config.NewConfigManager()

	app := newApplication(logger, cfgMgr)

	if err := app.rootCmd.Execute(); err != nil {
		app.logger.Error("Failed to start snag", err)
		os.Exit(1)
	}
}

func newApplication(log logger.Logger, cfgMgr config.CfgMgr) *application {
	app := &application{
		logger: log,
		cfgMgr: cfgMgr,
	}

	root := newRootCmd(app)
	root.AddCommand(
		initConfigCmd(app),
	)

	app.rootCmd = root

	return app
}
