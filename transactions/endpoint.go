package transactions

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type transData struct {
	UserID string
}

type transResponse struct {
	Err  string `json:""`
	Data []data `json:"data"`
}

type data struct {
	Type   string
	Name   string
	Note   string
	Amount string
	Time   string
}

func makeTransEndpoint(svc Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transData)
		transList, err := svc.Trans(req.UserID)
		if err != nil {
			return transResponse{Err: "Unable to get transactions"}, nil
		}
		return transResponse{Data: transList}, nil
	}
}
