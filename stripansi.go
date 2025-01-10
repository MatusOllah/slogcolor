package slogcolor

import (
	"bytes"
	"regexp"
)

// re is the regular expression used for removing ANSI colors.
var re = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")

// stripANSI removes ANSI escape sequences from the provided bytes.Buffer.
func stripANSI(bf *bytes.Buffer) {
	b := bf.Bytes()
	cleaned := re.ReplaceAll(b, nil)
	bf.Reset()
	bf.Write(cleaned)
}
