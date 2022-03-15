package org

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
	"strings"
)

var commandRules = []command.Rule{
	{
		Define: `obj list`,
		Help:   `List objectives`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
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
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			id := tokens[2].Value.(int64)

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
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			reply, err := comp.Org().CreateObjective(ctx, &pb.ObjectiveRequest{
				Objective: &pb.Objective{
					//Tag:  tokens[1].Value, // todo tag
					Name: tokens[2].Value.(string),
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
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
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
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			id := tokens[1].Value.(int64)

			reply, err := comp.Org().CreateKeyResult(ctx, &pb.KeyResultRequest{
				KeyResult: &pb.KeyResult{
					ObjectiveId: id,
					//Tag:         tokens[2].Value,// todo tag
					Name: tokens[3].Value.(string),
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
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			if comp.Org() == nil {
				return []string{"empty client"}
			}
			id := tokens[2].Value.(int64)

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
}
