package signin

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type signinRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type signinResponse struct {
	Data data `json:"data"`
}

type data struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIN    int    `json:"expires_in"`
}

// makeSigninEndpoint ...
func makeSigninEndpoint(svc Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signinRequest)
		token, err := svc.Signin(req.Username, req.Password)
		if err != nil {
			//
			return signinResponse{data{"", "", 0}}, nil
		}

		return signinResponse{data{token, "", 3600}}, nil
	}
}
