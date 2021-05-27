package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           WalletService
}

// Singin runtime behaviour
func (mw instrumentingMiddleware) Signin(id, password string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "signin", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Signin(id, password)
	return
}

func (mw instrumentingMiddleware) Signout(token string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "signout", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Signout(token)
	return
}
