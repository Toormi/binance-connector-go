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
func (s *AlphaKlinesService) Do(ctx context.Context) (res [][][]interface{}, err error) {
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
	res = make([][][]interface{}, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
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
func (s *AlphaTokensService) Do(ctx context.Context) (res AlphaTokensResponse, err error) {
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
	res = make(AlphaTokensResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AlphaTokensResponse represents the response from the token list endpoint.
type AlphaTokensResponse []TokenInfo

// TokenInfo represents a single token's information.
type TokenInfo struct {
	ContractAddress string `json:"contractAddress"`
	BaseAsset       string `json:"baseAsset"`
	QuoteAsset      string `json:"quoteAsset"`
	Operator        string `json:"operator"`
	TickSize        string `json:"tickSize"`
	LotSize         string `json:"lotSize"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
}
