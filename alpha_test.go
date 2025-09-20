package binance_connector

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type alphaTestSuite struct {
	baseTestSuite
}

func TestAlpha(t *testing.T) {
	suite.Run(t, new(alphaTestSuite))
}

func (s *alphaTestSuite) TestAlphaTokens() {
	data := []byte(`{
		"code": "000000",
		"message": null,
		"messageDetail": null,
		"data": [{
			"tokenId": "39883516E77864312AB4CE3A0BDAF1E4",
			"chainId": "56",
			"chainIconUrl": "https://bin.bnbstatic.com/image/admin_mgs_image_upload/20250228/d0216ce4-a3e9-4bda-8937-4a6aa943ccf2.png",
			"chainName": "BSC",
			"contractAddress": "0xd5df4d260d7a0145f655bcbf3b398076f21016c7",
			"name": "Ark Of Panda",
			"symbol": "AOP",
			"iconUrl": "https://bin.bnbstatic.com/images/web3-data/public/token/logos/849a8d47df8a49ff93f01f8e3241f990.png",
			"price": "0.05299655468651250336",
			"percentChange24h": "54.01",
			"volume24h": "55597436.742862718555984707547",
			"marketCap": "14362066.32004488",
			"fdv": "105993109.373025",
			"liquidity": "1965243.6703320750308283",
			"totalSupply": "2000000000",
			"circulatingSupply": "271000000",
			"holders": "7108",
			"decimals": 18,
			"listingCex": false,
			"hotTag": true,
			"cexCoinName": "",
			"canTransfer": false,
			"denomination": 1,
			"offline": false,
			"tradeDecimal": 8,
			"alphaId": "ALPHA_382",
			"offsell": false,
			"priceHigh24h": "0.07615578004512016091",
			"priceLow24h": "0.03443426501172903548",
			"count24h": "75739",
			"onlineTge": false,
			"onlineAirdrop": true,
			"score": 4411,
			"cexOffDisplay": false,
			"stockState": false,
			"listingTime": 1758283200000,
			"mulPoint": 4
		}]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewAlphaTokensService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("000000", res.Code)
	s.r().Nil(res.Message)
	s.r().Nil(res.MessageDetail)
	s.r().Len(res.Data, 1)

	token := res.Data[0]
	s.r().Equal("39883516E77864312AB4CE3A0BDAF1E4", token.TokenId)
	s.r().Equal("56", token.ChainId)
	s.r().Equal("BSC", token.ChainName)
	s.r().Equal("0xd5df4d260d7a0145f655bcbf3b398076f21016c7", token.ContractAddress)
	s.r().Equal("Ark Of Panda", token.Name)
	s.r().Equal("AOP", token.Symbol)
	s.r().Equal("0.05299655468651250336", token.Price)
	s.r().Equal("54.01", token.PercentChange24h)
	s.r().Equal("55597436.742862718555984707547", token.Volume24h)
	s.r().Equal("14362066.32004488", token.MarketCap)
	s.r().Equal(18, token.Decimals)
	s.r().Equal(false, token.ListingCex)
	s.r().Equal(true, token.HotTag)
	s.r().Equal("ALPHA_382", token.AlphaId)
	s.r().Equal(int64(1758283200000), token.ListingTime)
	s.r().Equal(4, token.MulPoint)
}

func (s *alphaTestSuite) TestAlphaTokensWithLimit() {
	data := []byte(`{
		"code": "000000",
		"message": null,
		"messageDetail": null,
		"data": []
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	limit := 10
	res, err := s.client.NewAlphaTokensService().Limit(limit).Do(newContext())
	s.r().NoError(err)
	s.r().Equal("000000", res.Code)
	s.r().Len(res.Data, 0)
}

func (s *alphaTestSuite) TestAlphaTokensError() {
	data := []byte(`{
		"code": "400001",
		"message": "Invalid request",
		"messageDetail": "Request parameters are invalid",
		"data": null
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewAlphaTokensService().Do(newContext())
	s.r().NoError(err)
	s.r().Equal("400001", res.Code)
	s.r().NotNil(res.Message)
	s.r().Equal("Invalid request", *res.Message)
	s.r().NotNil(res.MessageDetail)
	s.r().Equal("Request parameters are invalid", *res.MessageDetail)
	s.r().Nil(res.Data)
}

func (s *alphaTestSuite) TestAlphaTokensJSONUnmarshalError() {
	data := []byte(`invalid json`)
	s.mockDo(data, nil)
	defer s.assertDo()

	_, err := s.client.NewAlphaTokensService().Do(newContext())
	s.r().Error(err)
}

func (s *alphaTestSuite) TestAlphaKlines() {
	data := []byte(`{
		"code": "000000",
		"message": null,
		"messageDetail": null,
		"success": true,
		"data": [
			[
				"1758348000000",
				"0.00401127",
				"0.00401127",
				"0.00397600",
				"0.00397656",
				"435815.48000000",
				"1758351599999",
				"1732.80451313",
				"19",
				"435815.48000000",
				"1732.80451313",
				"0"
			],
			[
				"1758351600000",
				"0.00397656",
				"0.00397656",
				"0.00397656",
				"0.00397656",
				"0.00000000",
				"1758355199999",
				"0.00000000",
				"0",
				"0.00000000",
				"0.00000000",
				"0"
			]
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	interval := "1m"
	res, err := s.client.NewAlphaKlinesService().
		Symbol(symbol).
		Interval(interval).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("000000", res.Code)
	s.r().Nil(res.Message)
	s.r().Nil(res.MessageDetail)
	s.r().Equal(true, res.Success)
	s.r().Len(res.Data, 2)

	// Validate first kline data
	kline1 := res.Data[0]
	s.r().Len(kline1, 12)
	s.r().Equal("1758348000000", kline1[0])   // Open time
	s.r().Equal("0.00401127", kline1[1])      // Open price
	s.r().Equal("0.00401127", kline1[2])      // High price
	s.r().Equal("0.00397600", kline1[3])      // Low price
	s.r().Equal("0.00397656", kline1[4])      // Close price
	s.r().Equal("435815.48000000", kline1[5]) // Volume
	s.r().Equal("1758351599999", kline1[6])   // Close time
	s.r().Equal("1732.80451313", kline1[7])   // Quote asset volume
	s.r().Equal("19", kline1[8])              // Number of trades
	s.r().Equal("435815.48000000", kline1[9]) // Taker buy base asset volume
	s.r().Equal("1732.80451313", kline1[10])  // Taker buy quote asset volume
	s.r().Equal("0", kline1[11])              // Ignore field
}

func (s *alphaTestSuite) TestAlphaKlinesWithAllParams() {
	data := []byte(`{
		"code": "000000",
		"message": null,
		"messageDetail": null,
		"success": true,
		"data": [
			[
				"1609459200000",
				"29600.00000000",
				"29600.00000000",
				"29600.00000000",
				"29600.00000000",
				"0.00000000",
				"1609459259999",
				"0.00000000",
				"0",
				"0.00000000",
				"0.00000000",
				"0"
			]
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	interval := "1m"
	startTime := int64(1609459200000)
	endTime := int64(1609459260000)
	limit := 500

	res, err := s.client.NewAlphaKlinesService().
		Symbol(symbol).
		Interval(interval).
		StartTime(startTime).
		EndTime(endTime).
		Limit(limit).
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("000000", res.Code)
	s.r().Equal(true, res.Success)
	s.r().Len(res.Data, 1)
}

func (s *alphaTestSuite) TestAlphaKlinesError() {
	data := []byte(`{
		"code": "400001",
		"message": "Invalid symbol",
		"messageDetail": "The symbol parameter is invalid",
		"success": false,
		"data": null
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewAlphaKlinesService().
		Symbol("INVALID").
		Interval("1m").
		Do(newContext())
	s.r().NoError(err)
	s.r().Equal("400001", res.Code)
	s.r().NotNil(res.Message)
	s.r().Equal("Invalid symbol", *res.Message)
	s.r().NotNil(res.MessageDetail)
	s.r().Equal("The symbol parameter is invalid", *res.MessageDetail)
	s.r().Equal(false, res.Success)
	s.r().Nil(res.Data)
}

func (s *alphaTestSuite) TestAlphaKlinesJSONUnmarshalError() {
	data := []byte(`invalid json`)
	s.mockDo(data, nil)
	defer s.assertDo()

	_, err := s.client.NewAlphaKlinesService().
		Symbol("BTCUSDT").
		Interval("1m").
		Do(newContext())
	s.r().Error(err)
}
