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

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		return logger
	}

	return Root()
}

func WithFields(ctx context.Context, fields ...interface{}) context.Context {
	logger := FromContext(ctx)
	return ToContext(ctx, logger.With(fields...))
}
