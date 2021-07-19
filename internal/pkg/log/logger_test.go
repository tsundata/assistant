package log

import (
	"errors"
	"go.uber.org/zap"
	"testing"
)

func TestLogger(t *testing.T) {
	z := NewZapLogger(nil)
	l := NewAppLogger(z)
	l.Error(errors.New("test error"), zap.Any("t", t.Name()))
	l.Debug("debug", zap.Any("t", t.Name()))
	l.Info("info", zap.Any("t", t.Name()))
	l.Warn("info", zap.Any("t", t.Name()))
}
