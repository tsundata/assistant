package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"math/rand"
	"reflect"
	"testing"
)

func TestWorkflow_SyntaxCheck(t *testing.T) {
	s := NewWorkflow(nil, nil, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"ok syntax",
			s,
			args{context.Background(), &pb.WorkflowRequest{
				Type: enum.MessageTypeAction,
				Text: `get 1
echo 1
message "11"
`}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"error syntax",
			s,
			args{context.Background(), &pb.WorkflowRequest{
				Type: enum.MessageTypeAction,
				Text: `get a a;`}},
			&pb.StateReply{State: false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SyntaxCheck(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.SyntaxCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.SyntaxCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_RunAction(t *testing.T) {
	s := NewWorkflow(nil, nil, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"run action",
			s,
			args{context.Background(), &pb.WorkflowRequest{
				Text: `echo "ok"`}},
			&pb.WorkflowReply{Text: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RunAction(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.RunAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.RunAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_WebhookTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	nats, err := event.CreateNats(enum.Workflow)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewNatsBus(nats, nil)

	repo := mock.NewMockWorkflowRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetTriggerByFlag(gomock.Any(), gomock.Any(), gomock.Any()).Return(pb.Trigger{MessageId: 1, Secret: "test"}, nil),
	)

	s := NewWorkflow(bus, nil, repo, nil, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{Trigger: &pb.Trigger{Type: "webhook", Secret: "test"}}},
			&pb.WorkflowReply{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.WebhookTrigger(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.WebhookTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.WebhookTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_CronTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	rdb, err := vendors.CreateRedisClient(enum.User)
	if err != nil {
		t.Fatal(err)
	}

	nats, err := event.CreateNats(enum.Workflow)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewNatsBus(nats, nil)

	messageID := rand.Int63()
	repo := mock.NewMockWorkflowRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().
			ListTriggersByType(gomock.Any(), "cron").
			Return([]pb.Trigger{{MessageId: messageID, When: "* * * * *"}}, nil),
		repo.EXPECT().
			ListTriggersByType(gomock.Any(), "cron").
			Return([]pb.Trigger{{MessageId: messageID, When: "* * * * *"}}, nil),
	)

	s := NewWorkflow(bus, rdb, repo, nil, nil, nil)

	type args struct {
		ctx context.Context
		in1 *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{}},
			&pb.WorkflowReply{},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.TriggerRequest{}},
			&pb.WorkflowReply{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CronTrigger(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.CronTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.CronTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_CreateTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockWorkflowRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateTrigger(gomock.Any(), gomock.Any()).Return(int64(1), nil),
		repo.EXPECT().GetTriggerByFlag(gomock.Any(), gomock.Any(), gomock.Any()).Return(pb.Trigger{}, nil),
		repo.EXPECT().CreateTrigger(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	s := NewWorkflow(nil, nil, repo, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{
				Trigger: &pb.Trigger{
					Kind:      enum.MessageTypeAction,
					Type:      "webhook",
					MessageId: 1,
					//					MessageText: `
					//cron "* * * * *"
					//webhook "test"
					//`,
				},
			}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateTrigger(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.CreateTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.CreateTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_DeleteTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockWorkflowRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().DeleteTriggerByMessageID(gomock.Any(), gomock.Any()).Return(nil),
		repo.EXPECT().DeleteTriggerByMessageID(gomock.Any(), gomock.Any()).Return(errors.New("not record")),
	)

	s := NewWorkflow(nil, nil, repo, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{Trigger: &pb.Trigger{MessageId: 1}}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{Trigger: &pb.Trigger{MessageId: 2}}},
			&pb.StateReply{State: false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteTrigger(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.DeleteTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.DeleteTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkflow_ActionDoc(t *testing.T) {
	s := NewWorkflow(nil, nil, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.WorkflowRequest{}},
			&pb.WorkflowReply{Text: ""},
			false,
		},
		{
			"case1",
			s,
			args{context.Background(), &pb.WorkflowRequest{Text: "webhook"}},
			&pb.WorkflowReply{Text: "webhook [string] [string]?"},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ActionDoc(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.ActionDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if i == 0 {
				if got != nil && len(got.Text) == 0 {
					t.Errorf("Workflow.ActionDoc() = %v, want %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Workflow.ActionDoc() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestWorkflow_ListWebhook(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockWorkflowRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListTriggersByType(gomock.Any(), gomock.Any()).Return([]pb.Trigger{{Flag: "test1"}, {Flag: "test2"}}, nil),
	)

	s := NewWorkflow(nil, nil, repo, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Workflow
		args    args
		want    *pb.WebhooksReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.WorkflowRequest{}},
			&pb.WebhooksReply{Flag: []string{"test1", "test2"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListWebhook(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.ListWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.ListWebhook() = %v, want %v", got, tt.want)
			}
		})
	}
}
