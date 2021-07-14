package log

import (
	"fmt"
	"github.com/google/wire"
	"github.com/rollbar/rollbar-go"
	rb "github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(err error, fields ...zap.Field)
	Panic(err error, fields ...zap.Field)
	Fatal(err error, fields ...zap.Field)
}

type ZapLogger struct {
	Zap *zap.Logger
}

func NewZapLogger(r *rb.Rollbar) Logger {
	if r != nil {
		r.Config()
	}

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("can't initialize logger: %v\n", err)
	}
	defer func() { _ = logger.Sync() }()
	return &ZapLogger{Zap: logger}
}

func (l *ZapLogger) Debug(msg string, fields ...zap.Field) {
	l.Zap.Debug(msg, fields...)
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.Zap.Info(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...zap.Field) {
	l.Zap.Warn(msg, fields...)
}

func (l *ZapLogger) Error(err error, fields ...zap.Field) {
	rollbar.Error(err)
	l.Zap.Error(err.Error(), fields...)
}

func (l *ZapLogger) Panic(err error, fields ...zap.Field) {
	rollbar.Error(err)
	l.Zap.Panic(err.Error(), fields...)
}

func (l *ZapLogger) Fatal(err error, fields ...zap.Field) {
	rollbar.Error(err)
	l.Zap.Fatal(err.Error(), fields...)
}

var ProviderSet = wire.NewSet(NewZapLogger)
