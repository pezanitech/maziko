package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

// Custom handler for concise logging format
type conciseHandler struct {
	level  slog.Level
	writer io.Writer
}

// Creates a concise formatted logger
func createConciseLogger(logLevel slog.Level) (*slog.Logger, string) {
	handler := &conciseHandler{
		level:  logLevel,
		writer: os.Stdout,
	}
	return slog.New(handler), "concise"
}

// Implements slog.Handler interface
func (h *conciseHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *conciseHandler) Handle(ctx context.Context, r slog.Record) error {
	// format time to be human readable
	timeStr := r.Time.Format("2006-01-02 15:04:05.000")

	// Get log level as uppercase string
	levelStr := r.Level.String()

	// Format the main message
	fmt.Fprintf(h.writer, "%s [%s] %s",
		timeStr,
		levelStr,
		r.Message,
	)

	// Add attributes but in a cleaner format
	if r.NumAttrs() > 0 {
		fmt.Fprint(h.writer, " - ")
		r.Attrs(func(attr slog.Attr) bool {
			// Skip internal slog attributes
			if attr.Key == "time" || attr.Key == "level" || attr.Key == "msg" {
				return true
			}
			fmt.Fprintf(h.writer, "%s: %v ",
				attr.Key,
				attr.Value.Any(),
			)
			return true
		})
	}

	fmt.Fprintln(h.writer)
	return nil
}

func (h *conciseHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h // not yet storing attrs
}

func (h *conciseHandler) WithGroup(name string) slog.Handler {
	return h // not yet supporting groups
}
