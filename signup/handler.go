package signup

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func SignupHandler(db *gorm.DB, logger log.Logger, count metrics.Counter, latency metrics.Histogram) *httptransport.Server {
	var svc Servicer
	svc = service{db}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{count, latency, svc}
	handler := httptransport.NewServer(
		makeSignupEndpoint(svc),
		decodeSignupRequest,
		encodeSignupResponse,
	)

	return handler
}
