package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/mdobak/go-xerrors"
)

var _ Logger = (*SlogLogger)(nil)

type SlogLogger struct{}

type stackFrame struct {
	Fn   string `json:"fn"`
	Line int    `json:"line"`
	Src  string `json:"src"`
}

func replaceAttr(_ []string, a slog.Attr) slog.Attr {
	switch a.Value.Kind() {
	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = fmtErr(v)
		}
	}

	return a
}

// marshalStack extracts stack frames from the error
func marshalStack(err error) []stackFrame {
	trace := xerrors.StackTrace(err)

	if len(trace) == 0 {
		return nil
	}

	frames := trace.Frames()

	s := make([]stackFrame, len(frames))

	for i, v := range frames {
		f := stackFrame{
			Fn:   filepath.Base(v.Function),
			Line: v.Line,
			Src: filepath.Join(
				filepath.Base(filepath.Dir(v.File)),
				filepath.Base(v.File),
			),
		}

		s[i] = f
	}

	return s
}

// fmtErr returns a slog.Value with keys `msg` and `trace`. If the error
// does not implement interface { StackTrace() errors.StackTrace }, the `trace`
// key is omitted.
func fmtErr(err error) slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	frames := marshalStack(err)

	if frames != nil {
		groupValues = append(groupValues,
			slog.Any("trace", frames),
		)
	}

	return slog.GroupValue(groupValues...)
}

func NewSlogLogger() *SlogLogger {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
	})
	slog.SetDefault(slog.New(h))
	return &SlogLogger{}
}

func (sloglogger *SlogLogger) Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func (sloglogger *SlogLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	slog.DebugContext(ctx, msg, args...)
}

func (sloglogger *SlogLogger) Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func (sloglogger *SlogLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
}

func (sloglogger *SlogLogger) Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func (sloglogger *SlogLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	slog.InfoContext(ctx, msg, args...)
}

func (sloglogger *SlogLogger) Log(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	slog.Log(ctx, level.Level(), msg, args...)
}

func (sloglogger *SlogLogger) LogAttrs(ctx context.Context, level slog.Leveler, msg string, args ...any) {
	panic("not implemented") // TODO: Implement
}

func (sloglogger *SlogLogger) Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func (sloglogger *SlogLogger) WarnContext(ctx context.Context, msg string, args ...any) {
	slog.WarnContext(ctx, msg, args...)
}
