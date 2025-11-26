package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	initFlag  bool
	checkFlag bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config file generation",
	Long: `Config commands generate a dummy config
that can be updated for the project further by the
developer.`,

	Run: configFunc,
}

func init() {
	configCmd.Flags().BoolVarP(&initFlag, "init", "i", false, "Create a dummy config and then update according to your need.")
	configCmd.Flags().BoolVarP(&checkFlag, "check", "c", false, "Check whether the current config follow the norms of not.")
	rootCmd.AddCommand(configCmd)
}

func configFunc(cmd *cobra.Command, args []string) {
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
