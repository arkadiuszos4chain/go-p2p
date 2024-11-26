package p2p_logger

import (
	"context"
	"log/slog"
)

type ExtendedLogLevel slog.Level

const (
	LevelTrace ExtendedLogLevel = ExtendedLogLevel(slog.Level(-8))
	LevelDebug ExtendedLogLevel = ExtendedLogLevel(slog.LevelDebug)
	LevelInfo  ExtendedLogLevel = ExtendedLogLevel(slog.LevelInfo)
	LevelWarn  ExtendedLogLevel = ExtendedLogLevel(slog.LevelWarn)
	LevelError ExtendedLogLevel = ExtendedLogLevel(slog.LevelError)
)

type ExtendedLogger struct {
	slog.Logger
}

func FromLogger(l *slog.Logger, lvl ExtendedLogLevel) *ExtendedLogger {
	c := *l

	l.With(slog.Level(lvl))
}

// Trace logs at LevelTrace.
func (l *ExtendedLogger) Trace(msg string, args ...any) {
	l.Log(context.Background(), slog.Level(LevelTrace), msg, args...)
}

// TraceContext logs at LevelTrace with the given context.
func (l *ExtendedLogger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, slog.Level(LevelTrace), msg, args...)
}

func (l *ExtendedLogger) With(args ...any) *ExtendedLogger {
	c := &ExtendedLogger{}
	c.Logger = *l.Logger.With(args...)
	return c
}

func (l *ExtendedLogger) WithGroup(name string) *ExtendedLogger {
	c := &ExtendedLogger{}
	c.Logger = *l.Logger.WithGroup(name)
	return c
}
