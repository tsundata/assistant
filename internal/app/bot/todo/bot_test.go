package todo

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/version"
	"github.com/tsundata/assistant/mock"
	"strconv"
	"testing"
	"time"
)

func parseCommand(t *testing.T, comp command.Component, in string) []string {
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"https://qr.test/abc"}, res)
}

func TestUtCommand(t *testing.T) {
	cmd := "ut 1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{time.Unix(1, 0).String()}, res)
}

func TestRandCommand(t *testing.T) {
	cmd := "rand 1 100"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)

	i, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdCommand(t *testing.T) {
	cmd := "pwd 32"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestViewCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{Sequence: 1, Text: "test1"}}, nil),
	)

	cmd := "view 1"
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"test1"}, res)
}

func TestDocCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	workflow := mock.NewMockWorkflowSvcClient(ctl)
	gomock.InOrder(
		workflow.EXPECT().ActionDoc(gomock.Any(), gomock.Any()).Return(&pb.WorkflowReply{Text: "doc ..."}, nil),
	)

	cmd := "doc"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, workflow,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"stats ..."}, res)
}

func TestTodoList(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().GetTodos(gomock.Any(), gomock.Any()).Return(&pb.TodosReply{Todos: []*pb.Todo{
			{
				Id:        1,
				Priority:  1,
				Content:   "todo",
				Complete:  true,
				UpdatedAt: 946659600,
			},
		}}, nil),
	)

	cmd := "todo list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"  ID | PRIORITY | CONTENT | COMPLETE  \n-----+----------+---------+-----------\n   1 |        1 | todo    | true      \n"}, res)
}

func TestTodoCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "todo test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"success"}, res)
}

func TestPinyinCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	nlp := mock.NewMockNLPSvcClient(ctl)
	gomock.InOrder(
		nlp.EXPECT().Pinyin(gomock.Any(), gomock.Any()).Return(&pb.WordsReply{Text: []string{"a1", "a2"}}, nil),
	)

	cmd := "pinyin 测试"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nlp)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"a1, a2"}, res)
}

func TestRemindCommand(t *testing.T) {
	cmd := "remind test 19:50"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{}, res)
}

func TestDeleteCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: ""}, nil),
	)

	cmd := "del 1"
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"ok"}, res)
}

func TestGetFundCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	finance := mock.NewMockFinanceSvcClient(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		finance.EXPECT().GetFund(gomock.Any(), gomock.Any()).Return(&pb.FundReply{Name: "test"}, nil),
		middle.EXPECT().SetChartData(gomock.Any(), gomock.Any()).Return(&pb.ChartDataReply{ChartData: nil}, nil),
		middle.EXPECT().GetChartUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "http://127.0.0.1:7000/chart/test"}, nil),
	)

	cmd := "fund 000001"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"http://127.0.0.1:7000/chart/test"}, res)
}

func TestGetStockCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	finance := mock.NewMockFinanceSvcClient(ctl)
	gomock.InOrder(
		finance.EXPECT().GetStock(gomock.Any(), gomock.Any()).Return(&pb.StockReply{Name: "test"}, nil),
	)

	cmd := "stock sx000001"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"Code: \nName: test\nType: \nOpen: \nClose: \n"}, res)
}

func TestWebhookListCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	workflow := mock.NewMockWorkflowSvcClient(ctl)
	gomock.InOrder(
		workflow.EXPECT().ListWebhook(gomock.Any(), gomock.Any()).Return(&pb.WebhooksReply{Flag: []string{"test1", "test2"}}, nil),
	)

	cmd := "webhook list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, workflow,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"/webhook/test1\n/webhook/test2\n"}, res)
}
