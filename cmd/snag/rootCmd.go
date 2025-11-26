package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snag",
	Short: "Snag is a very fast hot-reload tool.",
	Long: `A fast and flexible way to hot-reload the project.
No need to manually build and run the project over and over
after modification.`,
}
