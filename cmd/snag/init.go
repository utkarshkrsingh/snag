package main

import "github.com/spf13/cobra"

func newInitCmd(app *application) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a configuration file",
		Long: `Create a default snag.yaml configuration file
in the current project directory.`,

		Example: `	snag init`,

		RunE: func(cmd *cobra.Command, args []string) error {
			app.ui.Success("config initiated")
			return nil
		},
	}

	return initCmd
}
