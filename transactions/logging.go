package transactions

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Servicer
}

// Balance ...
func (mw loggingMiddleware) Trans(userid string) (output []data, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "transaction",
			"userid", userid,
			"transactions", len(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Trans(userid)
	return
}
