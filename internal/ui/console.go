package ui

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

var (
	success = color.New(color.FgGreen)
	fail    = color.New(color.FgRed)
)

type Console struct {
	out io.Writer
}

func New(out io.Writer) *Console {
	return &Console{
		out: out,
	}
}

func (c *Console) Info(msg string) {
	fmt.Fprintln(c.out, "•", msg)
}

func (c *Console) Success(msg string) {
	success.Fprintln(c.out, "✔", msg)
}

func (c *Console) Error(msg string, args ...any) {
	fail.Fprintln(c.out, "✘", msg, ":", args)
}
