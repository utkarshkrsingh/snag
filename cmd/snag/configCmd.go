package main

import (
	"os"

	"github.com/spf13/cobra"
)

func initConfigCmd(app *application) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Config file generation",
		Long: `Config commands generate a dummy config
that can be updated for the project further by the
developer.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			if err := app.cfgMgr.GenerateConfig(cwd); err != nil {
				return err
			}
			return nil
		},
	}

	return configCmd
}
