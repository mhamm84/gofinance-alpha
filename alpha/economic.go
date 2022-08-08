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
	cpi               string = "CPI"
	consumerSentiment string = "CONSUMER_SENTIMENT"
	treasuryYield     string = "TREASURY_YIELD"
	retailSales       string = "RETAIL_SALES"
	realGdp           string = "REAL_GDP"
	realGdpPerCapita  string = "REAL_GDP_PER_CAPITA"
	federalFundsRate  string = "FEDERAL_FUNDS_RATE"
)

type Options struct {
	Interval Interval `url:"interval"`
	Maturity Maturity `url:"maturity"`
}

// FederalFundsRate Gets daily, weekly, and monthly federal funds rate (interest rate) of the United States.
func (c *Client) FederalFundsRate(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, federalFundsRate, opts)
}

// GdpPerCapita Gets data from the real GDP per capita of the United States
func (c *Client) GdpPerCapita(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, realGdpPerCapita, opts)
}

// Gdp gets the annual and quarterly Real GDP of the United States.
func (c *Client) Gdp(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, realGdp, opts)
}

// RetailSales Gets data from the retail sales endpoint
func (c *Client) RetailSales(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, retailSales, opts)
}

// TreasuryYield Gets data from the treasury yield endpoint
func (c *Client) TreasuryYield(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, treasuryYield, opts)
}

// Cpi Gets data from the CPI endpoint
func (c *Client) Cpi(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, cpi, opts)
}

// ConsumerSentiment Gets data from the consumer sentiment endpoint
func (c *Client) ConsumerSentiment(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, consumerSentiment, opts)
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
