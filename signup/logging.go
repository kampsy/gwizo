package signup

import (
	"time"

	"github.com/go-kit/kit/log"
)

// loggingMiddleware ...
type loggingMiddleware struct {
	logger log.Logger
	next   Servicer
}

// Signin ...
func (mw loggingMiddleware) Signup(FirstName,LastName, Phone_number, Password,Pin,Username, User_id,  Email, Balance , User_type_name string, Status_id, Account_type_id, Account_id, User_type_id int ) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "signup",
			"username", Username,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Signup(FirstName,LastName, Phone_number, Password,Pin,Username, User_id,Email, Balance , User_type_name, Status_id, Account_type_id, Account_id, User_type_id ) // i am having an error here 
	return
}