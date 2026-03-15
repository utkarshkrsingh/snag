package main

import (
	"github.com/spf13/cobra"
	"github.com/utkarshkrsingh/snag/internal/log"
	"github.com/utkarshkrsingh/snag/internal/ui"
)

type application struct {
	logger  *log.Logger
	ui      *ui.Console
	rootCmd *cobra.Command
	verbose bool
}

func newApplication() *application {
	app := application{
		logger:  log.New(false),
		verbose: false,
	}
	app.rootCmd = newRootCmd(&app)
	app.rootCmd.AddCommand(newRunCmd(&app))
	app.rootCmd.AddCommand(newVersionCmd(&app))
	app.rootCmd.AddCommand(newInitCmd(&app))
	return &app
}
