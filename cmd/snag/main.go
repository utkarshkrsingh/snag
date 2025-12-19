package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/utkarhskrsingh/snag/internal/config"
	"github.com/utkarhskrsingh/snag/internal/logger"
)

type application struct {
	logger        logger.Logger
	verbose       bool
	rootCmd       *cobra.Command
	configCmd     *cobra.Command
	watchCmd      *cobra.Command
	runCmd        *cobra.Command
	configManager config.ConfigManager
}

func main() {

	cfgMgr := config.NewConfigManager()
	logger := logger.NewConsoleUI(false)

	app := newApplication(logger, cfgMgr)

	if err := app.rootCmd.Execute(); err != nil {
		app.logger.Error("Failed to start snag", err)
		os.Exit(1)
	}
}

func newApplication(logger logger.Logger, cfgMgr config.ConfigManager) *application {
	app := &application{
		logger:        logger,
		configManager: cfgMgr,
	}
	app.configManager = config.NewConfigManager()

	app.rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		app.logger = logger.NewConsoleUI(app.verbose)
	}
	app.initRootCmd()
	app.initConfigCmd()
	app.initWatchCmd()
	app.initRunCmd()

	return app
}
