package main

import (
	"github.com/utkarhskrsingh/snag/internal/logger"
)

func init() {
	Logger = logger.NewConsoleUI()

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		Logger.Fatal("Failed to start snag", err)
	}
}
