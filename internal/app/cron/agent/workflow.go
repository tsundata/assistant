package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
)

func WorkflowCron(b *rulebot.Context) []result.Result {
	if b.Client == nil {
		return []result.Result{result.EmptyResult()}
	}
	_, err := rpcclient.GetWorkflowClient(b.Client).CronTrigger(context.Background(), &pb.TriggerRequest{})
	if err != nil {
		b.Logger.Error(err)
		return []result.Result{result.ErrorResult(err)}
	}
	return []result.Result{result.EmptyResult()}
}
