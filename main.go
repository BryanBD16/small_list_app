package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BryanBD16/smallListApp/list"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// MySQL DSN (username:password@tcp(host:port)/dbname)
	dsn := "root:1773548_Nype58@tcp(127.0.0.1:3306)/list_app"

	// Create a logger
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowInfo())

	// Initialize repository and service for your app
	repo, err := list.NewRepository(dsn)
	if err != nil {
		logger.Log("msg", "fatal error", "err", err)
		os.Exit(1)
	}
	defer repo.Close()
	s := list.NewService(repo)

	// Register Prometheus metrics for request counts
	getRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "get_requests_total",
		Help: "Total number of GET requests",
	})
	prometheus.MustRegister(getRequests)

	addRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "add_requests_total",
		Help: "Total number of ADD requests",
	})
	prometheus.MustRegister(addRequests)

	clearRequests := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "clear_requests_total",
		Help: "Total number of CLEAR requests",
	})
	prometheus.MustRegister(clearRequests)

	// Handler for GET /element
	http.HandleFunc("/element", func(w http.ResponseWriter, r *http.Request) {
		// Increment Prometheus metric
		getRequests.Inc()

		// Log request details with logger
		level.Info(logger).Log("msg", "Handling GET request")

		// Call the actual handler to process the request
		s.Get(w, r)
	})

	// Handler for POST /element/add
	http.HandleFunc("/element/add", func(w http.ResponseWriter, r *http.Request) {
		// Increment Prometheus metric
		addRequests.Inc()

		// Log request details with logger
		level.Info(logger).Log("msg", "Handling ADD request")

		// Call the actual handler to process the request
		s.Add(w, r)
	})

	// Handler for POST /element/clear
	http.HandleFunc("/element/clear", func(w http.ResponseWriter, r *http.Request) {
		// Increment Prometheus metric
		clearRequests.Inc()

		// Log request details with logger
		level.Info(logger).Log("msg", "Handling CLEAR request")

		// Call the actual handler to process the request
		s.Clear(w, r)
	})

	// Expose metrics on /metrics endpoint for Prometheus scraping
	http.Handle("/metrics", promhttp.Handler())

	// Start the server on port 3000
	fmt.Println("Serving on port 3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Log("msg", "fatal error", "err", err)
		os.Exit(1)
	}
}
