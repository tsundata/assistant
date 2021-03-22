package stage

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Message(b *rulebot.RuleBot, r result.Result) result.Result {
	if r.Kind == result.Message {
		_, err := b.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: r.Content.(string)})
		if err != nil {
			return result.ErrorResult(err)
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
