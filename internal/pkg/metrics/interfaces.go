package metrics

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type MetricsHTTP interface {
	IncreaseHits(status, method, path string)
	AddDurationToHandlerTimings(path, method string, duration time.Duration)
	AddDurationToMicroserviceTimings(mcrserviceMethod string, duration time.Duration)
	AddDurationToQueryTimings(repoMethod, queryName string, duration time.Duration)
	IncreaseExtSystemErr(systemName, errorType string)
	ServerMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	MetricsMiddleware(next http.Handler, replacePos int, altName string) http.Handler
	RegisterMetrics()
}
