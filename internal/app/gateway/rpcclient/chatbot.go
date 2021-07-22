package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewChatbotClient(client *rpc.Client) (pb.ChatbotSvcClient, error) {
	conn, err := client.Dial(enum.Chatbot, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "chatbot client dial error")
	}
	c := pb.NewChatbotSvcClient(conn)
	return c, nil
}
