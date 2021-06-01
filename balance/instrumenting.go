package balance

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

func (mw instrumentingMiddleware) Balance(userid string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"service", "balance", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Balance(userid)
	return
}
