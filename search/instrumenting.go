// package search adds search functionality to the wallet. search for people within dazwallet
package search

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

func (mw instrumentingMiddleware) Search(query string) (output []data, err error) {
	defer func(begin time.Time) {
		lvs := []string{"service", "search", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Search(query)
	return
}
