/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

API version: 1.0.0
*/

package models

// GetKlinesResponse represents the response from the klines endpoint.
type GetKlinesResponse struct {
	Code          string      `json:"code"`
	Message       *string     `json:"message"`
	MessageDetail *string     `json:"messageDetail"`
	Success       bool        `json:"success"`
	Data          []KlineData `json:"data"`
}

// KlineData represents a single kline/candlestick data point.
// Each element corresponds to:
// [0] Open time (millisecond timestamp)
// [1] Open price
// [2] High price
// [3] Low price
// [4] Close price
// [5] Volume
// [6] Close time (millisecond timestamp)
// [7] Quote asset volume
// [8] Number of trades
// [9] Taker buy base asset volume
// [10] Taker buy quote asset volume
// [11] Ignore (static value 0)
type KlineData []string
