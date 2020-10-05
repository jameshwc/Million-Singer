package prom

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "service"

var EndpointList = []string{
	"/metrics",
	"/api/users",
	"/api/users/check",
	"/api/users/register",
	"/api/users/login",
	"/api/users/check",
	"/api/game",
	"/api/game/tours",
	"/api/game/collects",
	"/api/game/songs",
	"/api/game/lyrics",
	"/swagger/",
}
var (
	labels = []string{"status", "endpoint", "method"}

	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "uptime",
			Help:      "HTTP service uptime.",
		}, nil,
	)

	reqCountPerEndpoint = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total_per_endpoint",
			Help:      "Total number of HTTP requests made per endpoint.",
		}, labels,
	)

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, nil,
	)

	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, labels,
	)

	reqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)

	respSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(uptime, reqCount, reqDuration, reqSizeBytes, respSizeBytes)
	go recordUptime()
}

// recordUptime increases service uptime per second.
func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}

// calcRequestSize returns the size of request object.
func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}

func PromMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		endpoint := c.Request.URL.Path
		method := c.Request.Method
		recordEndpoint := ""
		for i := range EndpointList {
			if strings.HasPrefix(endpoint, EndpointList[i]) && len(recordEndpoint) < len(EndpointList[i]) {
				recordEndpoint = EndpointList[i]
			}
		}
		if recordEndpoint == "" {
			recordEndpoint = "unknown"
		}
		lvs := []string{status, recordEndpoint, method}
		// TODO: filter unknown url: make a endpoint map, return unknown when key not found
		reqCount.WithLabelValues().Inc()
		reqCountPerEndpoint.WithLabelValues(lvs...).Inc()
		reqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
		reqSizeBytes.WithLabelValues(lvs...).Observe(calcRequestSize(c.Request))
		respSizeBytes.WithLabelValues(lvs...).Observe(float64(c.Writer.Size()))
	}
}
func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
