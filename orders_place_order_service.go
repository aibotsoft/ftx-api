package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type PlaceOrderService struct {
	c      *Client
	params PlaceOrderParams
}

type PlaceOrderParams struct {
	Market string    `json:"market"`
	Side   Side      `json:"side"`
	Price  float64   `json:"price"`
	Type   OrderType `json:"type"`
	Size   float64   `json:"size"`
	//https://help.ftx.com/hc/en-us/articles/360030802012
	//Reduce-only orders will only reduce your overall position. They will never increase your position size or open a position in the opposite direction
	ReduceOnly *bool `json:"reduceOnly,omitempty"`
	//Immediate or cancel orders are guaranteed to be the taker order when executed. If you send an IOC order that does not immediately trade, it will be cancelled
	Ioc *bool `json:"ioc,omitempty"`
	//Post only orders are guaranteed to be the maker order when executed. If a post only order would instead cross the book and take, it will be cancelled
	PostOnly *bool `json:"postOnly,omitempty"`
	//client order id
	ClientID *string `json:"clientId,omitempty"`

	RejectOnPriceBand *bool  `json:"rejectOnPriceBand,omitempty"`
	RejectAfterTs     *int64 `json:"rejectAfterTs,omitempty"`
}

func (s *PlaceOrderService) Params(params PlaceOrderParams) *PlaceOrderService {
	s.params = params
	return s
}

type PlaceOrderResponse struct {
	basicResponse
	Result *Order `json:"result"`
}

func (s *PlaceOrderService) Do(ctx context.Context) (*Order, error) {
	r := newRequest(http.MethodPost, endPointWithFormat("/orders"), true)
	body, err := json.Marshal(s.params)
	if err != nil {
		return nil, err
	}
	r.setBody(body)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(byteData))
	var result PlaceOrderResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
