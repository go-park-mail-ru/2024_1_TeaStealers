package middleware

import (
	"2024_1_TeaStealers/internal/pkg/metrics"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type GrpcMiddleware struct {
	// totalHitsHandler   *prometheus.CounterVec todo - агрегация??
	HitsHandlerStatusCode *prometheus.CounterVec   // хиты хэндлеров с разделением по кодам
	HandlerTimings        *prometheus.HistogramVec // тайминги по хэндлерам
	microserviceTimings   *prometheus.HistogramVec // тайминги микросервисов
	queryTimings          *prometheus.HistogramVec // тайминги запросов
	extSystemErrors       *prometheus.CounterVec   // ошибки запросов / микросервисов
}

func Create() metrics.MetricsHTTP {

	labelResponseStatusHits := []string{"status_code", "path", "method"}
	HitsHandlerStatusCode := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_status_hits_handler",
		Help: "number_of_requests_by_status_code",
	}, labelResponseStatusHits)
	prometheus.MustRegister(HitsHandlerStatusCode)

	labelLatency := []string{"path"}
	HandlerTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_latency_seconds",
		Help:    "Latency of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatency)
	prometheus.MustRegister(HandlerTimings)

	labelLatencymicro := []string{"microservice", "method"}
	microserviceTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "microservice_request_latency_seconds",
		Help:    "Latency of microservice requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatencymicro)
	prometheus.MustRegister(microserviceTimings)

	labelLatencyQuery := []string{"repo_method", "query_name", "method"}
	QueryTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "query_request_latency_seconds",
		Help:    "Latency of microservice requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatencyQuery)
	prometheus.MustRegister(QueryTimings)

	labelExtSystemErrors := []string{"system_name", "error_type"}
	extSystemErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "external_system_errors_total",
		Help: "number_of_external_system_errors",
	}, labelExtSystemErrors)
	prometheus.MustRegister(extSystemErrors)

	return &GrpcMiddleware{
		HitsHandlerStatusCode: HitsHandlerStatusCode,
		HandlerTimings:        HandlerTimings,
		microserviceTimings:   microserviceTimings,
		queryTimings:          QueryTimings,
		extSystemErrors:       extSystemErrors,
	}
}

func (m *GrpcMiddleware) IncreaseHits(status, method, path string) {
	m.HitsHandlerStatusCode.WithLabelValues(status, path, method).Inc()
}

func (m *GrpcMiddleware) AddDurationToHandlerTimings(path string, duration time.Duration) {
	m.HandlerTimings.WithLabelValues(path).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) AddDurationToMicroserviceTimings(mcrService, method string, duration time.Duration) {
	m.HandlerTimings.WithLabelValues(mcrService, method).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) AddDurationToQueryTimings(repo_method, query_name, method string, duration time.Duration) {
	m.HandlerTimings.WithLabelValues(repo_method, query_name, method).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) IncreaseExtSystemErr(systemName, errorType string) {
	m.extSystemErrors.WithLabelValues(systemName, errorType).Inc()
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()

	resp, err := handler(ctx, req)

	duration := time.Since(startTime)
	status := "200"
	if err != nil {
		status = "400" // todo
	}
	m.IncreaseHits(status, "", info.FullMethod)

	// Record the duration in the histogram for handler timings
	m.AddDurationToHandlerTimings("", duration)

	return resp, err
}

func (m *GrpcMiddleware) ServerMetricsMiddleware(next http.Handler, urlTruncCount int, replacePos int, altName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// start := time.Now()
		// ww := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(w, r)
		// tm := time.Since(start)
		// methodName, err1 := utils.TruncSlash(r.URL.String(), urlTruncCount)
		// methodName, err2 := utils.ReplaceURLPart(methodName, replacePos, altName)
		// if err1 == nil && err2 == nil {
		// m.IncreaseHits(methodName, r.Method)
		// m.AddDurationToHistogram(methodName, r.Method, tm)
		// m.IncreaseErr(fmt.Sprintf("%d", ww.statusCode), r.Method, methodName)
		// }
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
