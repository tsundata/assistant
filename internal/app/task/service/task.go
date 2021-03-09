package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"go.etcd.io/etcd/clientv3"
)

type Task struct {
	etcd *clientv3.Client
}

func NewTask(etcd *clientv3.Client) *Task {
	return &Task{etcd: etcd}
}

func (s *Task) Send(_ context.Context, _ *pb.JobRequest) (*pb.StateReply, error) {
	return &pb.StateReply{
		State: true,
	}, nil
}
