package signin

import (
	"context"
	"encoding/json"
	"net/http"
)

/* Transports encode and decode
 */

func DecodeSigninRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request signinRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
