package slogcolor

import (
	"github.com/fatih/color"
)

func Prefix(prefix string, msg string) string {
	return color.New(color.BgHiWhite, color.FgBlack).Sprint(prefix) + " " + msg
}
