package rule

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/tags"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var rules = []Rule{
	{
		Define: "version",
		Help:   `Version info`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			return []string{version.Info()}
		},
	},
	{
		Define: `qr [string]`,
		Help:   `Generate QR code`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			txt := tokens[1].Value
			reply, err := comp.Middle().GetQrUrl(ctx, &pb.TextRequest{
				Text: txt,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{
				reply.GetText(),
			}
		},
	},
	{
		Define: `ut [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			tt, err := strconv.ParseInt(tokens[1].Value, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			t := time.Unix(tt, 0)

			return []string{
				t.String(),
			}
		},
	},
	{
		Define: `rand [number] [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			minArg := tokens[1].Value
			maxArg := tokens[2].Value
			min, err := strconv.ParseInt(minArg, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			max, err := strconv.ParseInt(maxArg, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			nBing, err := rand.Int(rand.Reader, big.NewInt(max+1-max))
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			t := nBing.Int64() + min

			return []string{
				strconv.FormatInt(t, 10),
			}
		},
	},
	{
		Define: `pwd [number]`,
		Help:   `Generate Password`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			lenArg := tokens[1].Value
			length, err := strconv.Atoi(lenArg)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			pwd := util.GeneratePassword(length, "lowercase|uppercase|numbers")

			return []string{
				pwd,
			}
		},
	},
	{
		Define: "subs list",
		Help:   `List subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().ListSubscribe(ctx, &pb.SubscribeRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Subscribe) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Name", "Subscribe"})
				for _, v := range reply.Subscribe {
					table.Append([]string{v.Name, util.BoolToString(v.State)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"empty subscript"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: "subs open [string]",
		Help:   `Open subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().OpenSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `subs close [string]`,
		Help:   `Close subscribe`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().CloseSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `view [number]`,
		Help:   `View message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(tokens[1].Value, 10, 64)
			if err != nil {
				return []string{"error args"}
			}
			messageReply, err := comp.Message().Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			if messageReply.Message.Id == 0 {
				return []string{"no message"}
			}

			return []string{messageReply.Message.GetText()}
		},
	},
	{
		Define: `run [number]`,
		Help:   `Run message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Message() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			id, err := strconv.ParseInt(tokens[1].Value, 10, 64)
			if err != nil {
				return []string{"error args"}
			}

			reply, err := comp.Message().Run(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{reply.GetText()}
		},
	},
	{
		Define: `doc`,
		Help:   `Show action docs`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Workflow() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Workflow().ActionDoc(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			res := strings.Builder{}
			res.WriteString("Action:\n")
			res.WriteString(reply.GetText())
			res.WriteString("\n\nTag:\n")
			for k := range tags.Tags() {
				res.WriteString("#")
				res.WriteString(k)
				res.WriteString("\n")
			}
			return []string{res.String()}
		},
	},
	{
		Define: `test`,
		Help:   `Test`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			_, err := comp.Message().Send(ctx, &pb.MessageRequest{Message: &pb.Message{Text: "test"}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Middle().GetSubscribeStatus(ctx, &pb.SubscribeRequest{Text: "example"})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: "测试"})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Todo().GetTodo(ctx, &pb.TodoRequest{Todo: &pb.Todo{Id: 1}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.User().GetUser(ctx, &pb.UserRequest{User: &pb.User{Id: 1}})
			if err != nil {
				return []string{err.Error()}
			}
			_, err = comp.Workflow().SyntaxCheck(ctx, &pb.WorkflowRequest{Text: "echo 1"})
			if err != nil {
				return []string{err.Error()}
			}
			return []string{"test done"}
		},
	},
	{
		Define: `stats`,
		Help:   `Stats Info`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().GetStats(ctx, &pb.TextRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			return []string{reply.GetText()}
		},
	},
	{
		Define: `todo list`,
		Help:   `List todo`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Todo() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Todo().GetTodos(ctx, &pb.TodoRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Todos) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Id", "Priority", "Content", "Complete"})
				for _, v := range reply.Todos {
					table.Append([]string{strconv.Itoa(int(v.Id)), strconv.Itoa(int(v.Priority)), v.Content, util.BoolToString(v.Complete)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"Empty"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: `todo [string]`,
		Help:   "Todo something",
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Todo() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}
			reply, err := comp.Todo().CreateTodo(ctx, &pb.TodoRequest{
				Todo: &pb.Todo{Content: tokens[1].Value},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if !reply.GetState() {
				return []string{"failed"}
			}
			return []string{"success"}
		},
	},
	{
		Define: `pinyin [string]`,
		Help:   "chinese pinyin conversion",
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.NLP() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}
			reply, err := comp.NLP().Pinyin(ctx, &pb.TextRequest{Text: tokens[1].Value})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if len(reply.GetText()) <= 0 {
				return []string{"failed"}
			}
			return []string{strings.Join(reply.GetText(), ", ")}
		},
	},
	{
		Define: `remind [string] [string]`,
		Help:   `Remind something`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			arg1 := tokens[1].Value
			arg2 := tokens[2].Value
			fmt.Println(arg1, arg2) // todo remind message

			return []string{}
		},
	},
	{
		Define: `del [number]`,
		Help:   `Delete message`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			idStr := tokens[1].Value
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			_, err = comp.Message().Delete(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			return []string{fmt.Sprintf("Deleted %d", id)}
		},
	},
	{
		Define: "cron list",
		Help:   `List cron`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Middle().ListCron(ctx, &pb.CronRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Cron) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Name", "IsCron"})
				for _, v := range reply.Cron {
					table.Append([]string{v.Name, util.BoolToString(v.State)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"empty cron"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: "cron start [string]",
		Help:   `Start cron`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().StartCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `cron stop [string]`,
		Help:   `Stop cron`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Middle().StopCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `obj list`,
		Help:   `List objectives`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Org().GetObjectives(ctx, &pb.ObjectiveRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Objective) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Id", "Name"})
				for _, v := range reply.Objective {
					table.Append([]string{strconv.Itoa(int(v.Id)), v.Name})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"Empty"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: `obj del [number]`,
		Help:   `Delete objective`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			idStr := tokens[2].Value
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			reply, err := comp.Org().DeleteObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{Id: id},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `obj [string] [string]`,
		Help:   `Create Objective`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			reply, err := comp.Org().CreateObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{
					//Tag:  tokens[1].Value, // todo tag
					Name: tokens[2].Value,
				},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `kr list`,
		Help:   `List KeyResult`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Org().GetKeyResults(ctx, &pb.KeyResultRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			tableString := &strings.Builder{}
			if len(reply.Result) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Id", "Name", "OID", "Complete"})
				for _, v := range reply.Result {
					table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, strconv.Itoa(int(v.ObjectiveId)), util.BoolToString(v.Complete)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []string{"Empty"}
			}

			return []string{tableString.String()}
		},
	},
	{
		Define: `kr [number] [string] [string]`,
		Help:   `Create KeyResult`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 4 {
				return []string{"error args"}
			}

			idStr := tokens[1].Value
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			reply, err := comp.Org().CreateKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{
					ObjectiveId: id,
					//Tag:         tokens[2].Value,// todo tag
					Name: tokens[3].Value,
				},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `kr delete [number]`,
		Help:   `Delete KeyResult`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 3 {
				return []string{"error args"}
			}

			idStr := tokens[2].Value
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return []string{"error call: " + err.Error()}
			}

			reply, err := comp.Org().DeleteKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{Id: id},
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetState() {
				return []string{"ok"}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `fund [string]`,
		Help:   `Get fund`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Finance() == nil {
				return []string{"empty client"}
			}
			if comp.Middle() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			reply, err := comp.Finance().GetFund(ctx, &pb.TextRequest{
				Text: tokens[1].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetName() != "" {
				var xAxis []string
				var series []float64
				if reply.NetWorthDataDate == nil || len(reply.NetWorthDataDate) == 0 {
					xAxis = reply.MillionCopiesIncomeDataDate
					series = reply.MillionCopiesIncomeDataIncome
				} else {
					xAxis = reply.NetWorthDataDate
					series = reply.NetWorthDataUnit
				}

				chartReply, err := comp.Middle().SetChartData(ctx, &pb.ChartDataRequest{
					ChartData: &pb.ChartData{
						Title:    fmt.Sprintf("Fund %s (%s)", reply.Name, reply.Code),
						SubTitle: "Data for the last 90 days",
						XAxis:    xAxis,
						Series:   series,
					},
				})
				if err != nil {
					return []string{"chart failed"}
				}
				urlReply, err := comp.Middle().GetChartUrl(ctx, &pb.TextRequest{Text: chartReply.ChartData.GetUuid()})
				if err != nil {
					return []string{"url failed"}
				}

				return []string{urlReply.Text}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `stock [string]`,
		Help:   `Get stock`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Finance() == nil {
				return []string{"empty client"}
			}
			if len(tokens) != 2 {
				return []string{"error args"}
			}

			reply, err := comp.Finance().GetStock(ctx, &pb.TextRequest{
				Text: tokens[1].Value,
			})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			if reply.GetName() != "" {
				var res strings.Builder
				res.WriteString("Code: ")
				res.WriteString(reply.Code)
				res.WriteString("\n")
				res.WriteString("Name: ")
				res.WriteString(reply.Name)
				res.WriteString("\n")
				res.WriteString("Type: ")
				res.WriteString(reply.Type)
				res.WriteString("\n")
				res.WriteString("Open: ")
				res.WriteString(reply.Open)
				res.WriteString("\n")
				res.WriteString("Close: ")
				res.WriteString(reply.Close)
				res.WriteString("\n")
				return []string{res.String()}
			}

			return []string{"failed"}
		},
	},
	{
		Define: `webhook list`,
		Help:   `List webhook`,
		Parse: func(ctx context.Context, comp rulebot.IComponent, tokens []*Token) []string {
			if comp.Workflow() == nil {
				return []string{"empty client"}
			}

			reply, err := comp.Workflow().ListWebhook(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []string{"error call: " + err.Error()}
			}
			var result strings.Builder
			for _, flag := range reply.Flag {
				result.WriteString("/webhook/")
				result.WriteString(flag)
				result.WriteString("\n")
			}
			if result.Len() > 0 {
				return []string{result.String()}
			}

			return []string{"failed"}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
