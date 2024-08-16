package logger

import (
	"context"
	"log/slog"
	"os"
	"sync/atomic"
	"time"

	"github.com/SocBongDev/soc-bong/internal/apperr"
)

var globalLogger atomic.Value

// init automatically initializes the global logger when the package is imported
func init() {
	SetGlobalLogger(Default)
}

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() Logger {
	return globalLogger.Load().(Logger)
}

// SetGlobalLogger allows changing the global logger after initialization if needed
func SetGlobalLogger(loggerType LoggerType, options ...Option) {
	var (
		logger Logger
		err    error
	)

	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "dev"
	}

	switch loggerType {
	case Default:
		logLevel := slog.LevelDebug
		if env == "prod" {
			logLevel = slog.LevelInfo
		}

		logger = NewSlogLogger(WithLevel(logLevel))
	case File:
		logger, err = NewFileLogger("app.log")
		if err != nil {
			Error("SetFileLogger err", "err", apperr.New(err))
			return
		}
	case FileAndCentralized:
		logger, err = NewCompositeLogger("app.log", NewCentralLogStore(), 5*time.Minute)
		if err != nil {
			Error("SetCompositeLogger err", "err", apperr.New(err))
			return
		}
	default:
		logger = NewSlogLogger()
	}
	globalLogger.Store(logger)
}

// Global logger methods for convenience

func Debug(msg string, args ...any) {
	GetGlobalLogger().Debug(msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	GetGlobalLogger().DebugContext(ctx, msg, args...)
}

func Error(msg string, args ...any) {
	GetGlobalLogger().Error(msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	GetGlobalLogger().ErrorContext(ctx, msg, args...)
}

func Info(msg string, args ...any) {
	GetGlobalLogger().Info(msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	GetGlobalLogger().InfoContext(ctx, msg, args...)
}

func Log(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	GetGlobalLogger().Log(ctx, level, msg, args...)
}

func LogAttrs(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	GetGlobalLogger().LogAttrs(ctx, level, msg, args...)
}

func Warn(msg string, args ...any) {
	GetGlobalLogger().Warn(msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	GetGlobalLogger().WarnContext(ctx, msg, args...)
}
