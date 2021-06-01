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
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIN   int    `json:"expires_in"`
}

// MakeSigninEndpoint ...
func MakeSigninEndpoint(svc Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signinRequest)
		token, err := svc.Signin(req.Username, req.Password)
		if err != nil {
			// get the error message
			return signinResponse{"", token, 0}, nil
		}
		return signinResponse{token, "bearer", 3600}, nil
	}
}
