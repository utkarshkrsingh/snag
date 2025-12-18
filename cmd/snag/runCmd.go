package main

import "github.com/spf13/cobra"

// initRunCmd registers the 'run' subcommand with the root command.
func (app *application) initRunCmd() {
	app.runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the script",
		Long:  "Run the script once",
		Run:   app.runFunc,
	}

	app.rootCmd.AddCommand(app.runCmd)
}

// runFunc handles execution of the 'run' subcommand.
func (app *application) runFunc(cmd *cobra.Command, args []string) {
	app.logger.Info("run flag is used")
}
