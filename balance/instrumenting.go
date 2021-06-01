package balance

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

//InstrumentingMiddleware ...
type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           Servicer
}

func (mw InstrumentingMiddleware) Balance(userid string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "balance", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Balance(userid)
	return
}
