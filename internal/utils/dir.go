package utils

import (
	"os"

	"github.com/utkarhskrsingh/snag/internal/logger"
)

func GetCwD(logger *logger.ConsoleUI) string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatal("Unable to get current dir", err)
	}
	return dir
}
