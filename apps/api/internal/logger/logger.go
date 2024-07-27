package logger

import (
	"context"
	"log/slog"
	"sync"
)

var (
	globalLogger Logger
	once         sync.Once
)

// init automatically initializes the global logger when the package is imported
func init() {
	once.Do(func() {
		globalLogger = LoggerFactory(Default)
	})
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() Logger {
	return globalLogger
}

// SetGlobalLogger allows changing the global logger after initialization if needed
func SetGlobalLogger(logger Logger) {
	globalLogger = logger
}

// Global logger methods for convenience

func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	globalLogger.DebugContext(ctx, msg, args...)
}

func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	globalLogger.ErrorContext(ctx, msg, args...)
}

func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	globalLogger.InfoContext(ctx, msg, args...)
}

func Log(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	globalLogger.Log(ctx, level, msg, args...)
}

func LogAttrs(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	globalLogger.LogAttrs(ctx, level, msg, args...)
}

func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	globalLogger.WarnContext(ctx, msg, args...)
}
