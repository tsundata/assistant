package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func WorkflowCron(ctx rulebot.IContext) []result.Result {
	if ctx.Workflow() == nil {
		return []result.Result{result.EmptyResult()}
	}
	ctxB := context.Background()
	_, err := ctx.Workflow().CronTrigger(ctxB, &pb.TriggerRequest{})
	if err != nil {
		ctx.GetLogger().Error(err)
		return []result.Result{result.ErrorResult(err)}
	}
	return []result.Result{result.EmptyResult()}
}
