package send

import (
	"time"

	"github.com/go-kit/kit/log"
)

// loggingMiddleware ...
type loggingMiddleware struct {
	logger log.Logger
	next   Servicer
}

// Signin ...
func (mw loggingMiddleware) Send(req sendData) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"service", "send",
			"sender", req.sender,
			"receiver", req.receiver,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Send(req)
	return
}
