package finance

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strings"
)

var commandRules = []command.Rule{
	{
		Define: `fund [string]`,
		Help:   `Get fund`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Finance() == nil || comp.Middle() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			text, _ := tokens[1].Value.String()
			reply, err := comp.Finance().GetFund(ctx, &pb.TextRequest{
				Text: text,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
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
					return []pb.MsgPayload{pb.TextMsg{Text: "chart failed"}}
				}
				urlReply, err := comp.Middle().GetChartUrl(ctx, &pb.TextRequest{Text: chartReply.ChartData.GetUuid()})
				if err != nil {
					return []pb.MsgPayload{pb.TextMsg{Text: "url failed"}}
				}

				return []pb.MsgPayload{pb.TextMsg{Text: urlReply.Text}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
	{
		Define: `stock [string]`,
		Help:   `Get stock`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Finance() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			text, _ := tokens[1].Value.String()
			reply, err := comp.Finance().GetStock(ctx, &pb.TextRequest{
				Text: text,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
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
				return []pb.MsgPayload{pb.TextMsg{Text: res.String()}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
		},
	},
}
