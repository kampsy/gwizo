package transactions

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

//instrumentingMiddleware ...
type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Servicer
}

func (mw instrumentingMiddleware) Trans(userid string) (output []data, err error) {
	defer func(begin time.Time) {
		lvs := []string{"service", "transactions", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Trans(userid)
	return
}
