package slogcolor

type SourceFileLength int

const (
	// ShortFile produces only the filename (for example main.go:69).
	ShortFile SourceFileLength = iota

	// LongFile produces the full file path (for example /home/frajer/go/src/myapp/main.go:69).
	LongFile
)
