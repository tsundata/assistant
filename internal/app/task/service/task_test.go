package service

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/queue"
	"reflect"
	"testing"
	"time"

	"github.com/tsundata/assistant/api/pb"
)

func TestTask_Delay(t *testing.T) {
	q, err := queue.CreateQueueServer(enum.Task)
	if err != nil {
		t.Fatal(err)
	}
	s := NewTask(q)

	type args struct {
		ctx     context.Context
		payload *pb.JobRequest
	}
	tests := []struct {
		name    string
		s       *Task
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.JobRequest{
				Time: time.Now().Format("2006-01-02 15:04:05"),
				Name: enum.WorkflowRunTask,
				Args: `{"type":"action", "id":"1"}`,
			}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Delay(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.Delay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Delay() = %v, want %v", got, tt.want)
			}
		})
	}
}
