package log

import (
	"fmt"
	"github.com/google/wire"
	"github.com/rollbar/rollbar-go"
	rb "github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
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

func (l *ZapLogger) Debug(msg string, fields ...interface{}) {
	kvs := zapFields(fields)
	l.Zap.Debug(msg, kvs...)
}

func (l *ZapLogger) Info(msg string, fields ...interface{}) {
	kvs := zapFields(fields)
	l.Zap.Info(msg, kvs...)
}

func (l *ZapLogger) Warn(msg string, fields ...interface{}) {
	kvs := zapFields(fields)
	l.Zap.Warn(msg, kvs...)
}

func (l *ZapLogger) Error(err error, fields ...interface{}) {
	kvs := zapFields(fields)
	rollbar.Error(err)
	l.Zap.Error(err.Error(), kvs...)
}

func (l *ZapLogger) Panic(err error, fields ...interface{}) {
	kvs := zapFields(fields)
	rollbar.Error(err)
	l.Zap.Panic(err.Error(), kvs...)
}

func (l *ZapLogger) Fatal(err error, fields ...interface{}) {
	kvs := zapFields(fields)
	rollbar.Error(err)
	l.Zap.Fatal(err.Error(), kvs...)
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

var ProviderSet = wire.NewSet(NewZapLogger)
