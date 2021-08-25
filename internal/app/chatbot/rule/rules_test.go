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

func parseRule(t *testing.T, comp rulebot.IComponent, in string) []string {
	for _, rule := range rules {
		tokens, err := ParseCommand(in)
		if err != nil {
			t.Fatal(err)
		}
		check, err := SyntaxCheck(rule.Define, tokens)
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

func TestVersionRule(t *testing.T) {
	command := "version"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{version.Info()}, res)
}

func TestMenuRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetMenu(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "menu ..."}, nil),
	)

	command := "menu"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"menu ..."}, res)
}

func TestQrRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetQrUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "https://qr.test/abc"}, nil),
	)

	command := "qr abc"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"https://qr.test/abc"}, res)
}

func TestUtRule(t *testing.T) {
	command := "ut 1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{time.Unix(1, 0).String()}, res)
}

func TestRandRule(t *testing.T) {
	command := "rand 1 100"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)

	i, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdRule(t *testing.T) {
	command := "pwd 32"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, 32, len(res[0]))
}

func TestSubsListRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().ListSubscribe(gomock.Any(), gomock.Any()).Return(&pb.SubscribeReply{Subscribe: []*pb.Subscribe{{Name: "test1", State: true}}}, nil),
	)

	command := "subs list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"  NAME  | SUBSCRIBE  \n--------+------------\n  test1 | true       \n"}, res)
}

func TestSubsOpenRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().OpenSubscribe(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "subs open test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestSubsCloseRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().CloseSubscribe(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "subs close test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestViewRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&pb.MessageReply{Message: &pb.Message{Id: 1, Text: "test1"}}, nil),
	)

	command := "view 1"
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"test1"}, res)
}

func TestRunRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Run(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "test1"}, nil),
	)

	command := "run 1"
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"test1"}, res)
}

func TestDocRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	workflow := mock.NewMockWorkflowSvcClient(ctl)
	gomock.InOrder(
		workflow.EXPECT().ActionDoc(gomock.Any(), gomock.Any()).Return(&pb.WorkflowReply{Text: "doc ..."}, nil),
	)

	command := "doc"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, workflow,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Len(t, res, 1)
}

func TestStatsRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetStats(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "stats ..."}, nil),
	)

	command := "stats"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"stats ..."}, res)
}

func TestTodoList(t *testing.T) {
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
				UpdatedAt: "2000-01-01 01:00:00",
			},
		}}, nil),
	)

	command := "todo list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, todo, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"  ID | PRIORITY | CONTENT | COMPLETE |       UPDATE         \n-----+----------+---------+----------+----------------------\n   1 |        1 | todo    | true     | 2000-01-01 01:00:00  \n"}, res)
}

func TestTodoRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "todo test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, todo, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"success"}, res)
}

func TestRoleRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().GetRoleImageUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "https://web.test/role/test"}, nil),
	)

	command := "role"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{`https://web.test/role/test`}, res)
}

func TestPinyinRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	nlp := mock.NewMockNLPSvcClient(ctl)
	gomock.InOrder(
		nlp.EXPECT().Pinyin(gomock.Any(), gomock.Any()).Return(&pb.WordsReply{Text: []string{"a1", "a2"}}, nil),
	)

	command := "pinyin 测试"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nlp, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"a1, a2"}, res)
}

func TestRemindRule(t *testing.T) {
	command := "remind test 19:50"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{}, res)
}

func TestDeleteRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: ""}, nil),
	)

	command := "del 1"
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"Deleted 1"}, res)
}

func TestCronListRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().ListCron(gomock.Any(), gomock.Any()).Return(&pb.CronReply{Cron: []*pb.Cron{{Name: "test", State: true}}}, nil),
	)

	command := "cron list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"  NAME | ISCRON  \n-------+---------\n  test | true    \n"}, res)
}

func TestCronStartRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().StartCron(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "cron start test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestCronStopRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().StopCron(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "cron stop test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestObjList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().GetObjectives(gomock.Any(), gomock.Any()).Return(&pb.ObjectivesReply{Objective: []*pb.Objective{
			{
				Id:    1,
				Name:  "obj",
				Tag:   "obj-1",
				TagId: 1,
			},
		}}, nil),
	)

	command := "obj list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, org, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"  ID | NAME |  TAG   \n-----+------+--------\n   1 | obj  | obj-1  \n"}, res)
}

func TestObjCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().CreateObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "obj obj obj-1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, org, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestObjDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().DeleteObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "obj del 1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, org, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestKrList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().GetKeyResults(gomock.Any(), gomock.Any()).Return(&pb.KeyResultsReply{Result: []*pb.KeyResult{
			{
				Id:          1,
				ObjectiveId: 1,
				Name:        "kr",
				Tag:         "kr-1",
				TagId:       1,
			},
		}}, nil),
	)

	command := "kr list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, org, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"  ID | NAME | OID | TAG  | COMPLETE | UPDATE  \n-----+------+-----+------+----------+---------\n   1 | kr   |   1 | kr-1 | false    |         \n"}, res)
}

func TestKrCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().CreateKeyResult(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "kr 1 kr kr-1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, org, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestKrDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().DeleteKeyResult(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "kr delete 1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, org, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"ok"}, res)
}

func TestGetFund(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	finance := mock.NewMockFinanceSvcClient(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		finance.EXPECT().GetFund(gomock.Any(), gomock.Any()).Return(&pb.FundReply{Name: "test"}, nil),
		middle.EXPECT().SetChartData(gomock.Any(), gomock.Any()).Return(&pb.ChartDataReply{ChartData: nil}, nil),
		middle.EXPECT().GetChartUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "http://127.0.0.1:7000/chart/test"}, nil),
	)

	command := "fund 000001"
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil,
		nil, nil, nil, nil, nil, finance)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"http://127.0.0.1:7000/chart/test"}, res)
}

func TestGetStock(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	finance := mock.NewMockFinanceSvcClient(ctl)
	gomock.InOrder(
		finance.EXPECT().GetStock(gomock.Any(), gomock.Any()).Return(&pb.StockReply{Name: "test"}, nil),
	)

	command := "stock sx000001"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, finance)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"Code: \nName: test\nType: \nOpen: \nClose: \n"}, res)
}

func TestWebhookList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	workflow := mock.NewMockWorkflowSvcClient(ctl)
	gomock.InOrder(
		workflow.EXPECT().ListWebhook(gomock.Any(), gomock.Any()).Return(&pb.WebhooksReply{Flag: []string{"test1", "test2"}}, nil),
	)

	command := "webhook list"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, workflow,
		nil, nil, nil, nil, nil, nil)
	res := parseRule(t, comp, command)
	require.Equal(t, []string{"/webhook/test1\n/webhook/test2\n"}, res)
}
