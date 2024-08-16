package logger

import (
	"context"
	"log/slog"
	"os"
)

var _ Logger = (*SlogLogger)(nil)

type SlogLogger struct{}

type Option func(*slog.HandlerOptions)

func WithLevel(level slog.Level) Option {
	return func(otps *slog.HandlerOptions) {
		otps.Level = level
	}
}

func NewSlogLogger(options ...Option) *SlogLogger {
	defaultOptions := &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
	}
	for _, opt := range options {
		opt(defaultOptions)
	}

	h := slog.NewJSONHandler(os.Stdout, defaultOptions)
	slog.SetDefault(slog.New(h))
	return &SlogLogger{}
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func (l *SlogLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	slog.DebugContext(ctx, msg, args...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func (l *SlogLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func (l *SlogLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}

func (l *SlogLogger) Log(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	slog.Log(ctx, level.Level(), msg, args...)
}

func (l *SlogLogger) LogAttrs(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	panic("not implemented") // TODO: Implement
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func (l *SlogLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}
