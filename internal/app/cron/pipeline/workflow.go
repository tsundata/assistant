package pipeline

import (
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/stage"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Workflow(ctx rulebot.IContext, in result.Result) {
	for {
		switch in.Kind {
		case result.Done:
			in = stage.Done(ctx, in)
			return
		case result.Error:
			in = stage.Error(ctx, in)
		case result.Message:
			in = stage.Message(ctx, in)
		case result.Url:
			in = stage.URL(ctx, in)
		case result.Repos:
			in = stage.Repos(ctx, in)
		default:
			return
		}
	}
}
