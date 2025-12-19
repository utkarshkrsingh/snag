package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/utkarhskrsingh/snag/internal/logger"
)

type application struct {
	logger  logger.Logger
	rootCmd *cobra.Command
	verbose bool
}

func main() {
	logger := logger.NewConsoleUI(false)
	app := newApplication(logger)

	if err := app.rootCmd.Execute(); err != nil {
		app.logger.Error("Failed to start snag", err)
		os.Exit(1)
	}
}

func newApplication(log logger.Logger) *application {
	app := &application{
		logger: log,
	}

	root := newRootCmd(app)
	root.AddCommand(
		initConfigCmd(app),
	)

	app.rootCmd = root

	return app
}
