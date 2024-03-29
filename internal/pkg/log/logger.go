package log

import (
	"fmt"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(err error, fields ...interface{})
	Panic(err error, fields ...interface{})
	Fatal(err error, fields ...interface{})
}

func NewZapLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("can't initialize logger: %v\n", err)
	}
	defer func() { _ = logger.Sync() }()
	return logger
}

type AppLogger struct {
	logger *zap.Logger
}

func NewAppLogger(zap *zap.Logger) Logger {
	return &AppLogger{logger: zap}
}

func (l *AppLogger) Debug(msg string, fields ...interface{}) {
	kvs := zapFields(fields)
	l.logger.Debug(msg, kvs...)
}

func (l *AppLogger) Info(msg string, fields ...interface{}) {
	kvs := zapFields(fields)
	l.logger.Info(msg, kvs...)
}

func (l *AppLogger) Warn(msg string, fields ...interface{}) {
	kvs := zapFields(fields)
	l.logger.Warn(msg, kvs...)
}

func (l *AppLogger) Error(err error, fields ...interface{}) {
	kvs := zapFields(fields)
	l.logger.Error(err.Error(), kvs...)
}

func (l *AppLogger) Panic(err error, fields ...interface{}) {
	kvs := zapFields(fields)
	l.logger.Panic(err.Error(), kvs...)
}

func (l *AppLogger) Fatal(err error, fields ...interface{}) {
	kvs := zapFields(fields)
	l.logger.Fatal(err.Error(), kvs...)
}

func zapFields(fields []interface{}) []zap.Field {
	var res []zap.Field
	for _, i := range fields {
		if f, ok := i.(zap.Field); ok {
			res = append(res, f)
		}
	}
	return res
}

var ProviderSet = wire.NewSet(NewZapLogger, NewAppLogger)
