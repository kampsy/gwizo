package send

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// instrumentingMiddleware
type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Servicer
}

// Singin runtime behaviour
func (mw instrumentingMiddleware) Send(req sendData) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"service", "send", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Send(req)
	return
}
