// package send send money to a user
package send

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// sendRequest ...
type sendRequest struct {
	UserID   string `json:"userid,omitempty"`
	Username string `json:"username,omitempty"`
	Amount   string `json:"amount,omitempty"`
	Note     string `json:"note"`
}

type sendResponse struct {
	Err     string `json:"error"`
	Data    data   `json:"data"`
	Message string `json:"message"`
}

type data struct {
	Amount string `json:"amount"`
}

type sendData struct {
	sender   string
	receiver string
	amount   string
	note     string
}

func makeSendEndpoint(svc Servicer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d := request.(sendData)
		_, err := svc.Send(d)
		if err != nil {
			switch err {
			case errNotNumber:
				return sendResponse{Err: err.Error()}, nil
			case errSenderNotFound:
				return sendResponse{Err: err.Error()}, nil
			case errSenderAccountNotFound:
				return sendResponse{Err: err.Error()}, nil
			case errSenderBalanceDecimal:
				return sendResponse{Err: err.Error()}, nil
			case errSenderBalanceLessThanAmount:
				return sendResponse{Err: err.Error()}, nil
			case errReceiverNotFound:
				return sendResponse{Err: err.Error()}, nil
			case errReceiverAccountNotFound:
				return sendResponse{Err: err.Error()}, nil
			case errReceiverBalanceDecimal:
				return sendResponse{Err: err.Error()}, nil
			default:
				return sendResponse{Err: "Unable to send money"}, nil
			}
		}
		return sendResponse{Message: "Money sent"}, nil
	}
}
