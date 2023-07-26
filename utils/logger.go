package utils

import "go.uber.org/zap"

var apiLogger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if apiLogger == nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync() // flushes buffer, if any
		sugar := logger.Sugar()
		apiLogger = sugar
		return apiLogger
	}

	return apiLogger
}
