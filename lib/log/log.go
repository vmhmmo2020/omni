// Package log provides a wrapper around the slog package (might change implementation later).
// It provides an opinionated interface for logging structured data always with context.
package log

import (
	"context"
	"log/slog"
	"runtime"
	"strings"
	"time"

	pkgerrors "github.com/pkg/errors" //nolint:revive // Need this for stacktraces.
)

type attrsKey struct{}

// WithCtx returns a copy of the context with which the logging attributes are associated.
// Usage:
//
//	ctx := log.WithCtx(ctx, "height", 1234)
//	...
//	log.Info(ctx, "Height processed") // Will contain attribute: height=1234
func WithCtx(ctx context.Context, attrs ...any) context.Context {
	return context.WithValue(ctx, attrsKey{}, mergeAttrs(ctx, attrs))
}

// Debug logs the message and attributes at default level.
func Debug(ctx context.Context, msg string, attrs ...any) {
	log(ctx, slog.LevelDebug, msg, mergeAttrs(ctx, attrs)...)
}

// Info logs the message and attributes at info level.
func Info(ctx context.Context, msg string, attrs ...any) {
	log(ctx, slog.LevelInfo, msg, mergeAttrs(ctx, attrs)...)
}

// Warn logs the message and error and attributes at warning level.
// If err is nil, it will not be logged.
func Warn(ctx context.Context, msg string, err error, attrs ...any) {
	if err != nil {
		attrs = append(attrs, slog.String("err", err.Error()))
		attrs = append(attrs, errAttrs(err)...)
	}

	log(ctx, slog.LevelWarn, msg, mergeAttrs(ctx, attrs)...)
}

// Error logs the message and error and arguments at error level.
// If err is nil, it will not be logged.
func Error(ctx context.Context, msg string, err error, attrs ...any) {
	if err != nil {
		attrs = append(attrs, slog.String("err", err.Error()))
		attrs = append(attrs, errAttrs(err)...)
	}
	log(ctx, slog.LevelError, msg, mergeAttrs(ctx, attrs)...)
}

// log is the low-level logging method for methods that take ...any.
// It must always be called directly by an exported logging method
// or function, because it uses a fixed call depth to obtain the pc.
//
// Copied from stdlib since we wrap slog and the source/caller is incorrect otherwise.
// See https://github.com/golang/go/blob/master/src/log/slog/logger.go#L241
func log(ctx context.Context, level slog.Level, msg string, attrs ...any) {
	logTotal.WithLabelValues(strings.ToLower(level.String())).Inc()

	logger := getLogger(ctx)

	if !logger.Enabled(ctx, level) {
		return
	}

	// skip [runtime.Callers, this function, this function's caller]
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(attrs...)

	_ = logger.Handler().Handle(ctx, r)
}

// errFields is similar to z.Err and returns the structured error fields and
// stack trace but without the error message. It avoids duplication of the error message
// since it is used as the main log message in Error above.
func errAttrs(err error) []any {
	type structErr interface {
		StackTrace() pkgerrors.StackTrace
		Attrs() []any
	}

	var attrs []any
	var stackTrace pkgerrors.StackTrace

	// Go up the cause chain (from the outermost error to the innermost error)
	for {
		if serr, ok := err.(structErr); ok { //nolint:errorlint // Using cast instead of errors.As since we do custom logic
			// Use the first encountered structErr's attributes.
			if len(attrs) == 0 {
				attrs = serr.Attrs()
			}

			// Use the last encountered structErr's stack trace.
			stackTrace = serr.StackTrace()
		}

		if cause := pkgerrors.Unwrap(err); cause != nil {
			err = cause
			continue // Continue up the cause chain.
		}

		// Root cause reached, break the loop.
		if len(stackTrace) > 0 {
			attrs = append(attrs, slog.Any("stacktrace", stackTrace))
		}

		return attrs
	}
}

// mergeAttrs returns the attributes from the context merged with the provided attributes.
func mergeAttrs(ctx context.Context, attrs []any) []any {
	resp, _ := ctx.Value(attrsKey{}).([]any) //nolint:revive // We know the type.
	resp = append(resp, attrs...)

	return resp
}
