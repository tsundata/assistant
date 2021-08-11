package pipeline

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func TestWorkflowDone(t *testing.T) {
	comp := rulebot.NewComponent(nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil, nil)

	in := result.DoneResult()

	Workflow(context.Background(), comp, in)
}

func TestWorkflowError(t *testing.T) {
	z := log.NewZapLogger(nil)
	l := log.NewAppLogger(z)

	comp := rulebot.NewComponent(nil, nil, l,
		nil, nil, nil, nil,
		nil, nil, nil, nil)

	in := result.ErrorResult(errors.New("test"))

	Workflow(context.Background(), comp, in)
}

func TestWorkflowMessage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().
			Send(gomock.Any(), gomock.Any()).
			Return(&pb.StateReply{State: true}, nil),
	)

	comp := rulebot.NewComponent(nil, nil, nil, message,
		nil, nil, nil,
		nil, nil, nil, nil)

	in := result.MessageResult("test")

	Workflow(context.Background(), comp, in)
}

func TestWorkflowUrl(t *testing.T) {
	comp := rulebot.NewComponent(nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil, nil)

	in := result.Result{
		Kind:    result.Url,
		Content: map[string]interface{}{"test": "test"},
	}

	Workflow(context.Background(), comp, in)
}

func TestWorkflowRepos(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().
			GetAvailableApp(gomock.Any(), gomock.Any()).
			Return(&pb.AppReply{Token: ""}, nil),
	)

	comp := rulebot.NewComponent(nil, nil, nil, nil,
		middle, nil, nil,
		nil, nil, nil, nil)

	in := result.Result{
		Kind: result.Repos,
		Content: map[string]string{
			"test": "test",
		},
	}

	Workflow(context.Background(), comp, in)
}
func TestWorkflowDefault(t *testing.T) {
	comp := rulebot.NewComponent(nil, nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil)

	in := result.Result{
		Kind: result.Undefined,
	}

	Workflow(context.Background(), comp, in)
}
