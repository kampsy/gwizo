package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

/*Requests and responses for our wallet
 */
type signinRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type signinResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type signoutRequest struct {
	Token string `json:"token"`
}

type signoutResponse struct {
	Message string `json:"message"`
}

/*Endpoints
 */

func makeSigninEndpoint(svc WalletService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signinRequest)
		token, err := svc.Signin(req.ID, req.Password)
		if err != nil {
			// get the error message
			return signinResponse{err.Error(), ""}, err
		}
		msg := "Signin was successfull"
		return signinResponse{msg, token}, nil
	}
}

func makeSignoutEndpoint(svc WalletService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(signoutRequest)
		msg, err := svc.Signout(req.Token)

		if err != nil {
			return signoutResponse{err.Error()}, err
		}

		return signinResponse{msg, ""}, nil
	}
}

/* Transports encode and decode
 */

func decodeSigninRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request signinRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeSignoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request signoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
