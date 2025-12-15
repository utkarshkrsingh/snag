package main

import (
	"github.com/spf13/cobra"
)

// rootCmd is the root command for the snag application and
// is responsible for initiating all subcommands
var rootCmd = &cobra.Command{
	Use:   "snag",
	Short: "Snag is a very fast hot-reload tool.",
	Long: `Snag provides a fast and flexible way to hot-reload projects.
There is no need to manually build and run the application
after every modification.`,
}
