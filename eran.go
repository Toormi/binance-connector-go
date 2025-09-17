package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Get Simple Earn Flexible Product List(USER_DATA)
const (
	simpleEarnFlexibleProductListEndpoint = "/sapi/v1/simple-earn/flexible/list"
)

type SimpleEarnFlexibleProductListService struct {
	c     *Client
	asset *string
	size  *int
}

func (s *SimpleEarnFlexibleProductListService) SetAsset(asset string) *SimpleEarnFlexibleProductListService {
	s.asset = &asset
	return s
}

func (s *SimpleEarnFlexibleProductListService) SetSize(size int) *SimpleEarnFlexibleProductListService {
	s.size = &size
	return s
}

func (s *SimpleEarnFlexibleProductListService) Do(ctx context.Context) (res *SimpleEarnFlexibleProductListResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: simpleEarnFlexibleProductListEndpoint,
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(SimpleEarnFlexibleProductListResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SimpleEarnFlexibleProductListResponse struct {
	Rows []struct {
		Asset                      string `json:"asset"`
		LatestAnnualPercentageRate string `json:"latestAnnualPercentageRate"`
		TierAnnualPercentageRate   struct {
			BTC  float64 `json:"0-5BTC"`
			BTC1 float64 `json:"5-10BTC"`
		} `json:"tierAnnualPercentageRate"`
		AirDropPercentageRate string `json:"airDropPercentageRate"`
		CanPurchase           bool   `json:"canPurchase"`
		CanRedeem             bool   `json:"canRedeem"`
		IsSoldOut             bool   `json:"isSoldOut"`
		Hot                   bool   `json:"hot"`
		MinPurchaseAmount     string `json:"minPurchaseAmount"`
		ProductId             string `json:"productId"`
		SubscriptionStartTime int64  `json:"subscriptionStartTime"`
		Status                string `json:"status"`
	} `json:"rows"`
	Total int `json:"total"`
}

// Get Flexible Product Position (USER_DATA)
const (
	flexiblePositionEndpoint = "/sapi/v1/simple-earn/flexible/position"
)

type FlexiblePositionService struct {
	c    *Client
	size *int
}

func (s *FlexiblePositionService) SetSize(size int) *FlexiblePositionService {
	s.size = &size
	return s
}

func (s *FlexiblePositionService) Do(ctx context.Context) (res *FlexiblePositionResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: flexiblePositionEndpoint,
		secType:  secTypeSigned,
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(FlexiblePositionResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type FlexiblePositionResponse struct {
	Total int `json:"total"`
	Rows  []struct {
		TotalAmount                string `json:"totalAmount"`
		LatestAnnualPercentageRate string `json:"latestAnnualPercentageRate"`
		Asset                      string `json:"asset"`
		CanRedeem                  bool   `json:"canRedeem"`
		CollateralAmount           string `json:"collateralAmount"`
		ProductId                  string `json:"productId"`
		YesterdayRealTimeRewards   string `json:"yesterdayRealTimeRewards"`
		CumulativeBonusRewards     string `json:"cumulativeBonusRewards"`
		CumulativeRealTimeRewards  string `json:"cumulativeRealTimeRewards"`
		CumulativeTotalRewards     string `json:"cumulativeTotalRewards"`
		AutoSubscribe              bool   `json:"autoSubscribe"`
	} `json:"rows"`
}

// Redeem Flexible Product(TRADE)
const (
	redeemFlexibleEndpoint = "/sapi/v1/simple-earn/flexible/redeem"
)

type RedeemFlexibleService struct {
	c         *Client
	productId string
	amount    *float64
}

func (s *RedeemFlexibleService) SetProductId(productId string) *RedeemFlexibleService {
	s.productId = productId
	return s
}

func (s *RedeemFlexibleService) SetAmount(amount float64) *RedeemFlexibleService {
	s.amount = &amount
	return s
}

func (s *RedeemFlexibleService) Do(ctx context.Context) (res *RedeemFlexibleResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: redeemFlexibleEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("productId", s.productId)
	if s.amount != nil {
		r.setParam("amount", *s.amount)
	} else {
		r.setParam("redeemAll", true)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(RedeemFlexibleResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type RedeemFlexibleResponse struct {
	RedeemId int  `json:"redeemId"`
	Success  bool `json:"success"`
}

// Subscribe Flexible Product(TRADE)
const (
	subscribeFlexibleEndpoint = "/sapi/v1/simple-earn/flexible/subscribe"
)

type SubscribeFlexibleService struct {
	c         *Client
	productId string
	amount    float64
}

func (s *SubscribeFlexibleService) SetProductId(productId string) *SubscribeFlexibleService {
	s.productId = productId
	return s
}

func (s *SubscribeFlexibleService) SetAmount(amount float64) *SubscribeFlexibleService {
	s.amount = amount
	return s
}

func (s *SubscribeFlexibleService) Do(ctx context.Context) (res *SubscribeFlexibleResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: subscribeFlexibleEndpoint,
		secType:  secTypeSigned,
	}
	r.setParam("productId", s.productId)
	r.setParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(SubscribeFlexibleResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubscribeFlexibleResponse struct {
	SubscribeId int  `json:"subscribeId"`
	Success     bool `json:"success"`
}
