package balance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type balanceRequest struct {
	UserID string `json:"userid"`
}

type balanceResponse struct {
	Balance string `json:"balance"`
}

// MakeBalanceEndpoint ...
func MakeBalanceEndpoint(svc Servicer) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(balanceRequest)
		amount, err := svc.Balance(req.UserID)

		if err != nil {
			return balanceResponse{""}, nil
		}

		return balanceResponse{amount}, nil
	}
}
