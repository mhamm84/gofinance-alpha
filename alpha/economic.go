package alpha

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/mhamm84/gofinance-alpha/alpha/data"
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

func (c *Client) RetailSales(opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(c, retailSales, opts)
}

func (c *Client) TreasuryYield(opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(c, treasuryYield, opts)
}

func (c *Client) Cpi(opts *Options) (*data.EconomicResponse, error) {
	if opts == nil {
		opts = &Options{
			Interval: Monthly,
		}
	}
	if opts.Interval != Monthly && opts.Interval != SemiAnnual {
		opts.Interval = Monthly
	}
	return createAndSend(c, cpi, opts)
}

func (c *Client) ConsumerSentiment(opts *Options) (*data.EconomicResponse, error) {
	return createAndSend(c, consumerSentiment, opts)
}

func get(c *Client, endpoint string) (*data.EconomicResponse, error) {
	httpRes, err := c.httpClient.Get(endpoint)
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
	decoder.Decode(responseData)
	return responseData, nil
}

func createAndSend(c *Client, function string, opts *Options) (*data.EconomicResponse, error) {
	endpoint, err := createEndpoint(c.cfg.baseUrl, c.cfg.token, function, opts)
	if err != nil {
		return nil, err
	}
	return get(c, endpoint)
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
