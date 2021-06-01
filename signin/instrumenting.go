package signin

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware
type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	Next           Servicer
}

// Singin runtime behaviour
func (mw InstrumentingMiddleware) Signin(id, password string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "signin", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Signin(id, password)
	return
}
