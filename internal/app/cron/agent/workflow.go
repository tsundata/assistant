package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"log"
)

func WorkflowCron(b *rulebot.RuleBot) []result.Result {
	_, err := b.WfClient.CronTrigger(context.Background(), &pb.TriggerRequest{})
	if err != nil {
		log.Println(err)
		return []result.Result{result.ErrorResult(err)}
	}
	return []result.Result{result.EmptyResult()}
}
