package signup

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type signupData struct {
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	PhoneNumber     string `json:"phonenumber"`
	Password        string `json:"password"`
	Pin             string `json:"pin"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Balance         string
	AccountTypeName string
	UserTypeName    string
}

// just adding the missing information
func addSignupData(req signupRequest) signupData {
	return signupData{
		FirstName: req.FirstName, LastName: req.LastName, PhoneNumber: req.PhoneNumber,
		Password: req.Password, Pin: req.Pin, Username: req.Username, Email: req.Email, Balance: "1000.00",
		AccountTypeName: "something about the account", UserTypeName: "private user",
	}
}

type signupRequest struct {
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname.omitempty"`
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
	Pin         string `json:"pin"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}

type signupResponse struct {
	Data data `json:"data"`
}

type data struct {
	Message string `json:"message"`
}

func makeSignupEndpoint(svc Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(signupRequest)
		// missing data
		sd := addSignupData(req)
		msg, err := svc.Signup(sd)
		if err != nil {
			return signupResponse{}, err
		}
		return signupResponse{data{msg}}, nil
	}
}
