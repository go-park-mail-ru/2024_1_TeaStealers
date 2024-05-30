package main

import (
	genComplex "2024_1_TeaStealers/internal/pkg/complexes/delivery/grpc/gen"
	complexR "2024_1_TeaStealers/internal/pkg/complexes/repo"
	complexUc "2024_1_TeaStealers/internal/pkg/complexes/usecase"
	"2024_1_TeaStealers/internal/pkg/config"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	"context"
	"errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	grpcComplex "2024_1_TeaStealers/internal/pkg/complexes/delivery/grpc"
	metricsMw "2024_1_TeaStealers/internal/pkg/metrics/middleware"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	_ = godotenv.Load()
	logger := zap.Must(zap.NewDevelopment())

	cfg := config.MustLoad()
	maxConns := int32(10) // todo надо подобрать и объяснить
	dbPool.InitDatabasePool(fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Database.DBUser,
		cfg.Database.DBPass,
		cfg.Database.DBHost,
		cfg.Database.DBPort,
		cfg.Database.DBName), maxConns)
	pool := dbPool.GetDBPool()

	if err = pool.Ping(context.Background()); err != nil {
		log.Println("fail ping postgres")
		err = fmt.Errorf("error happened in db.Ping: %w", err)
		log.Println(err)
	}

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
	http.Handle("/", r)
	httpSrv := &http.Server{
		Addr:              ":8095",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {

		logger.Info("Starting HTTP server for metrics on :8095")
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("HTTP server listen: %s\n", err))
		}
	}()

	metricMw := metricsMw.Create()
	metricMw.RegisterMetrics()
	go metricMw.UpdatePSS()
	complexRepo := complexR.NewRepository(logger, metricMw)
	complexUsecase := complexUc.NewComplexUsecase(complexRepo, logger)
	complexHandler := grpcComplex.NewComplexServerHandler(complexUsecase, logger)

	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(metricMw.ServerMetricsInterceptor))
	genComplex.RegisterComplexServer(gRPCServer, complexHandler)

	go func() {
		logger.Info(fmt.Sprintf("Start server on %s\n", ":8085"))
		listener, err := net.Listen("tcp", ":8085")
		if err != nil {
			log.Fatal(err)
		}
		if err := gRPCServer.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	logger.Info(fmt.Sprintf("Received signal: %v\n", stop))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server shutdown failed: %s\n", err))
	}

	gRPCServer.GracefulStop()

	return nil
}
