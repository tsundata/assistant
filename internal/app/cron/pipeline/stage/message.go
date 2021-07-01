package stage

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Message(ctx rulebot.IContext, in result.Result) result.Result {
	if in.Kind == result.Message {
		if ctx.Message() == nil {
			return result.EmptyResult()
		}
		_, err := ctx.Message().Send(context.Background(), &pb.MessageRequest{Text: in.Content.(string)})
		if err != nil {
			return result.ErrorResult(err)
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
