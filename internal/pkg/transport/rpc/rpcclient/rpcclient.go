package rpcclient

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"time"
)

func NewIdClient(client *rpc.Client) (pb.IdSvcClient, error) {
	conn, err := client.Dial(enum.Id, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "id client dial error")
	}
	c := pb.NewIdSvcClient(conn)
	return c, nil
}

func NewMiddleClient(client *rpc.Client) (pb.MiddleSvcClient, error) {
	conn, err := client.Dial(enum.Middle, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewMiddleSvcClient(conn)
	return c, nil
}

func NewMessageClient(client *rpc.Client) (pb.MessageSvcClient, error) {
	conn, err := client.Dial(enum.Message, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "message client dial error")
	}
	c := pb.NewMessageSvcClient(conn)
	return c, nil
}

func NewChatbotClient(client *rpc.Client) (pb.ChatbotSvcClient, error) {
	conn, err := client.Dial(enum.Chatbot, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "chatbot client dial error")
	}
	c := pb.NewChatbotSvcClient(conn)
	return c, nil
}

func NewUserClient(client *rpc.Client) (pb.UserSvcClient, error) {
	conn, err := client.Dial(enum.User, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewUserSvcClient(conn)
	return c, nil
}

func NewNLPClient(client *rpc.Client) (pb.NLPSvcClient, error) {
	conn, err := client.Dial(enum.NLP, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "nlp client dial error")
	}
	c := pb.NewNLPSvcClient(conn)
	return c, nil
}

func NewStorageClient(client *rpc.Client) (pb.StorageSvcClient, error) {
	conn, err := client.Dial(enum.Storage, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "storage client dial error")
	}
	c := pb.NewStorageSvcClient(conn)
	return c, nil
}

func NewWorkflowClient(client *rpc.Client) (pb.WorkflowSvcClient, error) {
	conn, err := client.Dial(enum.Workflow, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "workflow client dial error")
	}
	c := pb.NewWorkflowSvcClient(conn)
	return c, nil
}

func NewFinanceClient(client *rpc.Client) (pb.FinanceSvcClient, error) {
	conn, err := client.Dial(enum.Bot, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewFinanceSvcClient(conn)
	return c, nil
}

func NewOrgClient(client *rpc.Client) (pb.OrgSvcClient, error) {
	conn, err := client.Dial(enum.Bot, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "user client dial error")
	}
	c := pb.NewOrgSvcClient(conn)
	return c, nil
}

func NewTodoClient(client *rpc.Client) (pb.TodoSvcClient, error) {
	conn, err := client.Dial(enum.Bot, rpc.WithTimeout(time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "todo client dial error")
	}
	c := pb.NewTodoSvcClient(conn)
	return c, nil
}

var ProviderSet = wire.NewSet(NewIdClient, NewMiddleClient, NewMessageClient,
	NewChatbotClient, NewUserClient, NewNLPClient, NewStorageClient, NewWorkflowClient,
	NewFinanceClient, NewOrgClient, NewTodoClient)