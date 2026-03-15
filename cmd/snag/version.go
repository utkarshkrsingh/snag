package main

import "github.com/spf13/cobra"

func newVersionCmd(app *application) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show the Snag version",
		Long:  "Display the current version of Snag.",

		Example: `	snag version`,

		RunE: func(cmd *cobra.Command, args []string) error {
			app.ui.Info("Snag version v1.0.0")
			return nil
		},
	}

	return versionCmd
}
