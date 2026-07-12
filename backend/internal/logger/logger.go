package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
)

type Handler struct {
	slog.Handler
}

func New(w io.Writer, level slog.Leveler) *slog.Logger {
	base := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: level,

		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
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

		// Ignore logger internals
		if strings.Contains(name, "/internal/logger") ||
			strings.Contains(name, "log/slog") {
			continue
		}

		name = shortenFunction(name)

		return fmt.Sprintf(
			"%s:%s:%d",
			name,
			filepath.Base(file),
			line,
		)
	}

	return "unknown"
}

func shortenFunction(name string) string {
	const prefix = "github.com/RhykerWells/RK/backend/internal/"

	name = strings.TrimPrefix(name, prefix)

	return name
}