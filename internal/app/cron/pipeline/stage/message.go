package stage

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func Message(ctx context.Context, comp rulebot.IComponent, in result.Result) result.Result {
	if in.Kind == result.Message {
		if comp.Message() == nil {
			return result.EmptyResult()
		}
		_, err := comp.Message().Send(ctx, &pb.MessageRequest{Message: &pb.Message{Text: in.Content.(string)}})
		if err != nil {
			return result.ErrorResult(err)
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
