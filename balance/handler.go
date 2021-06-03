package balance

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	httptransport "github.com/go-kit/kit/transport/http"
	"gorm.io/gorm"
)

func BalanceHandler(db *gorm.DB, logger log.Logger, requestCount metrics.Counter, requestLatency metrics.Histogram) *httptransport.Server {
	var svc Servicer
	svc = service{db}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, svc}
	handler := httptransport.NewServer(
		makeBalanceEndpoint(svc),
		decodeBalanceRequest,
		encodeBalanceResponse,
	)
	return handler
}
