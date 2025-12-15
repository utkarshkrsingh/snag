package main

import "github.com/spf13/cobra"

// runCmd is the 'run' subcommand
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the script",
	Long:  "Runs the script once.",
	Run:   runFunc,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

// runFunc executes the 'run' subcommand
func runFunc(cmd *cobra.Command, args []string) {
	Logger.Info("run flag is used")
}
