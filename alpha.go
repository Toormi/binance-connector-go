package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	alphaKlinesEndpoint = "/bapi/defi/v1/public/alpha-trade/klines"
)

// AlphaKlinesService fetches kline/candlestick data.
type AlphaKlinesService struct {
	c         *Client
	symbol    string
	interval  string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol sets the symbol parameter.
func (s *AlphaKlinesService) Symbol(symbol string) *AlphaKlinesService {
	s.symbol = symbol
	return s
}

// Interval sets the interval parameter.
func (s *AlphaKlinesService) Interval(interval string) *AlphaKlinesService {
	s.interval = interval
	return s
}

// StartTime sets the startTime parameter.
func (s *AlphaKlinesService) StartTime(startTime int64) *AlphaKlinesService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *AlphaKlinesService) EndTime(endTime int64) *AlphaKlinesService {
	s.endTime = &endTime
	return s
}

// Limit sets the limit parameter.
func (s *AlphaKlinesService) Limit(limit int) *AlphaKlinesService {
	s.limit = &limit
	return s
}

// Do sends the request and returns the kline data or an error.
func (s *AlphaKlinesService) Do(ctx context.Context) (res *AlphaKlinesResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: alphaKlinesEndpoint,
		secType:  secTypeNone,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AlphaKlinesResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

const (
	alphaTickerEndpoint = "/bapi/defi/v1/public/alpha-trade/ticker/24hr"
)

// AlphaTickerService fetches 24hr ticker price change statistics.
type AlphaTickerService struct {
	c      *Client
	symbol *string
}

// Symbol sets the symbol parameter.
func (s *AlphaTickerService) Symbol(symbol string) *AlphaTickerService {
	s.symbol = &symbol
	return s
}

// Do sends the request and returns the ticker data or an error.
func (s *AlphaTickerService) Do(ctx context.Context) (res *AlphaTickerResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: alphaTickerEndpoint,
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AlphaTickerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AlphaTickerResponse represents the response from the ticker endpoint.
type AlphaTickerResponse struct {
	Code          string            `json:"code"`
	Message       *string           `json:"message"`
	MessageDetail *string           `json:"messageDetail"`
	Success       bool              `json:"success"`
	Data          []AlphaTickerData `json:"data"`
}

// AlphaTickerData represents a single ticker's data.
type AlphaTickerData struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
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

const (
	alphaTokensEndpoint = "/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list"
)

// AlphaTokensService fetches token list.
type AlphaTokensService struct {
	c     *Client
	limit *int
}

// Limit sets the limit parameter.
func (s *AlphaTokensService) Limit(limit int) *AlphaTokensService {
	s.limit = &limit
	return s
}

// Do sends the request and returns the token list or an error.
func (s *AlphaTokensService) Do(ctx context.Context) (res *AlphaTokensResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: alphaTokensEndpoint,
		secType:  secTypeNone,
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(AlphaTokensResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AlphaTokensResponse represents the response from the token list endpoint.
type AlphaTokensResponse struct {
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

// AlphaKlinesResponse represents the response from the klines endpoint.
type AlphaKlinesResponse struct {
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
