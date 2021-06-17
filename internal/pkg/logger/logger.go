package logger

import (
	"fmt"
	"github.com/google/wire"
	"github.com/rollbar/rollbar-go"
	rb "github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"go.uber.org/zap"
)

type Logger struct {
	Zap *zap.Logger
}

func NewLogger(r *rb.Rollbar) *Logger {
	r.Config()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("can't initialize logger: %v\n", err)
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

var ProviderSet = wire.NewSet(NewLogger)
