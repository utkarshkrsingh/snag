package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

func New(verbose bool) *Logger {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})

	return &Logger{
		Logger: slog.New(handler),
	}
}