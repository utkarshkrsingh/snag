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

type Logger interface {
	Info(msg string)
	Error(msg string, err error)
	Debug(msg string, args ...any)
	Success(msg string)
}

type ConsoleUI struct {
	logger *slog.Logger
}

func NewConsoleUI(verbose bool) *ConsoleUI {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	return &ConsoleUI{
		logger: slog.New(slog.NewTextHandler(os.Stderr, opts)),
	}
}

func (c *ConsoleUI) Info(msg string) {
	fmt.Println("•", msg)
}

func (c *ConsoleUI) Error(msg string, err error) {
	errorColor.Printf("✘ %s: %v\n", msg, err)
}

func (c *ConsoleUI) Debug(msg string, args ...any) {
	c.logger.Debug(msg, args...)
}

func (c *ConsoleUI) Success(msg string) {
	successColor.Println("✔ ", msg)
}
