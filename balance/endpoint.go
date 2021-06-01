package balance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type balanceRequest struct {
	UserID string `json:"userid"`
}

type balanceResponse struct {
	Data data `json:"data"`
}

type data struct {
	Balance string `json:"balance"`
}

// makeBalanceEndpoint ...
func makeBalanceEndpoint(svc Servicer) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(balanceRequest)
		amount, err := svc.Balance(req.UserID)

		if err != nil {
			return balanceResponse{data{""}}, nil
		}

		return balanceResponse{data{amount}}, nil
	}
}
