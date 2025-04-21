package slogcolor

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// Handler is a colored slog handler.
type Handler struct {
	groups []string
	attrs  []slog.Attr

	opts Options

	mu  *sync.Mutex
	out io.Writer
}

// NewHandler creates a new [Handler] with the specified options. If opts is nil, uses [DefaultOptions].
func NewHandler(out io.Writer, opts *Options) *Handler {
	h := &Handler{out: out, mu: &sync.Mutex{}}
	if opts == nil {
		h.opts = *DefaultOptions
	} else {
		h.opts = *opts
	}

	tags := maps.Clone(DefaultLevelTags)
	if opts.LevelTags != nil {
		for k, v := range opts.LevelTags {
			tags[k] = v
		}
	}
	h.opts.LevelTags = tags

	return h
}

func (h *Handler) clone() *Handler {
	return &Handler{
		groups: h.groups,
		attrs:  h.attrs,
		opts:   h.opts,
		mu:     h.mu,
		out:    h.out,
	}
}

// Enabled implements slog.Handler.Enabled .
func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

// Handle implements slog.Handler.Handle .
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	bf := getBuffer()
	bf.Reset()

	if !r.Time.IsZero() {
		fmt.Fprint(bf, color.New(color.Faint).Sprint(r.Time.Format(h.opts.TimeFormat)))
		fmt.Fprint(bf, " ")
	}

	fmt.Fprint(bf, h.opts.LevelTags[r.Level])

	fmt.Fprint(bf, " ")

	if h.opts.SrcFileMode != Nop {
		if r.PC != 0 {
			f, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()

			var filename string
			switch h.opts.SrcFileMode {
			case Nop:
				break
			case ShortFile:
				filename = filepath.Base(f.File)
			case LongFile:
				filename = f.File
			}
			lineStr := fmt.Sprintf(":%d", f.Line)
			formatted := fmt.Sprintf("%s ", filename+lineStr)
			if h.opts.SrcFileLength > 0 {
				maxFilenameLen := h.opts.SrcFileLength - len(lineStr) - 1
				if len(filename) > maxFilenameLen {
					filename = filename[:maxFilenameLen] // Truncate if too long
				}
				lenStr := strconv.Itoa(h.opts.SrcFileLength)
				formatted = fmt.Sprintf("%-"+lenStr+"s", filename+lineStr)
			}
			fmt.Fprint(bf, formatted)
		}
	}

	// we need the attributes here, as we can print a longer string if there are no attributes
	var attrs []slog.Attr
	attrs = append(attrs, h.attrs...)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	fmt.Fprint(bf, h.opts.MsgPrefix)
	formattedMessage := r.Message
	if h.opts.MsgLength > 0 && len(attrs) > 0 {
		if len(formattedMessage) > h.opts.MsgLength {
			formattedMessage = formattedMessage[:h.opts.MsgLength-1] + "…" // Truncate and add ellipsis if too long
		} else {
			// Pad with spaces if too short
			lenStr := strconv.Itoa(h.opts.MsgLength)
			formattedMessage = fmt.Sprintf("%-"+lenStr+"s", formattedMessage)
		}
	}
	if h.opts.MsgColor == nil {
		h.opts.MsgColor = color.New() // set to empty otherwise we have a null pointer
	}
	fmt.Fprintf(bf, "%s", h.opts.MsgColor.Sprint(formattedMessage))

	for _, a := range attrs {
		fmt.Fprint(bf, " ")
		for i, g := range h.groups {
			fmt.Fprint(bf, color.New(color.FgCyan).Sprint(g))
			if i != len(h.groups) {
				fmt.Fprint(bf, color.New(color.FgCyan).Sprint("."))
			}
		}

		if strings.Contains(a.Key, "err") {
			fmt.Fprint(bf, color.New(color.FgRed).Sprintf("%s=", a.Key)+a.Value.String())
		} else {
			fmt.Fprint(bf, color.New(color.FgCyan).Sprintf("%s=", a.Key)+a.Value.String())
		}
	}

	fmt.Fprint(bf, "\n")

	if h.opts.NoColor {
		stripANSI(bf)
	}

	h.mu.Lock()
	_, err := io.Copy(h.out, bf)
	h.mu.Unlock()

	freeBuffer(bf)

	return err
}

// WithGroup implements slog.Handler.WithGroup .
func (h *Handler) WithGroup(name string) slog.Handler {
	h2 := h.clone()
	h2.groups = append(h2.groups, name)
	return h2
}

// WithAttrs implements slog.Handler.WithAttrs .
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := h.clone()
	h2.attrs = append(h2.attrs, attrs...)
	return h2
}
