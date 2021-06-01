package signup

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)




type signupRequest struct{
	FirstName  		string	`json:"firstname"`
	LastName		string	`json:"lastname"`
	Phone_number 	string	`json:"phone_number"`
	Password		string	`json:"password"`
	Pin 			string	`json:"pin"`
	Username		string	`json:"username"`
	Email			string	`json:"email"`
	Balance 		string 	`json:"balance"`
	User_type_name	string 	`json:"user_type_name"`
	Status_id		int		`json:"status_id"`
	Account_type_id	int 	`json:"account_type_id"`
	Account_id		int 	`json:"account_id"`
	User_type_id 	int 	`json:"user_type_id"`
}

type signupResponse struct{
	Message string 	`json:"message"`
}



func  MakeSignupEndpoint(svc Servicer) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
		data := request.(signupRequest) 
		signupData , err := svc.Signup(data.FirstName,data.LastName,data.Phone_number,data.Password,data.Pin,data.Username,data.Email,data.Balance,data.User_type_name,data.Status_id,data.Account_type_id,data.Account_id, data.User_type_id)
			if err != nil {
				return	signupResponse{}, err 
			}
		return signupResponse{signupData}, nil
	}
}