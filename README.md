# 🌈 slogcolor

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/slogcolor.svg)](https://pkg.go.dev/github.com/MatusOllah/slogcolor) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/slogcolor)](https://goreportcard.com/report/github.com/MatusOllah/slogcolor) [![GitHub license](https://img.shields.io/github/license/MatusOllah/slogcolor)](https://github.com/MatusOllah/slogcolor/blob/main/LICENSE)

![screenshot](https://github.com/MatusOllah/slogcolor/blob/main/screenshot.png)

**slogcolor** is a color handler for `log/slog`. It's output is inspired by XMRig and zerolog.

## Basic Usage

```go
package main

import (
    "os"
    "time"
    "errors"
    "log/slog"

    "github.com/MatusOllah/slogcolor"
)

func main() {
    slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

    slog.Info("Initializing")
    slog.Debug("Init done", "duration", 500*time.Millisecond)
    slog.Warn("Slow request!", "method", "GET", "path", "/api/users", "duration", 750*time.Millisecond)
    slog.Error("DB connection lost!", "err", errors.New("connection reset"), "db", "horalky")
}
```

### Customized output format

The output format can also be customized using the [`Options`](https://pkg.go.dev/github.com/MatusOllah/slogcolor#Options) like this:

```go
opts := &slogcolor.Options{
    Level:         slog.LevelDebug,
    TimeFormat:    time.RFC3339,
    SrcFileMode:   slog.Nop,
    SrcFileLength: 16,
}
slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, opts)))
```

### Prefixes

Messages can be prefixed with [`Prefix`](https://pkg.go.dev/github.com/MatusOllah/slogcolor#Prefix) like this:

```go
slog.Info(slogcolor.Prefix("MyPrefix", "kajšmentke"))
slog.Info(slogcolor.Prefix("SceneController", "switching scene"))
```

### Disable colors

Colors are enabled by default but can be disabled using `Options.NoColor`. Particularly useful for automatically enabling colors based on based on the terminal capabilities using e.g. the [`go-isatty`](https://github.com/mattn/go-isatty) package.

```go
w := os.Stderr

opts := slogcolor.DefaultOptions
opts.NoColor = !isatty.IsTerminal(w.Fd())

slog.SetDefault(slog.New(slogcolor.NewHandler(w, opts)))
```
