package service

import (
	"context"
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/tsundata/assistant/api/pb"
	"time"
)

type TaskSvcClient interface {
	Delay(context.Context, *pb.JobRequest) (*pb.StateReply, error)
}

type Task struct {
	server *machinery.Server
}

func NewTask(server *machinery.Server) TaskSvcClient {
	return &Task{server: server}
}

func (s *Task) Delay(ctx context.Context, payload *pb.JobRequest) (*pb.StateReply, error) {
	eta, err := time.ParseInLocation("2006-01-02 15:04:05", payload.GetTime(), time.Local)
	if err != nil {
		return nil, err
	}
	task := tasks.Signature{
		ETA:  &eta,
		Name: payload.GetName(),
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: payload.GetArgs(),
			},
		},
	}

	_, err = s.server.SendTaskWithContext(ctx, &task)
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{
		State: true,
	}, nil
}
