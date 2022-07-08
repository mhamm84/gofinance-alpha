package main

import (
	"fmt"
	"github.com/mhamm84/gofinance-alpha/alpha"
)

func main() {
	client := alpha.NewClient("YCVVI14X9XNBVGOO")

	data, err := client.Cpi(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
