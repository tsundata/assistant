package stage

import (
	"context"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"go.uber.org/zap"
)

func Error(_ context.Context, comp component.Component, in result.Result) result.Result {
	if in.Kind == result.Error {
		if comp.GetLogger() == nil {
			return result.EmptyResult()
		}
		if err, ok := in.Content.(error); ok {
			comp.GetLogger().Error(err, zap.Any("in", in))
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
