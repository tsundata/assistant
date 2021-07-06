package rpcclient

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewChatbotClient(client *rpc.Client) (pb.ChatbotClient, error) {
	conn, err := client.Dial(app.Chatbot, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "chatbot client dial error")
	}
	c := pb.NewChatbotClient(conn)
	return c, nil
}
