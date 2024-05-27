package metrics

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type MetricsHTTP interface {
	IncreaseHits(status, method, path string)
	AddDurationToHandlerTimings(path string, duration time.Duration)
	AddDurationToMicroserviceTimings(mcrService, method string, duration time.Duration)
	AddDurationToQueryTimings(repoMethod, queryName, method string, duration time.Duration)
	IncreaseExtSystemErr(systemName, errorType string)
	ServerMetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	ServerMetricsMiddleware(next http.Handler, urlTruncCount int, replacePos int, altName string) http.Handler
}
