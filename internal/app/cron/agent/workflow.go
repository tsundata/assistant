package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func WorkflowCron(ctx context.Context, comp rulebot.IComponent) []result.Result {
	if comp.Workflow() == nil {
		return []result.Result{result.EmptyResult()}
	}
	_, err := comp.Workflow().CronTrigger(ctx, &pb.TriggerRequest{})
	if err != nil {
		comp.GetLogger().Error(err)
		return []result.Result{result.ErrorResult(err)}
	}
	return []result.Result{result.EmptyResult()}
}
