// package search adds search functionality to the wallet. search for people within dazwallet
package search

import (
	"context"
	"encoding/json"
	"net/http"
)

// decodeSerchRequest get payload from the client
func decodeSearchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request searchRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

// encodeSerchResponse send payload to the client
func encodeSerchResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
