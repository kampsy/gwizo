// package search adds search functionality to the wallet. search for people within dazwallet
package search

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// searchRequest query we get from the client.
type searchRequest struct {
	Query string `json:"query"`
}

type searchResponse struct {
	Err  string `json:"error"`
	Data []data `json:"data"`
}
type data struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	UserName    string `json:"username"`
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
	UserID      string `json:"userid"`
}

func makeSearchEndpoint(svc Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)
		searchResult, err := svc.Search(req.Query)

		if err != nil {
			return searchResponse{Err: "unxpected error"}, nil
		}

		return searchResponse{Data: searchResult}, nil
	}
}
