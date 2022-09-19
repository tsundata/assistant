package service

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/vendors/doctorxiong"
	"net/http"
	"strconv"
	"time"
)

type Finance struct{}

func NewFinance() pb.FinanceSvcServer {
	return &Finance{}
}

func (f *Finance) CreateBill(ctx context.Context, payload *pb.BillRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (f *Finance) GetBill(ctx context.Context, payload *pb.BillRequest) (*pb.BillReply, error) {
	panic("implement me")
}

func (f *Finance) GetBills(ctx context.Context, payload *pb.BillRequest) (*pb.BillsReply, error) {
	panic("implement me")
}

func (f *Finance) DeleteBill(ctx context.Context, payload *pb.BillRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (f *Finance) GetFund(ctx context.Context, payload *pb.TextRequest) (*pb.FundReply, error) {
	code := payload.Text
	now := time.Now()
	startDate := now.AddDate(0, 0, -90).Format("2006-01-02")
	endDate := now.Format("2006-01-02")
	dx := doctorxiong.NewDoctorxiong("")
	resp, err := dx.GetFundDetail(ctx, code, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.New(resp.Message)
	}
	fund := resp.Data
	var netWorthDataDate []string
	var netWorthDataUnit []float64
	var netWorthDataIncrease []float64
	for _, item := range fund.NetWorthData {
		netWorthDataDate = append(netWorthDataDate, item[0].(string))
		f1, _ := strconv.ParseFloat(item[1].(string), 64)
		netWorthDataUnit = append(netWorthDataUnit, f1)
		f2, _ := strconv.ParseFloat(item[2].(string), 64)
		netWorthDataIncrease = append(netWorthDataIncrease, f2)
	}
	var millionCopiesIncomeDataDate []string
	var millionCopiesIncomeDataIncome []float64
	for _, item := range fund.MillionCopiesIncomeData {
		millionCopiesIncomeDataDate = append(millionCopiesIncomeDataDate, item[0].(string))
		f1, _ := strconv.ParseFloat(item[1].(string), 64)
		millionCopiesIncomeDataIncome = append(millionCopiesIncomeDataIncome, f1)
	}

	return &pb.FundReply{
		Code:                          fund.Code,
		Name:                          fund.Name,
		Type:                          fund.Type,
		NetWorth:                      fund.NetWorth,
		ExpectWorth:                   fund.ExpectWorth,
		TotalWorth:                    fund.TotalWorth,
		ExpectGrowth:                  fund.ExpectGrowth,
		DayGrowth:                     fund.DayGrowth,
		LastWeekGrowth:                fund.LastWeekGrowth,
		LastMonthGrowth:               fund.LastMonthGrowth,
		LastThreeMonthsGrowth:         fund.LastThreeMonthsGrowth,
		LastSixMonthsGrowth:           fund.LastSixMonthsGrowth,
		LastYearGrowth:                fund.LastYearGrowth,
		BuyMin:                        fund.BuyMin,
		BuySourceRate:                 fund.BuySourceRate,
		BuyRate:                       fund.BuyRate,
		Manager:                       fund.Manager,
		FundScale:                     fund.FundScale,
		NetWorthDate:                  fund.NetWorthDate,
		ExpectWorthDate:               fund.ExpectWorthDate,
		NetWorthDataDate:              netWorthDataDate,
		NetWorthDataUnit:              netWorthDataUnit,
		NetWorthDataIncrease:          netWorthDataIncrease,
		MillionCopiesIncomeDate:       "",
		SevenDaysYearIncome:           0,
		MillionCopiesIncomeDataDate:   millionCopiesIncomeDataDate,
		MillionCopiesIncomeDataIncome: millionCopiesIncomeDataIncome,
	}, nil
}

func (f *Finance) GetStock(ctx context.Context, payload *pb.TextRequest) (*pb.StockReply, error) {
	code := payload.Text
	dx := doctorxiong.NewDoctorxiong("")
	resp, err := dx.GetStockDetail(ctx, code)
	if err != nil {
		return nil, err
	}
	if resp.Code != http.StatusOK {
		return nil, errors.New(resp.Message)
	}
	stock := resp.Data
	if len(resp.Data) <= 0 {
		return &pb.StockReply{}, nil
	}
	return &pb.StockReply{
		Code:             stock[0].Code,
		Name:             stock[0].Name,
		Type:             stock[0].Type,
		PriceChange:      stock[0].PriceChange,
		ChangePercent:    stock[0].ChangePercent,
		Open:             stock[0].Open,
		Close:            stock[0].Close,
		Price:            stock[0].Price,
		High:             stock[0].High,
		Low:              stock[0].Low,
		Volume:           stock[0].Volume,
		Turnover:         stock[0].Turnover,
		TurnoverRate:     stock[0].TurnoverRate,
		TotalWorth:       stock[0].TotalWorth,
		CirculationWorth: stock[0].CirculationWorth,
		Date:             stock[0].Date,
		Buy:              nil,
		Sell:             nil,
		Pb:               stock[0].Pb,
		Spe:              stock[0].Spe,
		Pe:               stock[0].Pe,
	}, nil
}
