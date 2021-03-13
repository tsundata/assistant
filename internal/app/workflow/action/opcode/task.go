package opcode

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"strconv"
)

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return false, errors.New("error params")
	}

	if id, ok := params[0].(int64); ok {
		reply, err := ctx.MsgClient.Get(context.Background(), &pb.MessageRequest{Id: id})
		if err != nil {
			return false, err
		}

		j, err := json.Marshal(map[string]string{
			"type": reply.GetType(),
			"id":   strconv.FormatInt(id, 10),
		})
		if err != nil {
			return false, err
		}
		_, err = ctx.TaskClient.Send(context.Background(), &pb.JobRequest{Name: "run", Args: utils.ByteToString(j)})
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
