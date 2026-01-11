package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	GetTokens()
}

func GetTokens() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath("https://www.binance.com"),
	)
	apiClient := client.NewBinanceAlphaClient(
		client.WithRestAPI(configuration),
	)

	// Get all Alpha tokens
	resp, err := apiClient.RestApi.AlphaAPI.GetTokens(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))

	// Get tokens with limit
	respWithLimit, err := apiClient.RestApi.AlphaAPI.GetTokens(context.Background()).Limit(10).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue2, _ := json.MarshalIndent(respWithLimit.Data, "", "  ")
	log.Printf("Response with limit: %s\n", string(dataValue2))
}
