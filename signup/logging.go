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
func (mw loggingMiddleware) Signup(req signupData) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "signup",
			"username", req.Username,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Signup(req)
	return
}
