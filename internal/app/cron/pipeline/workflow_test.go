package pipeline

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func TestWorkflowDone(t *testing.T) {
	ctx := rulebot.NewContext(nil, nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil)

	in := result.DoneResult()

	Workflow(ctx, in)
}

func TestWorkflowError(t *testing.T) {
	l := logger.NewLogger(nil)

	ctx := rulebot.NewContext(nil, nil, l, nil,
		nil, nil, nil, nil,
		nil, nil, nil)

	in := result.ErrorResult(errors.New("test"))

	Workflow(ctx, in)
}

func TestWorkflowMessage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageClient(ctl)
	gomock.InOrder(
		message.EXPECT().
			Send(gomock.Any(), gomock.Any()).
			Return(&pb.StateReply{State: true}, nil),
	)

	ctx := rulebot.NewContext(nil, nil, nil, message,
		nil, nil, nil, nil,
		nil, nil, nil)

	in := result.MessageResult("test")

	Workflow(ctx, in)
}

func TestWorkflowUrl(t *testing.T) {
	ctx := rulebot.NewContext(nil, nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil)

	in := result.Result{
		Kind:    result.Url,
		Content: map[string]interface{}{"test": "test"},
	}

	Workflow(ctx, in)
}

func TestWorkflowRepos(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleClient(ctl)
	gomock.InOrder(
		middle.EXPECT().
			GetAvailableApp(gomock.Any(), gomock.Any()).
			Return(&pb.AppReply{Token: ""}, nil),
	)

	ctx := rulebot.NewContext(nil, nil, nil, nil,
		middle, nil, nil, nil,
		nil, nil, nil)

	in := result.Result{
		Kind: result.Repos,
		Content: map[string]string{
			"test": "test",
		},
	}

	Workflow(ctx, in)
}
func TestWorkflowDefault(t *testing.T) {
	ctx := rulebot.NewContext(nil, nil, nil, nil,
		nil, nil, nil, nil,
		nil, nil, nil)

	in := result.Result{
		Kind: result.Undefined,
	}

	Workflow(ctx, in)
}
