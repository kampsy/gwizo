package balance

import (
	"context"
	"dazwallet/auth"
	"encoding/json"
	"net/http"
	"strings"
)

// decodeBalanceRequest ...
func decodeBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request balanceRequest

	// Todo middleware will parse a token passed into the context
	token := r.Header.Get("Authorization")
	ls := strings.Split(token, " ")
	userid, err := auth.GetUserIDFromToken(ls[1])
	if err != nil {
		return "", err
	}

	request.UserID = userid
	return request, nil
}

// encodeBalanceResponse ...
func encodeBalanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
