package system

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/msg"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/trigger/tags"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/version"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var commandRules = []command.Rule{
	{
		Define: "version",
		Help:   `Version info`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{pb.TextMsg{Text: version.Info()}}
		},
	},
	{
		Define: `qr [string]`,
		Help:   `Generate QR code`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Middle().GetQrUrl(ctx, &pb.TextRequest{
				Text: tokens[1].Value.(string),
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: reply.GetText()}}
		},
	},
	{
		Define: `ut [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			t := time.Unix(tokens[1].Value.(int64), 0)

			return []pb.MsgPayload{pb.TextMsg{Text: t.String()}}
		},
	},
	{
		Define: `rand [number] [number]`,
		Help:   `Unix Timestamp`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			min := tokens[1].Value.(int64)
			max := tokens[2].Value.(int64)

			nBing, err := rand.Int(rand.Reader, big.NewInt(max+1-min))
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			t := nBing.Int64() + min

			return []pb.MsgPayload{pb.TextMsg{Text: strconv.FormatInt(t, 10)}}
		},
	},
	{
		Define: `pwd [number]`,
		Help:   `Generate Password`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			length := tokens[1].Value.(int64)

			pwd := util.RandString(int(length), "lowercase|uppercase|numbers")

			return []pb.MsgPayload{pb.TextMsg{Text: pwd}}
		},
	},
	{
		Define: "subs list",
		Help:   `List subscribe`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			reply, err := comp.Middle().ListSubscribe(ctx, &pb.SubscribeRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			tableString := &strings.Builder{}
			if len(reply.Subscribe) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Name", "Subscribe"})
				for _, v := range reply.Subscribe {
					table.Append([]string{v.Name, strconv.Itoa(int(v.Status))})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty subscript"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: tableString.String()}}
		},
	},
	{
		Define: "subs open [string]",
		Help:   `Open subscribe`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Middle().OpenSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
	{
		Define: `subs close [string]`,
		Help:   `Close subscribe`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Middle().CloseSubscribe(ctx, &pb.SubscribeRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
	{
		Define: `view [number]`,
		Help:   `View message`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Message() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			id := tokens[1].Value.(int64)
			messageReply, err := comp.Message().GetBySequence(ctx, &pb.MessageRequest{Message: &pb.Message{UserId: 0, Sequence: id}}) // fixme
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			if messageReply.Message.Sequence == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "no message"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: messageReply.Message.GetText()}}
		},
	},
	{
		Define: `run [number]`,
		Help:   `Run message`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Message() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			id := tokens[1].Value.(int64)

			reply, err := comp.Message().Run(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: reply.GetText()}}
		},
	},
	{
		Define: `doc`,
		Help:   `Show action docs`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Chatbot() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			reply, err := comp.Chatbot().ActionDoc(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
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
			return []pb.MsgPayload{pb.TextMsg{Text: res.String()}}
		},
	},
	{
		Define: `test`,
		Help:   `Test`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			_, err := comp.Message().Send(ctx, &pb.MessageRequest{Message: &pb.Message{Text: "test"}})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: err.Error()}}
			}
			_, err = comp.Middle().GetSubscribeStatus(ctx, &pb.SubscribeRequest{Text: "example"})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: err.Error()}}
			}
			_, err = comp.Middle().Pinyin(ctx, &pb.TextRequest{Text: "测试"})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: err.Error()}}
			}
			_, err = comp.User().GetUser(ctx, &pb.UserRequest{User: &pb.User{Id: 1}})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: err.Error()}}
			}
			_, err = comp.Chatbot().SyntaxCheck(ctx, &pb.WorkflowRequest{Text: "echo 1"})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: err.Error()}}
			}
			return []pb.MsgPayload{pb.TextMsg{Text: "test done"}}
		},
	},
	{
		Define: `stats`,
		Help:   `Stats Info`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			reply, err := comp.Middle().GetStats(ctx, &pb.TextRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			return []pb.MsgPayload{pb.TextMsg{Text: reply.GetText()}}
		},
	},
	{
		Define: `pinyin [string]`,
		Help:   "chinese pinyin conversion",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			reply, err := comp.Middle().Pinyin(ctx, &pb.TextRequest{Text: tokens[1].Value.(string)})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if len(reply.GetText()) <= 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
			}
			return []pb.MsgPayload{pb.TextMsg{Text: strings.Join(reply.GetText(), ", ")}}
		},
	},
	{
		Define: `del [number]`,
		Help:   `Delete message`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			id := tokens[1].Value.(int64)

			_, err := comp.Message().Delete(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: fmt.Sprintf("Deleted %d", id)}}
		},
	},
	{
		Define: "cron list",
		Help:   `List cron`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			reply, err := comp.Middle().ListCron(ctx, &pb.CronRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
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
				return []pb.MsgPayload{pb.TextMsg{Text: "empty cron"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: tableString.String()}}
		},
	},
	{
		Define: "cron start [string]",
		Help:   `Start cron`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Middle().StartCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
	{
		Define: `cron stop [string]`,
		Help:   `Stop cron`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Middle().StopCron(ctx, &pb.CronRequest{
				Text: tokens[2].Value.(string),
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
	{
		Define: `webhook list`,
		Help:   `List webhook`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Chatbot() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Chatbot().ListWebhook(ctx, &pb.WorkflowRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			var result strings.Builder
			for _, flag := range reply.Flag {
				result.WriteString("/webhook/")
				result.WriteString(flag)
				result.WriteString("\n")
			}
			if result.Len() > 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: result.String()}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
	{
		Define: `push switch`,
		Help:   "Push notification switch",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(formRules, PushSwitchFormID)}
		},
	},
}
