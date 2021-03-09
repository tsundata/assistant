package service

import (
	"context"
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/tsundata/assistant/api/pb"
)

type Task struct {
	server *machinery.Server
}

func NewTask(server *machinery.Server) *Task {
	return &Task{server: server}
}

func (s *Task) Send(ctx context.Context, _ *pb.JobRequest) (*pb.StateReply, error) {
	runTask := tasks.Signature{
		Name: "run",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	_, err := s.server.SendTaskWithContext(ctx, &runTask)
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{
		State: true,
	}, nil
}
