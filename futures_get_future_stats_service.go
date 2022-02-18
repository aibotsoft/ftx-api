package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type GetFutureStatsService struct {
	c          *Client
	futureName string
}

func (s *GetFutureStatsService) FutureName(futureName string) *GetFutureStatsService {
	s.futureName = futureName
	return s
}

type FutureStats struct {
	Volume                   float64   `json:"volume"`
	NextFundingRate          float64   `json:"nextFundingRate"`
	NextFundingTime          time.Time `json:"nextFundingTime"`
	ExpirationPrice          float64   `json:"expirationPrice"`
	PredictedExpirationPrice float64   `json:"predictedExpirationPrice"`
	StrikePrice              float64   `json:"strikePrice"`
	OpenInterest             float64   `json:"openInterest"`
}

type FutureStatsResponse struct {
	basicResponse
	Result *FutureStats `json:"result"`
}

func (s *GetFutureStatsService) Do(ctx context.Context) (*FutureStats, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/futures/%s/stats", s.futureName), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result FutureStatsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
