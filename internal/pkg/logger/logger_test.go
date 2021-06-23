package logger

import (
	"errors"
	"go.uber.org/zap"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewLogger(nil)
	l.Error(errors.New("test error"), zap.Any("t", t.Name()))
	l.Info("info", zap.Any("t", t.Name()))
	l.Warn("info", zap.Any("t", t.Name()))
}
