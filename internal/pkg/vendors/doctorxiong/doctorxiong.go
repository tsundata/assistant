package doctorxiong

import (
	"context"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const (
	ID    = "doctorxiong"
	Token = "token"
)

type FundDetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Code                    string          `json:"code"`
		Name                    string          `json:"name"`
		Type                    string          `json:"type"`
		NetWorth                float64         `json:"netWorth"`
		ExpectWorth             float64         `json:"expectWorth"`
		TotalWorth              float64         `json:"totalWorth"`
		ExpectGrowth            string          `json:"expectGrowth"`
		DayGrowth               string          `json:"dayGrowth"`
		LastWeekGrowth          string          `json:"lastWeekGrowth"`
		LastMonthGrowth         string          `json:"lastMonthGrowth"`
		LastThreeMonthsGrowth   string          `json:"lastThreeMonthsGrowth"`
		LastSixMonthsGrowth     string          `json:"lastSixMonthsGrowth"`
		LastYearGrowth          string          `json:"lastYearGrowth"`
		BuyMin                  string          `json:"buyMin"`
		BuySourceRate           string          `json:"buySourceRate"`
		BuyRate                 string          `json:"buyRate"`
		Manager                 string          `json:"manager"`
		FundScale               string          `json:"fundScale"`
		NetWorthDate            string          `json:"netWorthDate"`
		ExpectWorthDate         string          `json:"expectWorthDate"`
		NetWorthData            [][]interface{} `json:"netWorthData"`
		MillionCopiesIncomeData [][]interface{} `json:"millionCopiesIncomeData"`
		MillionCopiesIncomeDate string          `json:"millionCopiesIncomeDate"`
		SevenDaysYearIncome     float64         `json:"sevenDaysYearIncome"`
		SevenDaysYearIncomeDate [][]interface{} `json:"sevenDaysYearIncomeDate"`
	} `json:"data"`
	Meta string `json:"meta"`
}

type StockDetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Code             string   `json:"code"`
		Name             string   `json:"name"`
		Type             string   `json:"type"`
		PriceChange      string   `json:"priceChange"`
		ChangePercent    string   `json:"changePercent"`
		Open             string   `json:"open"`
		Close            string   `json:"close"`
		Price            string   `json:"price"`
		High             string   `json:"high"`
		Low              string   `json:"low"`
		Volume           string   `json:"volume"`
		Turnover         string   `json:"turnover"`
		TurnoverRate     string   `json:"turnoverRate"`
		TotalWorth       string   `json:"totalWorth"`
		CirculationWorth string   `json:"circulationWorth"`
		Date             string   `json:"date"`
		Buy              []string `json:"buy"`
		Sell             []string `json:"sell"`
		Pb               string   `json:"pb"`
		Spe              string   `json:"spe"`
		Pe               string   `json:"pe"`
	} `json:"data"`
	Meta interface{} `json:"meta"`
}

type Doctorxiong struct {
	token string
}

func NewDoctorxiong(token string) *Doctorxiong {
	return &Doctorxiong{token: token}
}

func (v *Doctorxiong) GetFundDetail(ctx context.Context, code, startDate, endDate string) (*FundDetailResponse, error) {
	c := resty.New()
	resp, err := c.R().
		SetContext(ctx).
		//SetHeader("token", v.token).
		SetQueryParam("code", code).
		SetQueryParam("startDate", startDate).
		SetQueryParam("endDate", endDate).
		SetResult(&FundDetailResponse{}).
		Get("https://api.doctorxiong.club/v1/fund/detail")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*FundDetailResponse)
		return result, nil
	}
	return nil, nil
}

func (v *Doctorxiong) GetStockDetail(ctx context.Context, code string) (*StockDetailResponse, error) {
	c := resty.New()
	resp, err := c.R().
		SetContext(ctx).
		//SetHeader("token", v.token).
		SetResult(&StockDetailResponse{}).
		SetQueryParam("code", code).
		Get("https://api.doctorxiong.club/v1/stock")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		result := resp.Result().(*StockDetailResponse)
		return result, nil
	}
	return nil, nil
}
