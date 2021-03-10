package logger

import (
	"github.com/rollbar/rollbar-go"
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	Zap *zap.Logger
}

func NewLogger() *Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()
	return &Logger{Zap: logger}
}

func (l *Logger) Error(err error, fields ...zap.Field) {
	rollbar.Error(err)
	l.Zap.Error(err.Error(), fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Zap.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Zap.Warn(msg, fields...)
}
