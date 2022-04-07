package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

func ScriptCron(ctx context.Context, comp component.Component) []result.Result {
	if comp.Chatbot() == nil {
		return []result.Result{result.EmptyResult()}
	}
	_, err := comp.Chatbot().CronTrigger(ctx, &pb.TriggerRequest{})
	if err != nil {
		comp.GetLogger().Error(err)
		return []result.Result{result.ErrorResult(err)}
	}
	return []result.Result{result.EmptyResult()}
}
