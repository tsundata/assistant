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
	command := "version"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[0]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[1]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[2]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"https://qr.test/abc"}, res)
}

func TestUtRule(t *testing.T) {
	command := "ut 1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[3]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{time.Unix(1, 0).String()}, res)
}

func TestRandRule(t *testing.T) {
	command := "rand 1 100"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[4]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)

	i, err := strconv.ParseInt(res[0], 10, 64)
	if err != nil {
		t.Fatal(err)
	}
	require.True(t, i > 0)
}

func TestPwdRule(t *testing.T) {
	command := "pwd 32"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[5]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[6]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"+-------+-----------+\n| NAME  | SUBSCRIBE |\n+-------+-----------+\n| test1 | true      |\n+-------+-----------+\n"}, res)
}

func TestSubsOpenRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().OpenSubscribe(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "subs open test1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[7]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[8]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[9]
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[10]
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[11]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, workflow, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[13]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"stats ..."}, res)
}

func TestTodoRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "todo test1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[14]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, todo, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[15]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[16]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nlp, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"a1, a2"}, res)
}

func TestRemindRule(t *testing.T) {
	command := "remind test 19:50"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[17]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[18]
	comp := rulebot.NewComponent(nil, nil, nil, message, nil, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"Deleted 1"}, res)
}

func TestCronListRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().ListCron(gomock.Any(), gomock.Any()).Return(&pb.CronReply{Text: []string{"test1"}}, nil),
	)

	command := "cron list"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[19]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"test1"}, res)
}

func TestCronStartRule(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		middle.EXPECT().StartCron(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "cron start test1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[20]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[21]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[22]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, org, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"+----+------+-------+\n| ID | NAME |  TAG  |\n+----+------+-------+\n|  1 | obj  | obj-1 |\n+----+------+-------+\n"}, res)
}

func TestObjCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().CreateObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "obj obj obj-1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[23]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, org, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"ok"}, res)
}

func TestObjDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().DeleteObjective(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "obj delete 1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[24]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, org, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[25]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, org, nil)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"+----+------+-----+------+----------+--------+\n| ID | NAME | OID | TAG  | COMPLETE | UPDATE |\n+----+------+-----+------+----------+--------+\n|  1 | kr   |   1 | kr-1 | false    |        |\n+----+------+-----+------+----------+--------+\n"}, res)
}

func TestKrCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	org := mock.NewMockOrgSvcClient(ctl)
	gomock.InOrder(
		org.EXPECT().CreateKeyResult(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	command := "kr 1 kr kr-1"
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[26]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, org, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[27]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, org, nil)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[28]
	comp := rulebot.NewComponent(nil, nil, nil, nil, middle, nil, nil, nil, nil, nil, nil, finance)
	res := r.Parse(context.Background(), comp, tokens)
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
	tokens, err := ParseCommand(command)
	if err != nil {
		t.Fatal(err)
	}
	r := rules[29]
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, finance)
	res := r.Parse(context.Background(), comp, tokens)
	require.Equal(t, []string{"Code: \nName: test\nType: \nOpen: \nClose: \n"}, res)
}
