package stage

import (
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Error(ctx rulebot.IContext, in result.Result) result.Result {
	if in.Kind == result.Error {
		if ctx.GetLogger() == nil {
			return result.EmptyResult()
		}
		if err, ok := in.Content.(error); ok {
			ctx.GetLogger().Error(err)
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
