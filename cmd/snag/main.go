package main

import (
	"github.com/spf13/cobra"
	"github.com/utkarhskrsingh/snag/internal/logger"
)

type application struct {
	logger    *logger.ConsoleUI
	rootCmd   *cobra.Command
	configCmd *cobra.Command
	watchCmd  *cobra.Command
	runCmd    *cobra.Command
}

func main() {

	app := newApplication()

	if err := app.rootCmd.Execute(); err != nil {
		app.logger.Fatal("Failed to start snag", err)
	}
}

func newApplication() *application {
	app := &application{
		logger: logger.NewConsoleUI(),
	}

	app.initRootCmd()
	app.initConfigCmd()
	app.initWatchCmd()
	app.initRunCmd()

	return app
}
