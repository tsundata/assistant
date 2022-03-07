package pipeline

import (
	"context"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/stage"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

func Workflow(ctx context.Context, comp component.Component, in result.Result) {
	for {
		switch in.Kind {
		case result.Done:
			in = stage.Done(ctx, comp, in)
			return
		case result.Error:
			in = stage.Error(ctx, comp, in)
		case result.Message:
			in = stage.Message(ctx, comp, in)
		case result.Url:
			in = stage.URL(ctx, comp, in)
		case result.Repos:
			in = stage.Repos(ctx, comp, in)
		default:
			return
		}
	}
}
