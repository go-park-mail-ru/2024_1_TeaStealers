package middleware

import (
	"2024_1_TeaStealers/internal/pkg/metrics"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type GrpcMiddleware struct {
	totalErrors    *prometheus.CounterVec
	totalHits      *prometheus.CounterVec
	requestLatency *prometheus.HistogramVec
}

func Create() metrics.MetricsHTTP {
	labelErrors := []string{"status_code", "path", "method"}
	totalErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_errors_total",
		Help: "number_of_all_errors",
	}, labelErrors)
	prometheus.MustRegister(totalErrors)

	labelHits := []string{"path", "method"}
	totalHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "number_of_all_requests",
	}, labelHits)
	prometheus.MustRegister(totalHits)

	labelLatency := []string{"path", "method"}
	requestLatency := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_latency_seconds",
		Help:    "Latency of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatency)
	prometheus.MustRegister(requestLatency)

	return &GrpcMiddleware{
		totalErrors:    totalErrors,
		totalHits:      totalHits,
		requestLatency: requestLatency,
	}
}

func (m *GrpcMiddleware) IncreaseHits(method, path string) {
	m.totalHits.WithLabelValues(path, method).Inc()
}

func (m *GrpcMiddleware) IncreaseErr(statusCode, method, path string) {
	m.totalErrors.WithLabelValues(statusCode, path, method).Inc()
}

func (m *GrpcMiddleware) AddDurationToHistogram(method, path string, duration time.Duration) {
	m.requestLatency.WithLabelValues(path, method).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	start := time.Now()
	h, err := handler(ctx, req)
	tm := time.Since(start)

	if err != nil { // todo - ошибка не всегда означает 500, надо подумать как различать
		m.IncreaseErr("500", info.FullMethod, "")
	}
	m.IncreaseHits(info.FullMethod, "")
	m.AddDurationToHistogram(info.FullMethod, "", tm)
	return h, err
}

func (m *GrpcMiddleware) ServerMetricsMiddleware(next http.Handler, urlTruncCount int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		next.ServeHTTP(w, r)
		tm := time.Since(start)
		methodName, err := utils.TruncSlash(r.URL.String(), urlTruncCount)
		if err == nil {
			m.IncreaseHits(methodName, "")
			m.AddDurationToHistogram(methodName, "", tm)
		}
	})
}
