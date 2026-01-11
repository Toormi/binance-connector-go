/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

API version: 1.0.0
*/

package models

// GetTokensResponse represents the response from the token list endpoint.
type GetTokensResponse struct {
	Code          string      `json:"code"`
	Message       *string     `json:"message"`
	MessageDetail *string     `json:"messageDetail"`
	Data          []TokenInfo `json:"data"`
}

// TokenInfo represents a single token's information.
type TokenInfo struct {
	TokenId           string `json:"tokenId"`
	ChainId           string `json:"chainId"`
	ChainIconUrl      string `json:"chainIconUrl"`
	ChainName         string `json:"chainName"`
	ContractAddress   string `json:"contractAddress"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	IconUrl           string `json:"iconUrl"`
	Price             string `json:"price"`
	PercentChange24h  string `json:"percentChange24h"`
	Volume24h         string `json:"volume24h"`
	MarketCap         string `json:"marketCap"`
	Fdv               string `json:"fdv"`
	Liquidity         string `json:"liquidity"`
	TotalSupply       string `json:"totalSupply"`
	CirculatingSupply string `json:"circulatingSupply"`
	Holders           string `json:"holders"`
	Decimals          int    `json:"decimals"`
	ListingCex        bool   `json:"listingCex"`
	HotTag            bool   `json:"hotTag"`
	CexCoinName       string `json:"cexCoinName"`
	CanTransfer       bool   `json:"canTransfer"`
	Denomination      int    `json:"denomination"`
	Offline           bool   `json:"offline"`
	TradeDecimal      int    `json:"tradeDecimal"`
	AlphaId           string `json:"alphaId"`
	Offsell           bool   `json:"offsell"`
	PriceHigh24h      string `json:"priceHigh24h"`
	PriceLow24h       string `json:"priceLow24h"`
	Count24h          string `json:"count24h"`
	OnlineTge         bool   `json:"onlineTge"`
	OnlineAirdrop     bool   `json:"onlineAirdrop"`
	Score             int    `json:"score"`
	CexOffDisplay     bool   `json:"cexOffDisplay"`
	StockState        bool   `json:"stockState"`
	ListingTime       int64  `json:"listingTime"`
	MulPoint          int    `json:"mulPoint"`
}
