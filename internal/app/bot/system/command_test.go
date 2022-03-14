package system

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/version"
	"github.com/tsundata/assistant/mock"
	"strconv"
	"testing"
	"time"
)

func parseCommand(t *testing.T, comp component.Component, in string) []string {
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

func TestVersionCommand(t *testing.T) {
	cmd := "version"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{version.Info()}, res)
}

func TestQrCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetQrUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "https://qr.test/abc"}, nil),
	)

	cmd := "qr abc"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"https://qr.test/abc"}, res)
}

func TestUtCommand(t *testing.T) {
	cmd := "ut 1"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{time.Unix(1, 0).String()}, res)
}

func TestRandCommand(t *testing.T) {
	cmd := "rand 1 100"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)

	i, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdCommand(t *testing.T) {
	cmd := "pwd 32"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, 32, len(res[0]))
}

func TestSubsListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().ListSubscribe(gomock.Any(), gomock.Any()).Return(&pb.SubscribeReply{Subscribe: []*pb.Subscribe{{Name: "test1", State: true}}}, nil),
	)

	cmd := "subs list"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"  NAME  | SUBSCRIBE  \n--------+------------\n  test1 | true       \n"}, res)
}

func TestSubsOpenCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().OpenSubscribe(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "subs open test1"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestSubsCloseCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().CloseSubscribe(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "subs close test1"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestViewCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{UserId: enum.SuperUserID, Sequence: 1, Text: "test1"}}, nil),
	)

	cmd := "view 1"
	comp := component.MockComponent(message)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"test1"}, res)
}

func TestRunCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Run(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "test1"}, nil),
	)

	cmd := "run 1"
	comp := component.MockComponent(message)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"test1"}, res)
}

func TestDocCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatbot := mock.NewMockChatbotSvcClient(ctl)
	gomock.InOrder(
		chatbot.EXPECT().ActionDoc(gomock.Any(), gomock.Any()).Return(&pb.WorkflowReply{Text: "doc ..."}, nil),
	)

	cmd := "doc"
	comp := component.MockComponent(chatbot)
	res := parseCommand(t, comp, cmd)
	require.Len(t, res, 1)
}

func TestStatsCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetStats(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "stats ..."}, nil),
	)

	cmd := "stats"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"stats ..."}, res)
}

func TestPinyinCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().Pinyin(gomock.Any(), gomock.Any()).Return(&pb.WordsReply{Text: []string{"a1", "a2"}}, nil),
	)

	cmd := "pinyin 测试"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"a1, a2"}, res)
}

func TestDeleteCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: ""}, nil),
	)

	cmd := "del 1"
	comp := component.MockComponent(message)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"Deleted 1"}, res)
}

func TestCronListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().ListCron(gomock.Any(), gomock.Any()).Return(&pb.CronReply{Cron: []*pb.Cron{{Name: "test", State: true}}}, nil),
	)

	cmd := "cron list"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"  NAME | ISCRON  \n-------+---------\n  test | true    \n"}, res)
}

func TestCronStartCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().StartCron(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "cron start test1"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestCronStopCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().StopCron(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "cron stop test1"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestWebhookListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatbot := mock.NewMockChatbotSvcClient(ctl)
	gomock.InOrder(
		chatbot.EXPECT().ListWebhook(gomock.Any(), gomock.Any()).Return(&pb.WebhooksReply{Flag: []string{"test1", "test2"}}, nil),
	)

	cmd := "webhook list"
	comp := component.MockComponent(chatbot)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"/webhook/test1\n/webhook/test2\n"}, res)
}
