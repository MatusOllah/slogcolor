package slogcolor

import (
	"log/slog"
	"time"
)

var DefaultOptions *Options = &Options{
	Level:       slog.LevelInfo,
	TimeFormat:  time.DateTime,
	SrcFileMode: ShortFile,
}

type Options struct {
	// Level reports the minimum level to log.
	// Levels with lower levels are discarded.
	// If nil, the Handler uses [slog.LevelInfo].
	Level slog.Leveler

	// The time format.
	TimeFormat string

	// SrcFileMode is the source file mode.
	SrcFileMode SourceFileMode
}
