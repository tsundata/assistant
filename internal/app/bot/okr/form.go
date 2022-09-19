package okr

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
)

var formRules = []bot.FormRule{
	{
		ID:    CreateObjectiveFormID,
		Title: "Create Objective",
		Field: []bot.FieldItem{
			{
				Key:      "title",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "memo",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "motive",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "feasibility",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "is_plan",
				Type:     bot.FieldItemTypeBool,
				Required: true,
			},
			{
				Key:      "plan_start",
				Type:     bot.FieldItemTypeDatetime,
				Required: true,
			},
			{
				Key:      "plan_end",
				Type:     bot.FieldItemTypeDatetime,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			var objective pb.Objective
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "title":
					objective.Title = item.Value.(string)
				case "memo":
					objective.Memo = item.Value.(string)
				case "motive":
					objective.Motive = item.Value.(string)
				case "feasibility":
					objective.Feasibility = item.Value.(string)
				}
			}

			reply, err := comp.Okr().CreateObjective(ctx, &pb.ObjectiveRequest{
				Objective: &objective,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
	{
		ID:    UpdateObjectiveFormID,
		Title: "Update Objective",
		Field: []bot.FieldItem{
			{
				Key:      "sequence",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
			{
				Key:      "title",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "memo",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "motive",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "feasibility",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "is_plan",
				Type:     bot.FieldItemTypeBool,
				Required: true,
			},
			{
				Key:      "plan_start",
				Type:     bot.FieldItemTypeDatetime,
				Required: true,
			},
			{
				Key:      "plan_end",
				Type:     bot.FieldItemTypeDatetime,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			var objective pb.Objective
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "sequence":
					sequence, _ := strconv.ParseInt(item.Value.(string), 10, 64)
					objective.Sequence = sequence
				case "title":
					objective.Title = item.Value.(string)
				case "memo":
					objective.Memo = item.Value.(string)
				case "motive":
					objective.Motive = item.Value.(string)
				case "feasibility":
					objective.Feasibility = item.Value.(string)
				}
			}

			reply, err := comp.Okr().UpdateObjective(ctx, &pb.ObjectiveRequest{
				Objective: &objective,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
	{
		ID:    CreateKeyResultFormID,
		Title: "Create Key Result",
		Field: []bot.FieldItem{
			{
				Key:      "objective_sequence",
				Type:     bot.FieldItemTypeInt,
				Required: true,
				Intro:    "Objective Sequence",
			},
			{
				Key:      "title",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "memo",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "initial_value",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
			{
				Key:      "target_value",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
			{
				Key:      "value_mode",
				Type:     bot.FieldItemTypeString,
				Required: true,
				Intro:    "avg|max|sum|last",
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			objectiveSequence := int64(0)
			var keyResult pb.KeyResult
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "objective_sequence":
					objectiveSequence, _ = strconv.ParseInt(item.Value.(string), 10, 64)
				case "title":
					keyResult.Title = item.Value.(string)
				case "memo":
					keyResult.Memo = item.Value.(string)
				case "initial_value":
					keyResult.InitialValue = util.ParseInt32(item.Value.(string))
				case "target_value":
					keyResult.TargetValue = util.ParseInt32(item.Value.(string))
				case "value_mode":
					keyResult.ValueMode = item.Value.(string)
				}
			}

			reply, err := comp.Okr().CreateKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult:         &keyResult,
				ObjectiveSequence: objectiveSequence,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
	{
		ID:    UpdateKeyResultFormID,
		Title: "Update Key Result",
		Field: []bot.FieldItem{
			{
				Key:      "sequence",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
			{
				Key:      "title",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "memo",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "target_value",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
			{
				Key:      "value_mode",
				Type:     bot.FieldItemTypeString,
				Required: true,
				Intro:    "avg|max|sum|last",
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			var keyResult pb.KeyResult
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "sequence":
					sequence, _ := strconv.ParseInt(item.Value.(string), 10, 64)
					keyResult.Sequence = sequence
				case "title":
					keyResult.Title = item.Value.(string)
				case "memo":
					keyResult.Memo = item.Value.(string)
				case "target_value":
					keyResult.TargetValue = util.ParseInt32(item.Value.(string))
				case "value_mode":
					keyResult.ValueMode = item.Value.(string)
				}
			}

			reply, err := comp.Okr().UpdateKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &keyResult,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
	{
		ID:    CreateKeyResultValueFormID,
		Title: "Create Key Result value",
		Field: []bot.FieldItem{
			{
				Key:      "key_result_sequence",
				Type:     bot.FieldItemTypeInt,
				Required: true,
				Intro:    "Key Result Sequence",
			},
			{
				Key:      "value",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Okr() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			keyResultSequence := int64(0)
			value := int32(0)
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "key_result_sequence":
					keyResultSequence, _ = strconv.ParseInt(item.Value.(string), 10, 64)
				case "value":
					value = util.ParseInt32(item.Value.(string))
				}
			}

			reply, err := comp.Okr().CreateKeyResultValue(ctx, &pb.KeyResultValueRequest{
				KeyResultSequence: keyResultSequence,
				Value:             value,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
}
