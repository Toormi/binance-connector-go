package main

import (
	"context"
	"fmt"

	"github.com/binance/binance-connector-go"
)

func main() {
	// Get 24hr Ticker Price Change Statistics
	client := binance_connector.NewClient("", "")
	alphaTicker, err := client.NewAlphaTickerService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(alphaTicker))

	// Get 24hr Ticker Price Change Statistics for a specific symbol
	alphaTickerWithSymbol, err := client.NewAlphaTickerService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(alphaTickerWithSymbol))
}
