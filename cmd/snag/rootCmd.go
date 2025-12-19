package main

import (
	"github.com/spf13/cobra"
	"github.com/utkarhskrsingh/snag/internal/logger"
)

func newRootCmd(app *application) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "snag",
		Short: "Snag is a very fast hot-reload tool",
		Long: `Snag provides a fast and flexible way to hot-reload projects.
There is no need to manually build and run the application
after every modification.`,
	}

	rootCmd.PersistentFlags().BoolVarP(
		&app.verbose,
		"verbose",
		"v",
		false,
		"enable verbose logging",
	)

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		app.logger = logger.NewConsoleUI(app.verbose)
		app.logger.Info("verbose mode is enabled")
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
