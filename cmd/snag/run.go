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
			if len(args) == 0 {
				app.ui.Info("runs the whole script once")
				return nil
			}
			app.ui.Info("Running command: " + args[0])
			return nil
		},
	}

	return runCmd
}
