package pipeline

import (
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/stage"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Workflow(ctx rulebot.IContext, r result.Result) {
	for {
		switch r.Kind {
		case result.Done:
			r = stage.Done(ctx, r)
			return
		case result.Error:
			r = stage.Error(ctx, r)
		case result.Message:
			r = stage.Message(ctx, r)
		case result.Url:
			r = stage.URL(ctx, r)
		case result.Repos:
			r = stage.Repos(ctx, r)
		default:
			return
		}
	}
}
