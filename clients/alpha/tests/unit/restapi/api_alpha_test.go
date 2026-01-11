/*
Binance Alpha REST API TEST

Testing AlphaAPIService

*/

package binancealpharestapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	client "github.com/binance/binance-connector-go/clients/alpha"
	"github.com/binance/binance-connector-go/clients/alpha/src/restapi/models"
	"github.com/binance/binance-connector-go/common/common"
	"github.com/stretchr/testify/require"
)

func Test_binancealpharestapi_AlphaAPIService(t *testing.T) {

	t.Run("Test AlphaAPIService GetKlines Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":null,"messageDetail":null,"success":true,"data":[["1758348000000","0.00401127","0.00401127","0.00397600","0.00397656","435815.48000000","1758351599999","1732.80451313","19","435815.48000000","1732.80451313","0"]]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/klines", r.URL.Path)
			require.Equal(t, "ALPHA_175USDT", r.URL.Query().Get("symbol"))
			require.Equal(t, "1h", r.URL.Query().Get("interval"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetKlinesResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetKlines(context.Background()).Symbol("ALPHA_175USDT").Interval("1h").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetKlinesResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetKlinesResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AlphaAPIService GetKlines Missing Required Params", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetKlines(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AlphaAPIService GetKlines Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetKlines(context.Background()).Symbol("BTCUSDT").Interval("1h").Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AlphaAPIService GetTicker Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":null,"messageDetail":null,"success":true,"data":{"symbol":"ALPHA_175USDT","priceChange":"-0.00022791","priceChangePercent":"-5.885","weightedAvgPrice":"0.00384878","lastPrice":"0.00364502","lastQty":"311.97000000","openPrice":"0.00387293","highPrice":"0.00401127","lowPrice":"0.00360000","volume":"3129500.16000000","quoteVolume":"12044.75859580","openTime":1758288480000,"closeTime":1758373411267,"firstId":83038,"lastId":83221,"count":220}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/ticker", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetTickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetTicker(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetTickerResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetTickerResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AlphaAPIService GetTicker With Symbol Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":null,"messageDetail":null,"success":true,"data":{"symbol":"BTCUSDT","priceChange":"1000.0","priceChangePercent":"1.0","weightedAvgPrice":"100000.0","lastPrice":"100000.0","lastQty":"0.1","openPrice":"99000.0","highPrice":"101000.0","lowPrice":"98000.0","volume":"1000.0","quoteVolume":"100000000.0","openTime":1758348000000,"closeTime":1758351599999,"firstId":1,"lastId":1000,"count":1000}}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/alpha-trade/ticker", r.URL.Path)
			require.Equal(t, "BTCUSDT", r.URL.Query().Get("symbol"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetTickerResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetTicker(context.Background()).Symbol("BTCUSDT").Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AlphaAPIService GetTicker Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetTicker(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

	t.Run("Test AlphaAPIService GetTokens Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":null,"messageDetail":null,"data":[{"tokenId":"39883516E77864312AB4CE3A0BDAF1E4","chainId":"56","chainIconUrl":"https://bin.bnbstatic.com/image/admin_mgs_image_upload/20250228/d0216ce4-a3e9-4bda-8937-4a6aa943ccf2.png","chainName":"BSC","contractAddress":"0xd5df4d260d7a0145f655bcbf3b398076f21016c7","name":"Ark Of Panda","symbol":"AOP","iconUrl":"https://bin.bnbstatic.com/images/web3-data/public/token/logos/849a8d47df8a49ff93f01f8e3241f990.png","price":"0.05299655468651250336","percentChange24h":"54.01","volume24h":"55597436.742862718555984707547","marketCap":"14362066.32004488","fdv":"105993109.373025","liquidity":"1965243.6703320750308283","totalSupply":"2000000000","circulatingSupply":"271000000","holders":"7108","decimals":18,"listingCex":false,"hotTag":true,"cexCoinName":"","canTransfer":false,"denomination":1,"offline":false,"tradeDecimal":8,"alphaId":"ALPHA_382","offsell":false,"priceHigh24h":"0.07615578004512016091","priceLow24h":"0.03443426501172903548","count24h":"75739","onlineTge":false,"onlineAirdrop":true,"score":4411,"cexOffDisplay":false,"stockState":false,"listingTime":1758283200000,"mulPoint":4}]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list", r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetTokensResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetTokens(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(
			t,
			reflect.TypeOf(&common.RestApiResponse[models.GetTokensResponse]{}),
			reflect.TypeOf(resp),
		)
		require.Equal(t, reflect.TypeOf(models.GetTokensResponse{}), reflect.TypeOf(resp.Data))
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AlphaAPIService GetTokens With Limit Success", func(t *testing.T) {

		mockedJSON := `{"code":"000000","message":null,"messageDetail":null,"data":[]}`
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list", r.URL.Path)
			require.Equal(t, "10", r.URL.Query().Get("limit"))
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(mockedJSON))
		}))
		defer mockServer.Close()

		var expected models.GetTokensResponse
		err := json.Unmarshal([]byte(mockedJSON), &expected)
		require.NoError(t, err)

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetTokens(context.Background()).Limit(10).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, resp.Status)
		require.Equal(t, expected, resp.Data)
	})

	t.Run("Test AlphaAPIService GetTokens Server Error", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()

		configuration := common.NewConfigurationRestAPI()
		configuration.BasePath = mockServer.URL
		configuration.Retries = 1
		configuration.Backoff = 1

		apiClient := client.NewBinanceAlphaClient(
			client.WithRestAPI(configuration),
		)

		resp, err := apiClient.RestApi.AlphaAPI.GetTokens(context.Background()).Execute()

		require.Error(t, err)
		require.Nil(t, resp)
	})

}
