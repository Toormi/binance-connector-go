package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetKlines()
}

func GetKlines() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath("https://www.binance.com"),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)

	// Get Klines data
	resp, err := apiClient.RestApi.AlphaAPI.GetKlines(context.Background()).
		Symbol("ALPHA_175USDT").
		Interval("1h").
		Limit(10).
		Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
