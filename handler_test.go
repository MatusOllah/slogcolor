package slogcolor_test

import (
	"errors"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/MatusOllah/slogcolor"
)

func Example() {
	slog.SetDefault(slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions)))

	slog.Info("Initializing")
	slog.Debug("Init done", "duration", 500*time.Millisecond)
	slog.Warn("Slow request!", "method", "GET", "path", "/api/users", "duration", 750*time.Millisecond)
	slog.Error("DB connection lost!", "err", errors.New("connection reset"), "db", "horalky")
	// Output:
}

func BenchmarkLog(b *testing.B) {
	b.StopTimer()
	l := slog.New(slogcolor.NewHandler(os.Stderr, slogcolor.DefaultOptions))

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		l.Info("benchmarking", "i", i)
		b.StopTimer()
	}
}
