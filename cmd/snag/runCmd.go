package main

import "github.com/spf13/cobra"

func (app *application) initRunCmd() {
	app.runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the script",
		Long:  "Run the script once",
		Run:   app.runFunc,
	}

	app.rootCmd.AddCommand(app.runCmd)
}

// runFunc executes the 'run' subcommand
func (app *application) runFunc(cmd *cobra.Command, args []string) {
	app.logger.Info("run flag is used")
}
