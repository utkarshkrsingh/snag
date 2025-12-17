package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initConfigCmd registers the 'config' subcommand with the root command.
func (app *application) initConfigCmd() {
	app.configCmd = &cobra.Command{
		Use:   "config",
		Short: "Config file generation",
		Long: `Config commands generate a dummy config
that can be updated for the project further by the
developer.`,

		Run: app.configFunc,
	}

	app.configCmd.Flags().BoolP("init", "i", false, "Create a dummy config and then udate according to your need")
	app.configCmd.Flags().BoolP("check", "c", false, "Check whether the current config follows the expected format")

	app.rootCmd.AddCommand(app.configCmd)
}

// configFunc handles execution of the 'config' subcommand.
func (app *application) configFunc(cmd *cobra.Command, args []string) {
	initFlag, err := cmd.Flags().GetBool("init")
	if err != nil {
		app.logger.Fatal("Unable to read flags", err)
	}

	checkFlag, err := cmd.Flags().GetBool("check")
	if err != nil {
		app.logger.Fatal("Unable to read flags", err)
	}

	if !initFlag && !checkFlag {
		app.logger.Info("No flags provided. Use --help for available options.")
	}

	if initFlag {
		app.initConfig()
	}

	if checkFlag {
		app.checkConfig()
	}
}

func (app *application) initConfig() {
	fmt.Println("init flag is used")
}

func (app *application) checkConfig() {
	fmt.Println("check flag is used")
}
