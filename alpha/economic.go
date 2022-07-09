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
)

type Options struct {
	Interval Interval `url:"interval"`
	Maturity Maturity `url:"maturity"`
}

func (c *Client) TreasuryYield(opts *Options) (*data.EconomicResponse, error) {
	endpoint, err := createEndpoint(c.cfg.baseUrl, c.cfg.token, treasuryYield, opts)
	if err != nil {
		return nil, err
	}
	return get(c, endpoint)
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
	endpoint, err := createEndpoint(c.cfg.baseUrl, c.cfg.token, cpi, opts)
	if err != nil {
		return nil, err
	}
	return get(c, endpoint)
}

func (c *Client) ConsumerSentiment(opts *Options) (*data.EconomicResponse, error) {
	endpoint, err := createEndpoint(c.cfg.baseUrl, c.cfg.token, consumerSentiment, opts)
	if err != nil {
		return nil, err
	}
	return get(c, endpoint)
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

	fmt.Println(params)

	v, _ := query.Values(&params)

	fmt.Println(v)
	return fmt.Sprintf("%s?%s", baseUrl, v.Encode()), nil
}
