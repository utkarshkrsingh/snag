package main

import (
	"strings"

	"github.com/spf13/cobra"
)

// watchCmd is the 'watch' subcommand
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches a specific directory",
	Long: `Watches the directory recursively or
any specified directory, given that it is
passed as an argument.`,

	Run: watchFunc,
}

func init() {
	rootCmd.AddCommand(watchCmd)
}

// watchFunc executes the 'watch' subcommand
func watchFunc(cmd *cobra.Command, args []string) {
	for _, dir := range args {
		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
		Logger.Info(dir)
	}
}
