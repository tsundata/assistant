package agent

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"log"
)

func WorkflowCron(b *rulebot.RuleBot) []string {
	_, err := b.WfClient.CronTrigger(context.Background(), &pb.TriggerRequest{})
	if err != nil {
		log.Println(err)
		return []string{}
	}
	return []string{}
}
