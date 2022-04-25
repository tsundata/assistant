package okr

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/msg"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strconv"
)

var commandRules = []command.Rule{
	{
		Define: `obj list`,
		Help:   `List objectives`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Okr().GetObjectives(ctx, &pb.ObjectiveRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			var header []string
			var row [][]interface{}
			if len(reply.Objective) > 0 {
				header = []string{"Sequence", "Title", "Current Value", "Total Value"}
				for _, v := range reply.Objective {
					row = append(row, []interface{}{strconv.Itoa(int(v.Sequence)), v.Title, strconv.Itoa(int(v.CurrentValue)), strconv.Itoa(int(v.TotalValue))})
				}
			}
			if len(row) == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "Empty"}}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "Objectives",
				Header: header,
				Row:    row,
			}}
		},
	},
	{
		Define: `obj del [number]`,
		Help:   `Delete objective`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			id := tokens[2].Value.(int64)

			reply, err := comp.Okr().DeleteObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{Id: id},
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
		Define: `obj create`,
		Help:   `Create Objective`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(Bot.FormRule, CreateObjectiveFormID)}
		},
	},
	{
		Define: `kr list`,
		Help:   `List KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Okr().GetKeyResults(ctx, &pb.KeyResultRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			var header []string
			var row [][]interface{}
			if len(reply.Result) > 0 {
				header = []string{"Sequence", "Title", "Current Value", "Target Value"}
				for _, v := range reply.Result {
					row = append(row, []interface{}{strconv.Itoa(int(v.Sequence)), v.Title, strconv.Itoa(int(v.CurrentValue)), strconv.Itoa(int(v.TargetValue))})
				}
			}
			if len(row) == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "Empty"}}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "KeyResult",
				Header: header,
				Row:    row,
			}}
		},
	},
	{
		Define: `kr create`,
		Help:   `Create KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(Bot.FormRule, CreateKeyResultFormID)}
		},
	},
	{
		Define: `kr delete [number]`,
		Help:   `Delete KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			id := tokens[2].Value.(int64)

			reply, err := comp.Okr().DeleteKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{Id: id},
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
		Define: `kr value`,
		Help:   `Create KeyResult value`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(Bot.FormRule, CreateKeyResultValueFormID)}
		},
	},
}
