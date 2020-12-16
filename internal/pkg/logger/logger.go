package logger

import (
	"go.uber.org/zap"
	"log"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()
	return logger
}
