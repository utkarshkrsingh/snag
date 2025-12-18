package main

import (
	"github.com/spf13/cobra"
)

// initRootCmd configures the root Cobra command for the application.
func (app *application) initRootCmd() {
	app.rootCmd = &cobra.Command{
		Use:   "snag",
		Short: "Snag is a very fast hot-reload tool",
		Long: `Snag provides a fast and flexible way to hot-reload projects.
There is no need to manually build and run the application
after every modification.`,
	}

	app.rootCmd.PersistentFlags().BoolVarP(
		&app.verbose,
		"verbose",
		"v",
		false,
		"enable verbose logging",
	)

	app.rootCmd.CompletionOptions.DisableDefaultCmd = true
}
