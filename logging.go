package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   WalletService
}

// Signin ...
func (mw loggingMiddleware) Signin(id, passsword string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "signin",
			"userid", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Signin(id, passsword)
	return
}

// Signout ...
func (mw loggingMiddleware) Signout(token string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "signout",
			"token", token,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Signout(token)
	return
}
