package main

import (
	"github.com/spf13/cobra"
	"github.com/utkarshkrsingh/snag/internal/log"
	"github.com/utkarshkrsingh/snag/internal/ui"
)

func newRootCmd(app *application) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "snag",
		Short: "Snag is a very fast hot-reload tool",
		Long: `Snag provides a fast and flexible way to hot-reload projects.
There is no need to manually build and run the application
after every modification.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			return newWatchCmd(app).RunE(cmd, args)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(
		&app.verbose,
		"verbose",
		"v",
		false,
		"enable verbose logging",
	)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		app.ui = ui.New(cmd.OutOrStdout())
		app.logger = log.New(app.verbose)

		if app.verbose {
			app.ui.Info("verbose mode enabled")
		}

		return nil
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
