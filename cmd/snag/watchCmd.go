package main

import (
	"strings"

	"github.com/spf13/cobra"
)

func (app *application) initWatchCmd() {
	app.watchCmd = &cobra.Command{
		Use:   "watch",
		Short: "Watch a specific directory",
		Long: `Watches the directory recursively or
any specified directory, given that it is
passed as an argument.`,

		Run: app.watchFunc,
	}

	app.rootCmd.AddCommand(app.watchCmd)
}

// watchFunc executes the 'watch' subcommand
func (app *application) watchFunc(cmd *cobra.Command, args []string) {
	for _, dir := range args {
		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
		app.logger.Info(dir)
	}
}
