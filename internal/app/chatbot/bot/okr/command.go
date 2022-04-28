package okr

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/msg"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
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
		Define: `obj [number]`,
		Help:   `View objective`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			sequence, _ := tokens[1].Value.Int64()

			reply, err := comp.Okr().GetObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{Sequence: sequence},
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			keyResultReply, err := comp.Okr().GetKeyResults(ctx, &pb.KeyResultRequest{ObjectiveSequence: sequence})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.OkrMsg{
				Title:     fmt.Sprintf("Objective #%d", reply.Objective.GetSequence()),
				Objective: *reply.Objective,
				KeyResult: keyResultReply.Result,
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
			sequence, _ := tokens[2].Value.Int64()

			reply, err := comp.Okr().DeleteObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{Sequence: sequence},
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
		Define: `obj update [number]`,
		Help:   `Update objective`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			sequence, _ := tokens[2].Value.Int64()

			reply, err := comp.Okr().GetObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{Sequence: sequence},
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.FormMsg{
				ID:    UpdateObjectiveFormID,
				Title: fmt.Sprintf("Update Objective #%d", sequence),
				Field: []pb.FormField{
					{
						Key:      "sequence",
						Type:     string(bot.FieldItemTypeInt),
						Required: true,
						Value:    reply.Objective.Sequence,
					},
					{
						Key:      "title",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Objective.Title,
					},
					{
						Key:      "memo",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Objective.Memo,
					},
					{
						Key:      "motive",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Objective.Motive,
					},
					{
						Key:      "feasibility",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Objective.Feasibility,
					},
					{
						Key:      "is_plan",
						Type:     string(bot.FieldItemTypeBool),
						Required: true,
						Value:    reply.Objective.IsPlan,
					},
					{
						Key:      "plan_start",
						Type:     string(bot.FieldItemTypeDatetime),
						Required: true,
						Value:    reply.Objective.PlanStart,
					},
					{
						Key:      "plan_end",
						Type:     string(bot.FieldItemTypeDatetime),
						Required: true,
						Value:    reply.Objective.PlanEnd,
					},
				},
			}}
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
		Define: `kr del [number]`,
		Help:   `Delete KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			sequence, _ := tokens[2].Value.Int64()

			reply, err := comp.Okr().DeleteKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{Sequence: sequence},
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
		Define: `kr update [number]`,
		Help:   `Update KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			sequence, _ := tokens[2].Value.Int64()

			reply, err := comp.Okr().GetKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{Sequence: sequence},
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.FormMsg{
				ID:    UpdateKeyResultFormID,
				Title: fmt.Sprintf("Update KeyResult #%d", sequence),
				Field: []pb.FormField{
					{
						Key:      "sequence",
						Type:     string(bot.FieldItemTypeInt),
						Required: true,
						Value:    reply.KeyResult.Sequence,
					},
					{
						Key:      "title",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.KeyResult.Title,
					},
					{
						Key:      "memo",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.KeyResult.Memo,
					},
					{
						Key:      "target_value",
						Type:     string(bot.FieldItemTypeInt),
						Required: true,
						Value:    reply.KeyResult.TargetValue,
					},
					{
						Key:      "value_mode",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Intro:    "avg|max|sum|last",
						Value:    reply.KeyResult.ValueMode,
					},
				},
			}}
		},
	},
	{
		Define: `kr value`,
		Help:   `Create KeyResult value`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(Bot.FormRule, CreateKeyResultValueFormID)}
		},
	},
	{
		Define: `kr value [number]`,
		Help:   `List KeyResult value`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			sequence, _ := tokens[2].Value.Int64()

			reply, err := comp.Okr().GetKeyResultValues(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{Sequence: sequence},
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			var header []string
			var row [][]interface{}
			if len(reply.Values) > 0 {
				header = []string{"Value", "Datetime"}
				for _, v := range reply.Values {
					row = append(row, []interface{}{strconv.Itoa(int(v.Value)), util.Format(v.CreatedAt)})
				}
			}
			if len(row) == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "Empty"}}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  fmt.Sprintf("KeyResult #%d Values", sequence),
				Header: header,
				Row:    row,
			}}
		},
	},
}
