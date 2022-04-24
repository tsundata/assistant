package system

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/msg"
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
			reply, err := comp.Middle().GetUserSubscribe(ctx, &pb.TextRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			var header []string
			var row [][]interface{}
			if len(reply.Subscribe) > 0 {
				header = []string{"Name", "Subscribe"}
				for _, v := range reply.Subscribe {
					row = append(row, []interface{}{v.Key, v.Value})
				}
			}
			if len(row) == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty subscript"}}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "Subscribes",
				Header: header,
				Row:    row,
			}}
		},
	},
	{
		Define: "subs switch",
		Help:   `Subscribe switch`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			var field []pb.FormField
			reply, err := comp.Middle().GetUserSubscribe(ctx, &pb.TextRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			for _, item := range reply.Subscribe {
				field = append(field, pb.FormField{
					Key:      item.Key,
					Type:     string(bot.FieldItemTypeString),
					Required: true,
				})
			}
			return []pb.MsgPayload{pb.FormMsg{
				ID:    SubscribeSwitchFormID,
				Title: "Subscribe switch",
				Field: field,
			}}
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
			messageReply, err := comp.Message().GetBySequence(ctx, &pb.MessageRequest{Message: &pb.Message{Sequence: id}})
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
			res.WriteString("\n")
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

			var header []string
			var row [][]interface{}
			if len(reply.Cron) > 0 {
				header = []string{"Name", "IsCron"}
				for _, v := range reply.Cron {
					row = append(row, []interface{}{v.Name, util.BoolToString(v.State)})
				}
			}
			if len(row) == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty cron"}}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "Cron",
				Header: header,
				Row:    row,
			}}
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

			var header []string
			var row [][]interface{}
			if len(reply.Flag) > 0 {
				header = []string{"No", "Webhook url"}
				for i, flag := range reply.Flag {
					row = append(row, []interface{}{i + 1, fmt.Sprintf("%s/webhook/%s", comp.GetConfig().Gateway.Url, flag)})
				}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "Webhook",
				Header: header,
				Row:    row,
			}}
		},
	},
	{
		Define: `push switch`,
		Help:   "Push notification switch",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(formRules, PushSwitchFormID)}
		},
	},
	{
		Define: "webhook switch",
		Help:   `Script Webhook switch`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Chatbot() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			var field []pb.FormField
			reply, err := comp.Chatbot().GetWebhookTriggers(ctx, &pb.TriggerRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			for _, item := range reply.List {
				field = append(field, pb.FormField{
					Key:      strconv.FormatInt(item.MessageId, 10),
					Type:     string(bot.FieldItemTypeString),
					Required: true,
					Intro:    fmt.Sprintf("/webhook/%s", item.Flag),
				})
			}
			return []pb.MsgPayload{pb.FormMsg{
				ID:    WebhookSwitchFormID,
				Title: "Script Webhook switch",
				Field: field,
			}}
		},
	},
	{
		Define: "cron switch",
		Help:   `Script cron switch`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Chatbot() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			var field []pb.FormField
			reply, err := comp.Chatbot().GetCronTriggers(ctx, &pb.TriggerRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			for _, item := range reply.List {
				field = append(field, pb.FormField{
					Key:      strconv.FormatInt(item.MessageId, 10),
					Type:     string(bot.FieldItemTypeString),
					Required: true,
					Intro:    fmt.Sprintf("#%d (%s)", item.Sequence, item.When),
				})
			}
			return []pb.MsgPayload{pb.FormMsg{
				ID:    CronSwitchFormID,
				Title: "Script cron switch",
				Field: field,
			}}
		},
	},
	{
		Define: `counters`,
		Help:   `List Counter`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			reply, err := comp.System().GetCounters(ctx, &pb.CounterRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			var header []string
			var row [][]interface{}
			if len(reply.Counters) > 0 {
				header = []string{"No", "Title", "Digit"}
				for i, item := range reply.Counters {
					row = append(row, []interface{}{i + 1, item.Flag, item.Digit})
				}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "Counter",
				Header: header,
				Row:    row,
			}}
		},
	},
	{
		Define: "counter [string]",
		Help:   `Count things`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			counter := &pb.Counter{Flag: tokens[1].Value.(string), Digit: int64(1)}
			find, err := comp.System().GetCounterByFlag(ctx, &pb.CounterRequest{Counter: counter})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if find.Counter.Id > 0 {
				return []pb.MsgPayload{pb.DigitMsg{
					Title: find.Counter.Flag,
					Digit: int(find.Counter.Digit),
				}}
			}
			_, err = comp.System().CreateCounter(ctx, &pb.CounterRequest{Counter: counter})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			return []pb.MsgPayload{pb.DigitMsg{
				Title: counter.Flag,
				Digit: int(counter.Digit),
			}}
		},
	},
	{
		Define: "increase [string]",
		Help:   `Increase Counter`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			counter := &pb.Counter{Flag: tokens[1].Value.(string), Digit: int64(1)}
			latest, err := comp.System().ChangeCounter(ctx, &pb.CounterRequest{Counter: counter})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			return []pb.MsgPayload{pb.DigitMsg{
				Title: latest.Counter.Flag,
				Digit: int(latest.Counter.Digit),
			}}
		},
	},
	{
		Define: "decrease [string]",
		Help:   `Decrease Counter`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			counter := &pb.Counter{Flag: tokens[1].Value.(string), Digit: int64(-1)}
			latest, err := comp.System().ChangeCounter(ctx, &pb.CounterRequest{Counter: counter})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			return []pb.MsgPayload{pb.DigitMsg{
				Title: latest.Counter.Flag,
				Digit: int(latest.Counter.Digit),
			}}
		},
	},
	{
		Define: "reset [string]",
		Help:   `Reset Counter`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			counter := &pb.Counter{Flag: tokens[1].Value.(string)}
			latest, err := comp.System().ResetCounter(ctx, &pb.CounterRequest{Counter: counter})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			return []pb.MsgPayload{pb.DigitMsg{
				Title: latest.Counter.Flag,
				Digit: int(latest.Counter.Digit),
			}}
		},
	},
}
