package system

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/version"
	"github.com/tsundata/assistant/mock"
	"strconv"
	"testing"
	"time"
)

func parseCommand(t *testing.T, comp component.Component, in string) []pb.MsgPayload {
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

func TestVersionCommand(t *testing.T) {
	cmd := "version"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: version.Info()}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "https://qr.test/abc"}}, res)
}

func TestUtCommand(t *testing.T) {
	cmd := "ut 1"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: time.Unix(1, 0).String()}}, res)
}

func TestRandCommand(t *testing.T) {
	cmd := "rand 1 100"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)

	i, err := strconv.ParseInt(res[0].(pb.TextMsg).Text, 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdCommand(t *testing.T) {
	cmd := "pwd 32"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, 32, len(res[0].(pb.TextMsg).Text))
}

func TestSubsListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetUserSubscribe(gomock.Any(), gomock.Any()).Return(&pb.GetUserSubscribeReply{Subscribe: []*pb.KV{{Key: "test1", Value: "1"}}}, nil),
	)

	cmd := "subs list"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TableMsg{
		Title:  "Subscribes",
		Header: []string{"Name", "Subscribe"},
		Row: [][]interface{}{
			{"test1", "1"},
		},
	}}, res)
}

func TestSubsOpenCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetUserSubscribe(gomock.Any(), gomock.Any()).Return(&pb.GetUserSubscribeReply{Subscribe: []*pb.KV{{Key: "test1", Value: "1"}}}, nil),
	)

	cmd := "subs switch"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.True(t, len(res) > 0)
}

func TestViewCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().GetBySequence(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{UserId: 1, Sequence: 1, Text: "test1"}}, nil),
	)

	cmd := "view 1"
	comp := component.MockComponent(message)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "test1"}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "test1"}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "stats ..."}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "a1, a2"}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "Deleted 1"}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TableMsg{
		Title:  "Cron",
		Header: []string{"Name", "IsCron"},
		Row: [][]interface{}{
			{"test", "true"},
		},
	}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "ok"}}, res)
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
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "ok"}}, res)
}

func TestWebhookListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatbot := mock.NewMockChatbotSvcClient(ctl)
	gomock.InOrder(
		chatbot.EXPECT().ListWebhook(gomock.Any(), gomock.Any()).Return(&pb.WebhooksReply{Flag: []string{"test1", "test2"}}, nil),
	)

	conf, err := config.CreateAppConfig(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}
	cmd := "webhook list"
	comp := component.MockComponent(chatbot, conf)
	res := parseCommand(t, comp, cmd)
	require.True(t, len(res) > 0)
}
