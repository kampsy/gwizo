// package search adds search functionality to the wallet. search for people within dazwallet
package search

import (
	"time"

	"github.com/go-kit/kit/log"
)

// loggingMiddleware ...
type loggingMiddleware struct {
	Logger log.Logger
	Next   Servicer
}

// Signin ...
func (mw loggingMiddleware) Search(query string) (output []data, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"service", "search",
			"query", query,
			"search results", len(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Search(query)
	return
}
