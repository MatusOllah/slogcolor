package slogcolor

import (
	"log/slog"
	"time"
)

var DefaultOptions *Options = &Options{
	Level:         slog.LevelInfo,
	TimeFormat:    time.DateTime,
	AddSource:     true,
	SrcFileLength: ShortFile,
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler

	// The time format.
	TimeFormat string

	// Enable source code location / trace.
	AddSource bool

	// SrcFileLength is the source file length.
	SrcFileLength SourceFileLength
}
