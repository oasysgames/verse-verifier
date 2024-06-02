package logger

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"log/slog"
)

const (
	skipLevel = 3
)

func NewCallerHandler(parent slog.Handler) slog.Handler {
	return &CallerHandler{parent}
}

type CallerHandler struct {
	slog.Handler
}

func (h *CallerHandler) Handle(ctx context.Context, r slog.Record) error {
	_, file, line, ok := runtime.Caller(skipLevel)
	if ok {
		r.AddAttrs(slog.String("caller", fmt.Sprintf("%s:%d", filepath.Base(file), line)))
	}
	return h.Handler.Handle(ctx, r)
}

func (h *CallerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CallerHandler{h.Handler.WithAttrs(attrs)}
}
