package ftxapi

import (
	"context"
	"encoding/json"
	"net/http"

	"errors"
)

type GetBalancesService struct {
	c *Client
}

type BalancesResponse struct {
	basicResponse
	Result []Balance `json:"result"`
}

func (s *GetBalancesService) Do(ctx context.Context) ([]Balance, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/wallet/balances"), true)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result BalancesResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
