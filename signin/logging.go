package signin

import (
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware ...
type LoggingMiddleware struct {
	Logger log.Logger
	Next   Servicer
}

// Signin ...
func (mw LoggingMiddleware) Signin(id, passsword string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "signin",
			"userid", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Signin(id, passsword)
	return
}
