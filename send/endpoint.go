// package send send money to a user
package send

// sendRequest ...
type sendRequest struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Amount   string `json:"amount"`
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
