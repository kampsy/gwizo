package main

import (
	"dazwallet/balance"
	"dazwallet/database"
	"dazwallet/signin"
	"flag"
	"fmt"
	stdlog "log"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/joho/godotenv"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Load .env
func init() {
	err := godotenv.Load()
	if err != nil {
		stdlog.Println("Unable to load .env file")
	}
}

func httpMethodCtl(method string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	// mysql database
	host := flag.String("host", "0.0.0.0", "The ip or host name E.G. 0.0.0.0")
	port := flag.Int("port", 8080, "An open port E.G. 8080")
	mysqlStr := flag.String("mysql", "dev", "Mysql credentials used, dev or pro")
	flag.Parse()

	var dbConn string
	if *mysqlStr == "dev" {
		dbConn = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_DEV_USERNAME"), os.Getenv("DB_DEV_PASSWORD"), os.Getenv("DB_DEV_DATABASE"))
	} else if *mysqlStr == "pro" {
		dbConn = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_PRO_USERNAME"), os.Getenv("DB_PRO_PASSWORD"), os.Getenv("DB_PRO_DATABASE"))
	}
	db, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		logger.Log("method", "Critical", "error", fmt.Sprint(err != nil))
	}

	// run Database Migrations
	database.Migrate(db)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var ssvc signin.Servicer
	ssvc = signin.Service{db}
	ssvc = signin.LoggingMiddleware{logger, ssvc}
	ssvc = signin.InstrumentingMiddleware{requestCount, requestLatency, ssvc}

	signinHandler := httptransport.NewServer(
		signin.MakeSigninEndpoint(ssvc),
		signin.DecodeSigninRequest,
		signin.EncodeResponse,
	)

	var balanceSvc balance.Servicer
	balanceSvc = balance.Service{db}
	balanceSvc = balance.LoggingMiddleware{logger, balanceSvc}
	balanceSvc = balance.InstrumentingMiddleware{requestCount, requestLatency, balanceSvc}
	balanceHandler := httptransport.NewServer(
		balance.MakeBalanceEndpoint(balanceSvc),
		balance.DecodeBalanceRequest,
		balance.EncodeResponse,
	)

	http.Handle("/signin", httpMethodCtl("POST", signinHandler))
	http.Handle("/balance", balanceHandler)
	http.Handle("/metrics", httpMethodCtl("GET", promhttp.Handler()))
	fmt.Printf("Running on %s:%d\n", *host, *port)
	stdlog.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil))

}
