package robot

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

func Register(ctx context.Context, chatbot pb.ChatbotSvcClient, metadata ...Metadata) error {
	for _, item := range metadata {
		_, err := chatbot.Register(ctx, &pb.BotRequest{Bot: &pb.Bot{
			Name:       item.Name,
			Identifier: item.Identifier,
			Detail:     item.Detail,
			Avatar:     item.Avatar,
			Extend:     "",
		}})
		if err != nil {
			return err
		}
	}
	return nil
}
