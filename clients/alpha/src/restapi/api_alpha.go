/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

API version: 1.0.0
*/

package binancealpharestapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/binance/binance-connector-go/clients/alpha/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
)

// AlphaAPIService AlphaAPI Service
type AlphaAPIService Service

// ==================== GetKlines ====================

type ApiGetKlinesRequest struct {
	ctx        context.Context
	ApiService *AlphaAPIService
	symbol     *string
	interval   *string
	startTime  *int64
	endTime    *int64
	limit      *int32
}

func (r ApiGetKlinesRequest) Symbol(symbol string) ApiGetKlinesRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetKlinesRequest) Interval(interval string) ApiGetKlinesRequest {
	r.interval = &interval
	return r
}

func (r ApiGetKlinesRequest) StartTime(startTime int64) ApiGetKlinesRequest {
	r.startTime = &startTime
	return r
}

func (r ApiGetKlinesRequest) EndTime(endTime int64) ApiGetKlinesRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetKlinesRequest) Limit(limit int32) ApiGetKlinesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetKlinesRequest) Execute() (*common.RestApiResponse[models.GetKlinesResponse], error) {
	return r.ApiService.GetKlinesExecute(r)
}

/*
GetKlines Kline/Candlestick Data

Get /bapi/defi/v1/public/alpha-trade/klines

Get kline/candlestick bars for a symbol.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol - Symbol to query
@param interval - Kline interval (1m, 5m, 15m, 1h, 4h, 1d, etc.)
@param startTime - Start time in milliseconds
@param endTime - End time in milliseconds
@param limit - Number of klines to return. Default: 500; Maximum: 1000
@return ApiGetKlinesRequest
*/
func (a *AlphaAPIService) GetKlines(ctx context.Context) ApiGetKlinesRequest {
	return ApiGetKlinesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetKlinesResponse
func (a *AlphaAPIService) GetKlinesExecute(r ApiGetKlinesRequest) (*common.RestApiResponse[models.GetKlinesResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/alpha-trade/klines"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol == nil {
		return nil, common.ReportError("symbol is required and must be specified")
	}
	if r.interval == nil {
		return nil, common.ReportError("interval is required and must be specified")
	}

	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	common.ParameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "form", "")
	if r.startTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "form", "")
	}
	if r.endTime != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "form", "")
	}
	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.GetKlinesResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

// ==================== GetTicker ====================

type ApiGetTickerRequest struct {
	ctx        context.Context
	ApiService *AlphaAPIService
	symbol     *string
}

func (r ApiGetTickerRequest) Symbol(symbol string) ApiGetTickerRequest {
	r.symbol = &symbol
	return r
}

func (r ApiGetTickerRequest) Execute() (*common.RestApiResponse[models.GetTickerResponse], error) {
	return r.ApiService.GetTickerExecute(r)
}

/*
GetTicker 24hr Ticker Price Change Statistics

Get /bapi/defi/v1/public/alpha-trade/ticker

Get 24 hour rolling window price change statistics.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param symbol - Symbol to query (optional)
@return ApiGetTickerRequest
*/
func (a *AlphaAPIService) GetTicker(ctx context.Context) ApiGetTickerRequest {
	return ApiGetTickerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTickerResponse
func (a *AlphaAPIService) GetTickerExecute(r ApiGetTickerRequest) (*common.RestApiResponse[models.GetTickerResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/alpha-trade/ticker"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.symbol != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "symbol", r.symbol, "form", "")
	}

	resp, err := SendRequest[models.GetTickerResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}

// ==================== GetTokens ====================

type ApiGetTokensRequest struct {
	ctx        context.Context
	ApiService *AlphaAPIService
	limit      *int32
}

func (r ApiGetTokensRequest) Limit(limit int32) ApiGetTokensRequest {
	r.limit = &limit
	return r
}

func (r ApiGetTokensRequest) Execute() (*common.RestApiResponse[models.GetTokensResponse], error) {
	return r.ApiService.GetTokensExecute(r)
}

/*
GetTokens Get Token List

Get /bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list

Get all available tokens in Alpha.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param limit - Number of tokens to return (optional)
@return ApiGetTokensRequest
*/
func (a *AlphaAPIService) GetTokens(ctx context.Context) ApiGetTokensRequest {
	return ApiGetTokensRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetTokensResponse
func (a *AlphaAPIService) GetTokensExecute(r ApiGetTokensRequest) (*common.RestApiResponse[models.GetTokensResponse], error) {
	localVarHTTPMethod := http.MethodGet
	localVarPath := a.client.cfg.BasePath + "/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list"

	localVarQueryParams := url.Values{}
	localVarBodyParameters := make(map[string]interface{})

	if r.limit != nil {
		common.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}

	resp, err := SendRequest[models.GetTokensResponse](r.ctx, localVarPath, localVarHTTPMethod, localVarQueryParams, localVarBodyParameters, a.client.cfg)
	if err != nil || resp == nil {
		return nil, err
	}

	return resp, nil
}
