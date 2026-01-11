/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

API version: 1.0.0
*/

package models

// GetTickerResponse represents the response from the ticker endpoint.
type GetTickerResponse struct {
	Code          string     `json:"code"`
	Message       *string    `json:"message"`
	MessageDetail *string    `json:"messageDetail"`
	Success       bool       `json:"success"`
	Data          TickerData `json:"data"`
}

// TickerData represents a single ticker's data.
type TickerData struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}
