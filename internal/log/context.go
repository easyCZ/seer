package log

import (
	"context"

	"go.uber.org/zap"
)

type loggerContextKeyType int

const (
	loggerContextKey loggerContextKeyType = iota
)

func ToContext(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func FromContext(ctx context.Context) *
