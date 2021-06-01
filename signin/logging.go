package signin

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
func (mw loggingMiddleware) Signin(username, passsword string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "signin",
			"username", username,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Signin(username, passsword)
	return
}
