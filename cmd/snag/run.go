package main

import "github.com/spf13/cobra"

func newRunCmd(app *application) *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run [command]",
		Short: "Run a command once",
		Long: `Run a configured command once without starting
the file watcher.

If a command name is provided, only that command
will be executed.`,

		Example: `	snag run
	snag run build
	snag run test`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.configMgr.Load(); err != nil {
				return err
			}
			app.ui.Info("config loaded")
			return nil
		},
	}

	return runCmd
}
