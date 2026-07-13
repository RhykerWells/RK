package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type Handler struct {
	slog.Handler
}

var (
	log  *slog.Logger
	once sync.Once
)

// Init initializes the global logger.
// It should only be called once, typically from main().
func Init(level slog.Leveler) {
	once.Do(func() {
		log = New(os.Stdout, level)
	})
}

// Logger returns the global logger.
func Logger() *slog.Logger {
	if log == nil {
		panic("logger.Init() has not been called")
	}

	return log
}

// Convenience wrappers.

func Debug(msg string, args ...any) {
	Logger().Debug(msg, args...)
}

func Info(msg string, args ...any) {
	Logger().Info(msg, args...)
}

func Warn(msg string, args ...any) {
	Logger().Warn(msg, args...)
}

func Error(msg string, args ...any) {
	Logger().Error(msg, args...)
}

// With returns a child logger with additional attributes.
func With(args ...any) *slog.Logger {
	return Logger().With(args...)
}

func New(w io.Writer, level slog.Leveler) *slog.Logger {
	base := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(
					"time",
					a.Value.Time().Format("02/01/2006 @ 15:04"),
				)
			}

			return a
		},
	})

	return slog.New(&Handler{
		Handler: base,
	})
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(
		slog.String(
			"stck",
			getCaller(),
		),
	)

	return h.Handler.Handle(ctx, r)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithAttrs(attrs),
	}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Handler: h.Handler.WithGroup(name),
	}
}

func getCaller() string {
	for i := 2; i < 20; i++ {
		pc, file, line, ok := runtime.Caller(i)

		if !ok {
			continue
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		name := fn.Name()

		if strings.Contains(name, "/internal/logger") ||
			strings.Contains(name, "log/slog") {
			continue
		}

		return fmt.Sprintf(
			"%s:%s:%d",
			shortenFunction(name),
			filepath.Base(file),
			line,
		)
	}

	return "unknown"
}

func shortenFunction(name string) string {
	const prefix = "github.com/RhykerWells/RK/backend/internal/"

	return strings.TrimPrefix(name, prefix)
}
