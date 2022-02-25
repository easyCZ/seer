package log

import "go.uber.org/zap"

var (
	rootLogger *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewDevelopment()
	rootLogger = logger.Sugar()
}

func SetRootLogger(logger *zap.SugaredLogger) {
	rootLogger = logger
}

func Root() *zap.SugaredLogger {
	return rootLogger
}
