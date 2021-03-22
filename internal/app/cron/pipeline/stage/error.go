package stage

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Error(b *rulebot.RuleBot, r result.Result) result.Result {
	if r.Kind == result.Message {
		_, err := b.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: fmt.Sprintf("Error: %s", r.Content)})
		if err != nil {
			return result.EmptyResult()
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
