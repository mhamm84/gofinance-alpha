package alpha

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/mhamm84/gofinance-alpha/alpha/data"
)

//CPI
//This API returns the monthly and semiannual consumer price index (CPI) of the United States. CPI is widely regarded as the barometer of inflation levels in the broader economy.
//https://www.alphavantage.co/query?function=CPI&interval=monthly&apikey=demo

type CpiOptions struct {
	Interval CpiInterval `url:"interval"`
}

func (c *AlphaClient) Cpi(opts *CpiOptions) (*data.CpiResponse, error) {
	endpoint, err := createCpiEndpoint(c.cfg.baseUrl, c.cfg.token, opts)
	if err != nil {
		return nil, err
	}
	response, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			c.logger.PrintError(err, nil)
		}
	}()

	cpiResponse := &data.CpiResponse{}
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(cpiResponse)
	return cpiResponse, nil
}

func createCpiEndpoint(baseUrl string, token string, opts *CpiOptions) (string, error) {
	if opts == nil {
		opts = &CpiOptions{
			Interval: Monthly,
		}
	}
	if opts.Interval != Monthly && opts.Interval != SemiAnnual {
		opts.Interval = Monthly
	}
	params := struct {
		Function string      `url:"function"`
		Interval CpiInterval `url:"interval"`
		Token    string      `url:"apikey"`
	}{
		Function: "CPI",
		Interval: opts.Interval,
		Token:    token,
	}
	v, _ := query.Values(&params)
	return fmt.Sprintf("%s?%s", baseUrl, v.Encode()), nil
}
