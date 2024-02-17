# slogcolor

![screenshot](https://github.com/MatusOllah/slogcolor/blob/main/screenshot.png)

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/slogcolor.svg)](https://pkg.go.dev/github.com/MatusOllah/slogcolor) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/slogcolor)](https://goreportcard.com/report/github.com/MatusOllah/slogcolor)

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
