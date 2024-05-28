package middleware

import (
	"2024_1_TeaStealers/internal/pkg/metrics"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

type GrpcMiddleware struct {
	TotalHitsHandler      *prometheus.CounterVec   // все хиты хэндлера
	HitsHandlerStatusCode *prometheus.CounterVec   // хиты хэндлеров с разделением по кодам
	HandlerTimings        *prometheus.HistogramVec // тайминги по хэндлерам
	MicroserviceTimings   *prometheus.HistogramVec // тайминги микросервисов
	QueryTimings          *prometheus.HistogramVec // тайминги запросов
	ExtSystemErrors       *prometheus.CounterVec   // ошибки запросов / микросервисов
}

func Create() metrics.MetricsHTTP {
	labelResponseHits := []string{"path", "method"}
	totalHitsHandler := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_status_total_hits",
		Help: "number_of_method_total_hits",
	}, labelResponseHits)

	labelResponseStatusHits := []string{"status_code", "path", "method"}
	HitsHandlerStatusCode := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_status_hits_handler",
		Help: "number_of_requests_by_status_code",
	}, labelResponseStatusHits)

	labelLatency := []string{"path", "method"}
	HandlerTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_latency_seconds",
		Help:    "Latency of HTTP requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatency)

	labelLatencymicro := []string{"microservice_method"}
	microserviceTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "microservice_request_latency_seconds",
		Help:    "Latency of microservice requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatencymicro)

	labelLatencyQuery := []string{"repo_method", "query_name"}
	QueryTimings := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "query_request_latency_seconds",
		Help:    "Latency of microservice requests.",
		Buckets: prometheus.DefBuckets,
	}, labelLatencyQuery)

	labelExtSystemErrors := []string{"system_name", "error_type"}
	ExtSystemErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "external_system_errors_total",
		Help: "number_of_external_system_errors",
	}, labelExtSystemErrors)

	return &GrpcMiddleware{
		TotalHitsHandler:      totalHitsHandler,
		HitsHandlerStatusCode: HitsHandlerStatusCode,
		HandlerTimings:        HandlerTimings,
		MicroserviceTimings:   microserviceTimings,
		QueryTimings:          QueryTimings,
		ExtSystemErrors:       ExtSystemErrors,
	}
}

func (m *GrpcMiddleware) RegisterMetrics() {
	prometheus.MustRegister(m.TotalHitsHandler)
	prometheus.MustRegister(m.HitsHandlerStatusCode)
	prometheus.MustRegister(m.HandlerTimings)
	prometheus.MustRegister(m.MicroserviceTimings)
	prometheus.MustRegister(m.QueryTimings)
	prometheus.MustRegister(m.ExtSystemErrors)
}

func (m *GrpcMiddleware) IncreaseHits(status, method, path string) {
	m.TotalHitsHandler.WithLabelValues(path, method).Inc()
	m.HitsHandlerStatusCode.WithLabelValues(status, path, method).Inc()
}

func (m *GrpcMiddleware) AddDurationToHandlerTimings(path, method string, duration time.Duration) {
	m.HandlerTimings.WithLabelValues(path, method).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) AddDurationToMicroserviceTimings(mcrserviceMethod string, duration time.Duration) {
	m.MicroserviceTimings.WithLabelValues(mcrserviceMethod).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) AddDurationToQueryTimings(repoMethod, queryName string, duration time.Duration) {
	m.QueryTimings.WithLabelValues(repoMethod, queryName).Observe(duration.Seconds())
}

func (m *GrpcMiddleware) IncreaseExtSystemErr(systemName, errorType string) {
	m.ExtSystemErrors.WithLabelValues(systemName, errorType).Inc()
}

func (m *GrpcMiddleware) ServerMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	startTime := time.Now()
	resp, err := handler(ctx, req)
	duration := time.Since(startTime)

	m.AddDurationToMicroserviceTimings(info.FullMethod, duration)

	respCode, err2 := utils.GetValueFromInterface(resp, "RespCode")
	if err2 != nil {
		return resp, err
	}

	var intCode int32
	var ok bool
	if intCode, ok = respCode.(int32); !ok {
		return resp, err
	}

	strCode := strconv.Itoa(int(intCode))
	m.IncreaseHits(strCode, "", info.FullMethod)

	return resp, err
}

func (m *GrpcMiddleware) MetricsMiddleware(next http.Handler, replacePos int, altName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rwCode := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		start := time.Now()
		next.ServeHTTP(rwCode, r)
		tm := time.Since(start)

		statusCode := rwCode.statusCode

		m.IncreaseHits(strconv.Itoa(statusCode), r.Method, utils.ReplaceURLPart(r.URL.String(), altName, replacePos))
		m.AddDurationToHandlerTimings(r.URL.String(), r.Method, tm)
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
