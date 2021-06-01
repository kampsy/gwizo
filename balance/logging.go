package balance

import (
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Servicer
}

// Balance ...
func (mw LoggingMiddleware) Balance(userid string) (output string, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "signout",
			"token", userid,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Balance(userid)
	return
}
