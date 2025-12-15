package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd is the 'config' subcommand
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config file generation",
	Long: `Config commands generate a dummy config
that can be updated for the project further by the
developer.`,

	Run: configFunc,
}

func init() {
	configCmd.Flags().BoolP("init", "i", false, "Create a dummy config and then update according to your need.")
	configCmd.Flags().BoolP("check", "c", false, "Check whether the current config follows the expected format.")
	rootCmd.AddCommand(configCmd)
}

// configFunc executes 'config' subcommand
func configFunc(cmd *cobra.Command, args []string) {
	initFlag, err := cmd.Flags().GetBool("init")
	if err != nil {
		Logger.Fatal("Unable to read flags", err)
	}

	checkFlag, err := cmd.Flags().GetBool("check")
	if err != nil {
		Logger.Fatal("Unable to read flags", err)
	}

	if !initFlag && !checkFlag {
		Logger.Info("No flags provided. Use --help for available options.")
	}

	if initFlag {
		initConfig()
	}

	if checkFlag {
		checkConfig()
	}
}

func initConfig() {
	fmt.Println("init flag is used")
}

func checkConfig() {
	fmt.Println("check flag is used")
}
