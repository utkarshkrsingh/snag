package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

var (
	successColor = color.New(color.FgGreen)
	errorColor   = color.New(color.FgRed)
)

type ConsoleUI struct {
	logger  *slog.Logger
	verbose bool
}

func NewConsoleUI() *ConsoleUI {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))

	return &ConsoleUI{
		logger:  logger,
		verbose: false,
	}
}

func (c *ConsoleUI) Success(msg string) {
	successColor.Println("✔ ", msg)
}

func (c *ConsoleUI) Error(msg string, err error) {
	errorColor.Printf("✘ %s: %v\n", msg, err)
}

func (c *ConsoleUI) Info(msg string) {
	fmt.Println("•", msg)
}

func (c *ConsoleUI) Debug(msg string, keyAndValue ...any) {
	if !c.verbose {
		return
	}

	c.logger.Debug(msg, keyAndValue...)
}

func (c *ConsoleUI) Fatal(msg string, err error) {
	errorColor.Printf("✘ %s: %v\n", msg, err)
	os.Exit(-1)
}

func (c *ConsoleUI) SetVerbose(verbose bool) {
	if !c.verbose {
		c.verbose = true
	}
}
