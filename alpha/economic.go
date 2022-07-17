package alpha

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/mhamm84/gofinance-alpha/alpha/data"
	"net/http"
)

const (
	cpi               string = "CPI"
	consumerSentiment string = "CONSUMER_SENTIMENT"
	treasuryYield     string = "TREASURY_YIELD"
	retailSales       string = "RETAIL_SALES"
)

type Options struct {
	Interval Interval `url:"interval"`
	Maturity Maturity `url:"maturity"`
}

func (c *Client) RetailSales(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, retailSales, opts)
}

func (c *Client) TreasuryYield(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, treasuryYield, opts)
}

func (c *Client) Cpi(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	if opts == nil {
		opts = &Options{
			Interval: Monthly,
		}
	}
	if opts.Interval != Monthly && opts.Interval != SemiAnnual {
		opts.Interval = Monthly
	}
	return createAndSend(ctx, c, cpi, opts)
}

func (c *Client) ConsumerSentiment(ctx context.Context, opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(ctx, c, consumerSentiment, opts)
}

func get(ctx context.Context, c *Client, endpoint string) (*data.EconomicResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	httpRes, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := httpRes.Body.Close()
		if err != nil {
			c.logger.PrintError(err, nil)
		}
	}()

	responseData := &data.EconomicResponse{}
	decoder := json.NewDecoder(httpRes.Body)
	err = decoder.Decode(responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func createAndSend(ctx context.Context, c *Client, function string, opts *Options) (*data.EconomicResponse, error) {
	endpoint, err := createEndpoint(c.cfg.baseUrl, c.cfg.token, function, opts)
	if err != nil {
		return nil, err
	}
	return get(ctx, c, endpoint)
}

func createEndpoint(baseUrl string, token string, function string, opts *Options) (string, error) {
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

	return fmt.Sprintf("%s?%s", baseUrl, v.Encode()), nil
}
