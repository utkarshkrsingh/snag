package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/utkarhskrsingh.snag/internal/color"
)

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

func watchFunc(cmd *cobra.Command, args []string) {
	for _, dir := range args {
		if string(dir[len(dir)-1]) != "/" {
			dir += "/"
		}
		fmt.Println("[SNAG]", color.CompletionCoding(dir))
	}
}
