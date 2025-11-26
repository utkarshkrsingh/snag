package color

import "github.com/fatih/color"

var completion = color.New(color.FgGreen)
var warning = color.New(color.FgYellow)
var error = color.New(color.FgRed)

func CompletionCoding(message string) string {
	return completion.Sprint(message)
}

func WarningCoding(message string) string {
	return warning.Sprint(message)
}

func ErrorCoding(message string) string {
	return error.Sprint(message)
}
