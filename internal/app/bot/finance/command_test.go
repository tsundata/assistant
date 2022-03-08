package finance

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func parseCommand(t *testing.T, comp component.Component, in string) []string { //nolint
	for _, rule := range Bot.CommandRule {
		tokens, err := command.ParseCommand(in)
		if err != nil {
			t.Fatal(err)
		}
		check, err := command.SyntaxCheck(rule.Define, tokens)
		if err != nil {
			t.Fatal(err)
		}
		if !check {
			continue
		}

		if ret := rule.Parse(context.Background(), comp, tokens); len(ret) > 0 {
			return ret
		}
	}

	return []string{}
}

func TestGetFundCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	finance := mock.NewMockFinanceSvcClient(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		finance.EXPECT().GetFund(gomock.Any(), gomock.Any()).Return(&pb.FundReply{Name: "test"}, nil),
		middle.EXPECT().SetChartData(gomock.Any(), gomock.Any()).Return(&pb.ChartDataReply{ChartData: nil}, nil),
		middle.EXPECT().GetChartUrl(gomock.Any(), gomock.Any()).Return(&pb.TextReply{Text: "http://127.0.0.1:7000/chart/test"}, nil),
	)

	cmd := "fund 000001"
	comp := component.MockComponent(middle)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"http://127.0.0.1:7000/chart/test"}, res)
}

func TestGetStockCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	finance := mock.NewMockFinanceSvcClient(ctl)
	gomock.InOrder(
		finance.EXPECT().GetStock(gomock.Any(), gomock.Any()).Return(&pb.StockReply{Name: "test"}, nil),
	)

	cmd := "stock sx000001"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"Code: \nName: test\nType: \nOpen: \nClose: \n"}, res)
}
