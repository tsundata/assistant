package pipeline

import (
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/stage"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Workflow(b *rulebot.RuleBot, r result.Result) {
	for {
		switch r.Kind {
		case result.Done:
			r = stage.Done(b, r)
			return
		case result.Error:
			r = stage.Error(b, r)
		case result.Message:
			r = stage.Message(b, r)
		case result.Url:
			r = stage.URL(b, r)
		case result.Repos:
			r = stage.Repos(b, r)
		default:
			return
		}
	}
}
