package balance

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Servicer
}

// Balance ...
func (mw loggingMiddleware) Balance(userid string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "balance",
			"token", userid,
			"amount", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Balance(userid)
	return
}
