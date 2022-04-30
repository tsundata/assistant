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

func TestSearchMetadata(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().
			CollectMetadata(gomock.Any(), gomock.Any()).
			Return(&pb.StateReply{State: true}, nil),
	)

	comp := component.MockComponent(middle)

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
			if got := SearchMetadata(context.Background(), tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
