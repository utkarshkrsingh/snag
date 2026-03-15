package main

import "os"

func main() {
	app := newApplication()

	if err := app.rootCmd.Execute(); err != nil {
		app.ui.Error("Failed to start snag", err)
		os.Exit(1)
	}
}
