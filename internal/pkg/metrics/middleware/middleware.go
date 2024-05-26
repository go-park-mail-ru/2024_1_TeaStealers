package middleware

import (
	"2024_1_TeaStealers/internal/pkg/metrics"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type GrpcMiddleware struct {
	totalErrors        *prometheus.CounterVec
	totalHits          *prometheus.CounterVec
	responseStatusHits *prometheus.CounterVec
	requestTimings     *prometheus.HistogramVec
	extSystemErrors    *prometheus.CounterVec
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

	labelResponseStatusHits := []string{"status_code", "path", "method"}
	responseStatusHits := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_status_hits_total",
		Help: "number_of_requests_by_status_code",
	}, labelResponseStatusHits)
	prometheus.MustRegister(responseStatusHits)

	labelLatency := []string{"path", "method"}
	requestTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_latency_seconds",
		Help:    "Latency of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatency)
	prometheus.MustRegister(requestTimings)

	labelExtSystemErrors := []string{"system_name", "error_type"}
	extSystemErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "external_system_errors_total",
		Help: "number_of_external_system_errors",
	}, labelExtSystemErrors)
	prometheus.MustRegister(extSystemErrors)

	return &GrpcMiddleware{
		totalErrors:        totalErrors,
		totalHits:          totalHits,
		responseStatusHits: responseStatusHits,
		requestTimings:     requestTimings,
		extSystemErrors:    extSystemErrors,
	}
}

func (m *GrpcMiddleware) IncreaseHits(method, path string) {
	m.totalHits.WithLabelValues(path, method).Inc()
}

func (m *GrpcMiddleware) IncreaseErr(statusCode, method, path string) {
	m.totalErrors.WithLabelValues(statusCode, path, method).Inc()
	m.responseStatusHits.WithLabelValues(statusCode, path, method).Inc()
}

func (m *GrpcMiddleware) AddDurationToHistogram(method, path string, duration time.Duration) {
	m.requestTimings.WithLabelValues(path, method).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) IncreaseExtSystemErr(systemName, errorType string) {
	m.extSystemErrors.WithLabelValues(systemName, errorType).Inc()
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

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

func (m *GrpcMiddleware) ServerMetricsMiddleware(next http.Handler, urlTruncCount int, replacePos int, altName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		ww := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(ww, r)
		tm := time.Since(start)
		methodName, err1 := utils.TruncSlash(r.URL.String(), urlTruncCount)
		methodName, err2 := utils.ReplaceURLPart(methodName, replacePos, altName)
		if err1 == nil && err2 == nil {
			m.IncreaseHits(methodName, r.Method)
			m.AddDurationToHistogram(methodName, r.Method, tm)
			m.IncreaseErr(fmt.Sprintf("%d", ww.statusCode), r.Method, methodName)
		}
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
