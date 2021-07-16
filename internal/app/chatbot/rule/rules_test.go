package rule

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/version"
	"github.com/tsundata/assistant/mock"
	"strconv"
	"testing"
	"time"
)

func TestVersionRule(t *testing.T) {
	r := rules[0]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "version", []string{})
	require.Equal(t, []string{version.Info()}, res)
}

func TestMenuRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetMenu(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "menu ..."}, nil),
	)

	r := rules[1]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "menu", []string{})
	require.Equal(t, []string{"menu ..."}, res)
}

func TestQrRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetQrUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "https://qr.test/abc"}, nil),
	)

	r := rules[2]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "qr abc", []string{"qr abc", "abc"})
	require.Equal(t, []string{"https://qr.test/abc"}, res)
}

func TestUtRule(t *testing.T) {
	r := rules[3]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "ut 1", []string{"ut 1", "1"})
	require.Equal(t, []string{time.Unix(1, 0).String()}, res)
}

func TestRandRule(t *testing.T) {
	r := rules[4]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "rand 1 100", []string{"rand 1 100", "1", "100"})

	i, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdRule(t *testing.T) {
	r := rules[5]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "pwd 32", []string{"psd 32", "32"})
	require.Equal(t, 32, len(res[0]))
}

func TestSubsListRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	subscribe := mock.NewMockSubscribeClient(ctl)
	gomock.InOrder(
		subscribe.EXPECT().List(gomock.Any(), gomock.Any()).Return(&pb.SubscribeReply{Text: []string{"test1"}}, nil),
	)

	r := rules[6]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, subscribe, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "subs list", []string{"subs list"})
	require.Equal(t, []string{"test1"}, res)
}

func TestSubsOpenRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	subscribe := mock.NewMockSubscribeClient(ctl)
	gomock.InOrder(
		subscribe.EXPECT().Open(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	r := rules[7]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, subscribe, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "subs open test1", []string{"subs open test1", "test1"})
	require.Equal(t, []string{"ok"}, res)
}

func TestSubsCloseRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	subscribe := mock.NewMockSubscribeClient(ctl)
	gomock.InOrder(
		subscribe.EXPECT().Close(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	r := rules[8]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, subscribe, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "subs close test1", []string{"subs close test1", "test1"})
	require.Equal(t, []string{"ok"}, res)
}

func TestViewRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageClient(ctl)
	gomock.InOrder(
		message.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&pb.MessageReply{Id: 1, Text: "test1"}, nil),
	)

	r := rules[9]
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "view 1", []string{"view 1", "1"})
	require.Equal(t, []string{"test1"}, res)
}

func TestRunRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageClient(ctl)
	gomock.InOrder(
		message.EXPECT().Run(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "test1"}, nil),
	)

	r := rules[10]
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "run 1", []string{"run 1", "1"})
	require.Equal(t, []string{"test1"}, res)
}

func TestDocRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	workflow := mock.NewMockWorkflowClient(ctl)
	gomock.InOrder(
		workflow.EXPECT().ActionDoc(gomock.Any(), gomock.Any()).Return(&pb.WorkflowReply{Text: "doc ..."}, nil),
	)

	r := rules[11]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, workflow, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "doc", []string{"doc"})
	require.Equal(t, []string{"doc ..."}, res)
}

func TestStatsRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetStats(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "stats ..."}, nil),
	)

	r := rules[13]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "stats", []string{"stats"})
	require.Equal(t, []string{"stats ..."}, res)
}

func TestTodoRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	r := rules[14]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, todo, nil, nil)
	res := r.Parse(context.Background(), comp, "todo test1", []string{"todo test1", "test1"})
	require.Equal(t, []string{"success"}, res)
}

func TestRoleRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetRoleImageUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "https://web.test/role/test"}, nil),
	)

	r := rules[15]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "role", []string{"role"})
	require.Equal(t, []string{`https://web.test/role/test`}, res)
}

func TestPinyinRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	nlp := mock.NewMockNLPClient(ctl)
	gomock.InOrder(
		nlp.EXPECT().Pinyin(gomock.Any(), gomock.Any()).Return(&pb.WordsReply{Text: []string{"a1", "a2"}}, nil),
	)

	r := rules[16]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nlp)
	res := r.Parse(context.Background(), comp, "pinyin 测试", []string{"pinyin 测试", "测试"})
	require.Equal(t, []string{"a1, a2"}, res)
}

func TestRemindRule(t *testing.T) {
	r := rules[17]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "remind test 19:50", []string{"remind test 19:50", "test", "19:50"})
	require.Equal(t, []string{}, res)
}

func TestDeleteRule(t *testing.T) {
	r := rules[18]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, "del 1", []string{"del 1", "1"})
	require.Equal(t, []string{}, res)
}
