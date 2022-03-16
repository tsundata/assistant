package work

import (
	"encoding/json"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
)

type WorkflowTask struct {
	bus     event.Bus
	message pb.MessageSvcClient
	chatbot pb.ChatbotSvcClient
}

func NewWorkflowTask(bus event.Bus, message pb.MessageSvcClient, chatbot pb.ChatbotSvcClient) *WorkflowTask {
	return &WorkflowTask{bus: bus, message: message, chatbot: chatbot}
}

func (t *WorkflowTask) Run(data string) (bool, error) {
	var args map[string]string
	err := json.Unmarshal(util.StringToByte(data), &args)
	if err != nil {
		return false, err
	}

	tp, ok := args["type"]
	if !ok {
		return false, app.ErrInvalidParameter
	}

	idStr, ok := args["id"]
	if !ok {
		return false, app.ErrInvalidParameter
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return false, err
	}

	ctx := md.BuildAuthContext(enum.SuperUserID)
	message, err := t.message.GetById(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
	if err != nil {
		return false, err
	}

	switch enum.MessageType(tp) {
	case enum.MessageTypeAction:
		_, err = t.chatbot.RunAction(ctx, &pb.WorkflowRequest{Message: message.Message})
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, app.ErrInvalidParameter
	}
}
