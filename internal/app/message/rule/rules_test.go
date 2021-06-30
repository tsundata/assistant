package rule

import (
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "version", []string{})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "menu", []string{})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "qr abc", []string{"qr abc", "abc"})
	require.Equal(t, []string{"https://qr.test/abc"}, res)
}

func TestUtRule(t *testing.T) {
	r := rules[3]
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "ut 1", []string{"ut 1", "1"})
	require.Equal(t, []string{time.Unix(1, 0).String()}, res)
}

func TestRandRule(t *testing.T) {
	r := rules[4]
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "rand 1 100", []string{"rand 1 100", "1", "100"})

	i, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdRule(t *testing.T) {
	r := rules[5]
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "pwd 32", []string{"psd 32", "32"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, subscribe, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "subs list", []string{"subs list"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, subscribe, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "subs open test1", []string{"subs open test1", "test1"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, subscribe, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "subs close test1", []string{"subs close test1", "test1"})
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
	ctx := rulebot.NewContext(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "view 1", []string{"view 1", "1"})
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
	ctx := rulebot.NewContext(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "run 1", []string{"run 1", "1"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, workflow, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "doc", []string{"doc"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "stats", []string{"stats"})
	require.Equal(t, []string{"stats ..."}, res)
}

func TestTodoRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoClient(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	r := rules[14]
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, nil, nil, todo, nil, nil)
	res := r.ParseMessage(ctx, "todo test1", []string{"todo test1", "test1"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil)
	res := r.ParseMessage(ctx, "role", []string{"role"})
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
	ctx := rulebot.NewContext(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nlp)
	res := r.ParseMessage(ctx, "pinyin 测试", []string{"pinyin 测试", "测试"})
	require.Equal(t, []string{"a1, a2"}, res)
}
