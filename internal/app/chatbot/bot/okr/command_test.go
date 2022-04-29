package okr

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func parseCommand(t *testing.T, comp component.Component, in string) []pb.MsgPayload { //nolint
	for _, rule := range Bot.CommandRule {
		tokens, err := command.ParseCommand(in)
		if err != nil {
			t.Fatal(err)
		}
		check, err := command.SyntaxCheck(rule.Define, tokens)
		if err != nil {
			t.Fatal(err)
		}
		if !check {
			continue
		}

		if ret := rule.Parse(context.Background(), comp, tokens); len(ret) > 0 {
			return ret
		}
	}

	return []pb.MsgPayload{}
}

func TestObjListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	okr := mock.NewMockOkrSvcServer(ctl)
	gomock.InOrder(
		okr.EXPECT().GetObjectives(gomock.Any(), gomock.Any()).Return(&pb.ObjectivesReply{Objective: []*pb.Objective{
			{
				Id:           1,
				UserId:       1,
				Sequence:     1,
				Title:        "obj",
				CurrentValue: 1,
				TotalValue:   10,
			},
		}}, nil),
	)

	cmd := "obj list"
	comp := component.MockComponent(okr)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TableMsg{
		Title:  "Objectives",
		Header: []string{"Sequence", "Title", "Current Value", "Total Value"},
		Row: [][]interface{}{
			{"1", "obj", "1", "10"},
		},
	}}, res)
}

func TestObjCreateCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	okr := mock.NewMockOkrSvcServer(ctl)

	cmd := "obj create"
	comp := component.MockComponent(okr)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, 1, len(res))
}

func TestObjDeleteCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	okr := mock.NewMockOkrSvcServer(ctl)
	gomock.InOrder(
		okr.EXPECT().DeleteObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "obj del 1"
	comp := component.MockComponent(okr)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "ok"}}, res)
}

func TestKrListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	okr := mock.NewMockOkrSvcServer(ctl)
	gomock.InOrder(
		okr.EXPECT().GetKeyResults(gomock.Any(), gomock.Any()).Return(&pb.KeyResultsReply{Result: []*pb.KeyResult{
			{
				Id:           1,
				ObjectiveId:  1,
				Title:        "kr",
				Sequence:     1,
				CurrentValue: 1,
				TargetValue:  10,
			},
		}}, nil),
	)

	cmd := "kr list"
	comp := component.MockComponent(okr)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TableMsg{
		Title:  "KeyResult",
		Header: []string{"Sequence", "Title", "Current Value", "Target Value"},
		Row: [][]interface{}{
			{"1", "kr", "1", "10"},
		},
	}}, res)
}

func TestKrCreateCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	okr := mock.NewMockOkrSvcServer(ctl)

	cmd := "kr create"
	comp := component.MockComponent(okr)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, 1, len(res))
}

func TestKrDeleteCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	okr := mock.NewMockOkrSvcServer(ctl)
	gomock.InOrder(
		okr.EXPECT().DeleteKeyResult(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "kr del 1"
	comp := component.MockComponent(okr)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "ok"}}, res)
}
