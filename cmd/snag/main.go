package main

import (
	"log"
)

func init() {
	// Remove auto-generated completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[ERR]: %v\n", err)
	}
}
