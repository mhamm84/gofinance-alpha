package alpha

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/mhamm84/gofinance-alpha/alpha/data"
	"strings"
)

const (
	cpi                  string = "CPI"
	consumerSentiment    string = "CONSUMER_SENTIMENT"
	treasuryYield        string = "TREASURY_YIELD"
	retailSales          string = "RETAIL_SALES"
	realGdp              string = "REAL_GDP"
	realGdpPerCapita     string = "REAL_GDP_PER_CAPITA"
	federalFundsRate     string = "FEDERAL_FUNDS_RATE"
	durableGoodsOrders   string = "DURABLES"
	unemployment         string = "UNEMPLOYMENT"
	nonfarmPayroll       string = "NONFARM_PAYROLL"
	inflation            string = "INFLATION"
	inflationExpectation string = "INFLATION_EXPECTATION"
)

type ReportType int

const (
	CPI ReportType = iota
	CONSUMER_SENTIMENT
	TREASURY_YIELD
	RETAIL_SALES
	REAL_GDP
	REAL_GDP_PER_CAPITA
	FED_FUNDS_RATE
	DURABLE_GOODS
	UNEMPLOYMENT
	NONFARM_PAYROLL
	INFLATION
	INFLATION_EXPECTATION
)

func getApiFunction(reportType ReportType) string {
	switch reportType {
	case CPI:
		return cpi
	case CONSUMER_SENTIMENT:
		return consumerSentiment
	case TREASURY_YIELD:
		return treasuryYield
	case RETAIL_SALES:
		return retailSales
	case REAL_GDP:
		return realGdp
	case REAL_GDP_PER_CAPITA:
		return realGdpPerCapita
	case FED_FUNDS_RATE:
		return federalFundsRate
	case DURABLE_GOODS:
		return durableGoodsOrders
	case UNEMPLOYMENT:
		return unemployment
	case NONFARM_PAYROLL:
		return nonfarmPayroll
	case INFLATION:
		return inflation
	case INFLATION_EXPECTATION:
		return inflationExpectation
	default:
		return "unknown"
	}
}

type Options struct {
	Interval Interval `url:"interval"`
	Maturity Maturity `url:"maturity"`
}

// EconomicData Gets the monthly manufacturers' new orders of durable goods in the United States.
func (c *Client) EconomicData(ctx context.Context, reportType ReportType, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, getApiFunction(reportType), opts)
}

func createAndSend(ctx context.Context, c *Client, function string, opts *Options) (*data.EconomicResponse, error) {
	endpoint, err := createEndpoint(c.cfg.baseUrl, c.cfg.token, function, opts)
	if err != nil {
		return nil, err
	}

	responseData := &data.EconomicResponse{}
	err = GetAndDecode(ctx, c.httpClient, endpoint, responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func createEndpoint(baseUrl string, token string, function string, opts *Options) (string, error) {
	if strings.TrimSpace(baseUrl) == "" {
		return "", errors.New("baseUrl cannot be empty")
	}
	if strings.TrimSpace(token) == "" {
		return "", errors.New("token cannot be empty")
	}
	if strings.TrimSpace(function) == "" {
		return "", errors.New("function cannot be empty")
	}
	if opts == nil {
		opts = &Options{Interval: 0, Maturity: 0}
	}

	params := struct {
		Function string   `url:"function"`
		Interval Interval `url:"interval,omitempty"`
		Maturity Maturity `url:"maturity,omitempty"`
		Token    string   `url:"apikey"`
	}{
		Function: function,
		Interval: opts.Interval,
		Maturity: opts.Maturity,
		Token:    token,
	}
	v, _ := query.Values(&params)

	return strings.TrimSpace(fmt.Sprintf("%s?%s", baseUrl, v.Encode())), nil
}
