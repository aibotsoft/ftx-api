package ftxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetFutureIndexWeightsService struct {
	c         *Client
	indexName string
}

type FutureIndexWeights map[string]float64

type FutureIndexWeightsResponse struct {
	basicResponse
	Result FutureIndexWeights `json:"result"`
}

func (s *GetFutureIndexWeightsService) Do(ctx context.Context) (FutureIndexWeights, error) {
	r := newRequest(http.MethodGet, endPointWithFormat("/indexes/%s/weights", s.indexName), false)
	byteData, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	var result FutureIndexWeightsResponse
	if err := json.Unmarshal(byteData, &result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Error)
	}
	return result.Result, nil
}
