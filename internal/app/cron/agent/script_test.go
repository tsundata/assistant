package agent

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"
)

func TestWorkflowCron(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatbot := mock.NewMockChatbotSvcClient(ctl)
	gomock.InOrder(
		chatbot.EXPECT().
			CronTrigger(gomock.Any(), gomock.Any()).
			Return(&pb.WorkflowReply{Text: ""}, nil),
	)

	comp := component.MockComponent(chatbot)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want []result.Result
	}{
		{
			"case1",
			args{comp},
			[]result.Result{result.EmptyResult()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScriptCron(context.Background(), tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptCron() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWatchCron(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatbot := mock.NewMockChatbotSvcClient(ctl)
	gomock.InOrder(
		chatbot.EXPECT().
			WatchTrigger(gomock.Any(), gomock.Any()).
			Return(&pb.WorkflowReply{Text: ""}, nil),
	)

	comp := component.MockComponent(chatbot)

	type args struct {
		comp component.Component
	}
	tests := []struct {
		name string
		args args
		want []result.Result
	}{
		{
			"case1",
			args{comp},
			[]result.Result{result.EmptyResult()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScriptWatch(context.Background(), tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
