package signup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

/* Transports encode and decode
 */

func decodeSignupRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request signupRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(request)
		return nil, err
	}
	return request, nil
}

func encodeSignupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
