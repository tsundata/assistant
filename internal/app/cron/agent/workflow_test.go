package agent

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"
)

func TestWorkflowCron(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	workflow := mock.NewMockWorkflowSvcClient(ctl)
	gomock.InOrder(
		workflow.EXPECT().
			CronTrigger(gomock.Any(), gomock.Any()).
			Return(&pb.WorkflowReply{Text: ""}, nil),
	)

	comp := rulebot.NewComponent(nil, nil, nil,
		nil, nil, workflow, nil,
		nil, nil, nil, nil, nil)

	type args struct {
		comp rulebot.IComponent
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
			if got := WorkflowCron(context.Background(), tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WorkflowCron() = %v, want %v", got, tt.want)
			}
		})
	}
}
