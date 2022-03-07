package org

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

var commandRules = []command.Rule{
	{
		Define: `obj list`,
		Help:   `List objectives`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Org() == nil {
			//	return []string{"empty client"}
			//}
			//
			//reply, err := comp.Org().GetObjectives(ctx, &pb.ObjectiveRequest{})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//
			//tableString := &strings.Builder{}
			//if len(reply.Objective) > 0 {
			//	table := tablewriter.NewWriter(tableString)
			//	table.SetBorder(false)
			//	table.SetHeader([]string{"Id", "Name"})
			//	for _, v := range reply.Objective {
			//		table.Append([]string{strconv.Itoa(int(v.Id)), v.Name})
			//	}
			//	table.Render()
			//}
			//if tableString.String() == "" {
			//	return []string{"Empty"}
			//}
			//
			//return []string{tableString.String()}
			return []string{}
		},
	},
	{
		Define: `obj del [number]`,
		Help:   `Delete objective`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Org() == nil {
			//	return []string{"empty client"}
			//}
			//idStr := tokens[2].Value
			//id, err := strconv.ParseInt(idStr, 10, 64)
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//
			//reply, err := comp.Org().DeleteObjective(ctx, &pb.ObjectiveRequest{
			//	Objective: &pb.Objective{Id: id},
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if reply.GetState() {
			//	return []string{"ok"}
			//}
			//
			//return []string{"failed"}
			return []string{}
		},
	},
	{
		Define: `obj [string] [string]`,
		Help:   `Create Objective`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Org() == nil {
			//	return []string{"empty client"}
			//}
			//reply, err := comp.Org().CreateObjective(ctx, &pb.ObjectiveRequest{
			//	Objective: &pb.Objective{
			//		//Tag:  tokens[1].Value, // todo tag
			//		Name: tokens[2].Value,
			//	},
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if reply.GetState() {
			//	return []string{"ok"}
			//}
			//
			//return []string{"failed"}
			return []string{}
		},
	},
	{
		Define: `kr list`,
		Help:   `List KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Org() == nil {
			//	return []string{"empty client"}
			//}
			//
			//reply, err := comp.Org().GetKeyResults(ctx, &pb.KeyResultRequest{})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//
			//tableString := &strings.Builder{}
			//if len(reply.Result) > 0 {
			//	table := tablewriter.NewWriter(tableString)
			//	table.SetBorder(false)
			//	table.SetHeader([]string{"Id", "Name", "OID", "Complete"})
			//	for _, v := range reply.Result {
			//		table.Append([]string{strconv.Itoa(int(v.Id)), v.Name, strconv.Itoa(int(v.ObjectiveId)), util.BoolToString(v.Complete)})
			//	}
			//	table.Render()
			//}
			//if tableString.String() == "" {
			//	return []string{"Empty"}
			//}
			//
			//return []string{tableString.String()}
			return []string{}
		},
	},
	{
		Define: `kr [number] [string] [string]`,
		Help:   `Create KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Org() == nil {
			//	return []string{"empty client"}
			//}
			//idStr := tokens[1].Value
			//id, err := strconv.ParseInt(idStr, 10, 64)
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//
			//reply, err := comp.Org().CreateKeyResult(ctx, &pb.KeyResultRequest{
			//	KeyResult: &pb.KeyResult{
			//		ObjectiveId: id,
			//		//Tag:         tokens[2].Value,// todo tag
			//		Name: tokens[3].Value,
			//	},
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if reply.GetState() {
			//	return []string{"ok"}
			//}
			//
			//return []string{"failed"}
			return []string{}
		},
	},
	{
		Define: `kr delete [number]`,
		Help:   `Delete KeyResult`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Org() == nil {
			//	return []string{"empty client"}
			//}
			//idStr := tokens[2].Value
			//id, err := strconv.ParseInt(idStr, 10, 64)
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//
			//reply, err := comp.Org().DeleteKeyResult(ctx, &pb.KeyResultRequest{
			//	KeyResult: &pb.KeyResult{Id: id},
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if reply.GetState() {
			//	return []string{"ok"}
			//}
			//
			//return []string{"failed"}
			return []string{}
		},
	},
	{
		Define: `fund [string]`,
		Help:   `Get fund`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Finance() == nil {
			//	return []string{"empty client"}
			//}
			//if comp.Middle() == nil {
			//	return []string{"empty client"}
			//}
			//
			//reply, err := comp.Finance().GetFund(ctx, &pb.TextRequest{
			//	Text: tokens[1].Value,
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if reply.GetName() != "" {
			//	var xAxis []string
			//	var series []float64
			//	if reply.NetWorthDataDate == nil || len(reply.NetWorthDataDate) == 0 {
			//		xAxis = reply.MillionCopiesIncomeDataDate
			//		series = reply.MillionCopiesIncomeDataIncome
			//	} else {
			//		xAxis = reply.NetWorthDataDate
			//		series = reply.NetWorthDataUnit
			//	}
			//
			//	chartReply, err := comp.Middle().SetChartData(ctx, &pb.ChartDataRequest{
			//		ChartData: &pb.ChartData{
			//			Title:    fmt.Sprintf("Fund %s (%s)", reply.Name, reply.Code),
			//			SubTitle: "Data for the last 90 days",
			//			XAxis:    xAxis,
			//			Series:   series,
			//		},
			//	})
			//	if err != nil {
			//		return []string{"chart failed"}
			//	}
			//	urlReply, err := comp.Middle().GetChartUrl(ctx, &pb.TextRequest{Text: chartReply.ChartData.GetUuid()})
			//	if err != nil {
			//		return []string{"url failed"}
			//	}
			//
			//	return []string{urlReply.Text}
			//}
			//
			//return []string{"failed"}
			return []string{}
		},
	},
	{
		Define: `stock [string]`,
		Help:   `Get stock`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []string {
			//if comp.Finance() == nil {
			//	return []string{"empty client"}
			//}
			//reply, err := comp.Finance().GetStock(ctx, &pb.TextRequest{
			//	Text: tokens[1].Value,
			//})
			//if err != nil {
			//	return []string{"error call: " + err.Error()}
			//}
			//if reply.GetName() != "" {
			//	var res strings.Builder
			//	res.WriteString("Code: ")
			//	res.WriteString(reply.Code)
			//	res.WriteString("\n")
			//	res.WriteString("Name: ")
			//	res.WriteString(reply.Name)
			//	res.WriteString("\n")
			//	res.WriteString("Type: ")
			//	res.WriteString(reply.Type)
			//	res.WriteString("\n")
			//	res.WriteString("Open: ")
			//	res.WriteString(reply.Open)
			//	res.WriteString("\n")
			//	res.WriteString("Close: ")
			//	res.WriteString(reply.Close)
			//	res.WriteString("\n")
			//	return []string{res.String()}
			//}
			//
			//return []string{"failed"}
			return []string{}
		},
	},
}
