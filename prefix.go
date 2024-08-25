package slogcolor

import (
	"github.com/fatih/color"
)

// Prefix prepends a colored prefix to msg.
func Prefix(prefix string, msg string) string {
	return color.New(color.BgHiWhite, color.FgBlack).Sprint(prefix) + " " + msg
}
