package transactions

import (
	"context"
	"dazwallet/auth"
	"encoding/json"
	"net/http"
	"strings"
)

func decodeTransRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var td transData

	// Todo middleware will parse a token passed into the context
	token := r.Header.Get("Authorization")
	ls := strings.Split(token, " ")
	userid, err := auth.GetUserIDFromToken(ls[1])

	if err != nil {
		return td, err
	}

	td.UserID = userid
	return td, nil
}

func encodeTransResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)

}
