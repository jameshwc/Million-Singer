package prom

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/stat"
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

	userCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "user_cpu_usage",
			Help:      "User CPU Usage.",
		}, nil,
	)

	systemCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "system_cpu_usage",
			Help:      "System CPU Usage.",
		}, nil,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(uptime, reqCount, reqCountPerEndpoint, userCPU, systemCPU)
	go recordUptime()
	go recordCpuUsage()
}

// recordUptime increases service uptime per second.
func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}

func recordCpuUsage() {
	prev := stat.CalCpuUsage()
	for range time.Tick(time.Second) {
		cur := stat.CalCpuUsage()
		total := float64(cur.Total - prev.Total)
		userCPU.WithLabelValues().Set(float64(cur.User-prev.User) / total * 100)
		systemCPU.WithLabelValues().Set(float64(cur.System-prev.System) / total * 100)
		prev = cur
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

		reqCount.WithLabelValues().Inc()
		reqCountPerEndpoint.WithLabelValues(lvs...).Inc()
	}
}
func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
