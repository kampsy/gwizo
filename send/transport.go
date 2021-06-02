// package send transfers money
package send

import (
	"context"
	"dazwallet/auth"
	"encoding/json"
	"net/http"
	"strings"
)

func decodeSendRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request sendRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	// Todo middleware will parse a token passed into the context
	token := r.Header.Get("Authorization")
	ls := strings.Split(token, " ")
	userid, err := auth.GetUserIDFromToken(ls[1])
	if err != nil {
		return "", err
	}

	var sendData sendData
	sendData.sender = userid
	sendData.receiver = request.UserID
	sendData.amount = request.Amount
	sendData.note = request.Note

	return sendData, nil
}

func encodeSendResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
