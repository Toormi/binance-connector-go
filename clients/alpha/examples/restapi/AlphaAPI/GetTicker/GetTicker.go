package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetTicker()
}

func GetTicker() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath("https://www.binance.com"),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)

	// Get 24hr Ticker Price Change Statistics
	resp, err := apiClient.RestApi.AlphaAPI.GetTicker(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))

	// Get 24hr Ticker Price Change Statistics for a specific symbol
	respWithSymbol, err := apiClient.RestApi.AlphaAPI.GetTicker(context.Background()).Symbol("ALPHA_175USDT").Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue2, _ := json.MarshalIndent(respWithSymbol.Data, "", "  ")
	log.Printf("Response with symbol: %s\n", string(dataValue2))
}
