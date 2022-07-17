package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mhamm84/gofinance-alpha/alpha"
)

func main() {

	var token, baseUrl string
	flag.StringVar(&token, "token", "TOKEN", "API token for Alpha Vantage")
	flag.StringVar(&baseUrl, "baseUrl", "https://www.alphavantage.co/query", "Base URL for Alpha Vantage")

	client := alpha.NewClient(baseUrl, token)

	opts := &alpha.Options{
		Interval: alpha.Daily,
		Maturity: alpha.ThreeMonth,
	}

	ctx := context.Background()
	data, err := client.TreasuryYield(ctx, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
