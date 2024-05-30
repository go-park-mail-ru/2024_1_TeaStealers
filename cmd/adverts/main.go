package main

import (
	genAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	advertsR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertsUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	"2024_1_TeaStealers/internal/pkg/config"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpcAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc"
	metricsMw "2024_1_TeaStealers/internal/pkg/metrics/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r.Handle("/metrics", promhttp.Handler())
	http.Handle("/", r)
	httpSrv := &http.Server{
		Addr:              ":8093",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {

		logger.Info("Starting HTTP server for metrics on :8093")
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("HTTP server listen: %s\n", err))
		}
	}()

	metricMw := metricsMw.Create()
	metricMw.RegisterMetrics()
	go metricMw.UpdatePSS()
	advertsRepo := advertsR.NewRepository(logger, metricMw)
	advertsUsecase := advertsUc.NewAdvertUsecase(advertsRepo, logger)
	authHandler := grpcAdverts.NewServerAdvertsHandler(advertsUsecase, logger)

	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(metricMw.ServerMetricsInterceptor))
	genAdverts.RegisterAdvertsServer(gRPCServer, authHandler)

	go func() {
		logger.Info("Starting gRPC server on :8083")
		listener, err := net.Listen("tcp", ":8083")
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
