package stage

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
)

func Message(b *rulebot.Context, r result.Result) result.Result {
	if r.Kind == result.Message {
		_, err := rpcclient.GetMessageClient(b.Client).Send(context.Background(), &pb.MessageRequest{Text: r.Content.(string)})
		if err != nil {
			return result.ErrorResult(err)
		}
		return result.DoneResult()
	}
	return result.EmptyResult()
}
