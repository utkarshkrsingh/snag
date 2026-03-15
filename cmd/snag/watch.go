package main

import "github.com/spf13/cobra"

func newWatchCmd(app *application) *cobra.Command {
	watchCmd := &cobra.Command{
		Use:   "watch",
		Short: "Watch project directories for changes",
		Long: `Watch the current project directory and all
subdirectories for file changes and automatically
reload the application.`,

		Example: `    snag watch
    snag watch .
    snag watch ./cmd/api`,

		RunE: func(cmd *cobra.Command, args []string) error {
			app.ui.Info("started watching")
			return nil
		},
	}

	return watchCmd
}
