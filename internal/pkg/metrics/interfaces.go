package metrics

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type MetricsHTTP interface {
	IncreaseHits(string, string)
	IncreaseErr(string, string, string)
	AddDurationToHistogram(string, string, time.Duration)
	IncreaseExtSystemErr(string, string)
	ServerMetricsInterceptor(context.Context,
		interface{},
		*grpc.UnaryServerInfo,
		grpc.UnaryHandler) (interface{}, error)
	ServerMetricsMiddleware(next http.Handler, urlTruncCount int, replacePos int, altName string) http.Handler
}
