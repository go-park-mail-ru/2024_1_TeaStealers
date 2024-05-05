package metrics

import (
	"context"
	"google.golang.org/grpc"
)

type MetricsHTTP interface {
	IncreaseHits(string, string)
	IncreaseErr(string, string, string)
	ServerMetricsInterceptor(context.Context,
		interface{},
		*grpc.UnaryServerInfo,
		grpc.UnaryHandler) (interface{}, error)
}
