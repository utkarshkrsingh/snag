package main

import (
	"fmt"
	"os"

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

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		app.logger = logger.NewConsoleUI(app.verbose)
		if app.verbose {
			app.logger.Info("verbose mode is enabled")
		}

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		if err := app.cfgMgr.LoadConfig(cwd); err != nil && cmd.Name() != "config" {
			return fmt.Errorf(
				"snag.yaml not found or invalid (run `snag config`): %w",
				err,
			)
		}

		return nil
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd
}
