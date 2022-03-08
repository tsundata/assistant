package org

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

func parseCommand(t *testing.T, comp component.Component, in string) []string { //nolint
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

	return []string{}
}

func TestObjListCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().GetObjectives(gomock.Any(), gomock.Any()).Return(&pb.ObjectivesReply{Objective: []*pb.Objective{
			{
				Id:    1,
				Name:  "obj",
				TagId: 1,
			},
		}}, nil),
	)

	cmd := "obj list"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"  ID | NAME  \n-----+-------\n   1 | obj   \n"}, res)
}

func TestObjCreateCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().CreateObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "obj obj obj-1"
	comp := component.MockComponent(org)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestObjDeleteCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().DeleteObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "obj del 1"
	comp := component.MockComponent(org)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestKrListCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().GetKeyResults(gomock.Any(), gomock.Any()).Return(&pb.KeyResultsReply{Result: []*pb.KeyResult{
			{
				Id:          1,
				ObjectiveId: 1,
				Name:        "kr",
				TagId:       1,
			},
		}}, nil),
	)

	cmd := "kr list"
	comp := component.MockComponent(org)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"  ID | NAME | OID | COMPLETE  \n-----+------+-----+-----------\n   1 | kr   |   1 | false     \n"}, res)
}

func TestKrCreateCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().CreateKeyResult(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "kr 1 kr kr-1"
	comp := component.MockComponent(org)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestKrDeleteCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().DeleteKeyResult(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "kr delete 1"
	comp := component.MockComponent(org)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}
