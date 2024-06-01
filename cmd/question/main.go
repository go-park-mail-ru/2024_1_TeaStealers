package main

import (
	"2024_1_TeaStealers/internal/pkg/config"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	genQuestion "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/grpc/gen"
	questionR "2024_1_TeaStealers/internal/pkg/questionnaire/repo"
	questionUc "2024_1_TeaStealers/internal/pkg/questionnaire/usecase"
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

	metricsMw "2024_1_TeaStealers/internal/pkg/metrics/middleware"
	grpcQuestion "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/grpc"

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
	maxConns := int32(5)
	cfg := config.MustLoad()
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
		Addr:              ":8094",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {

		logger.Info("Starting HTTP server for metrics on :8094")
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("HTTP server listen: %s\n", err))
		}
	}()

	questionRepo := questionR.NewRepository(logger)
	questionUsecase := questionUc.NewQuestionnaireUsecase(questionRepo, logger)
	questionHandler := grpcQuestion.NewQuestionServerHandler(questionUsecase, logger)
	metricMw := metricsMw.Create()
	metricMw.RegisterMetrics()
	go metricMw.UpdatePSS()
	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(metricMw.ServerMetricsInterceptor))
	genQuestion.RegisterQuestionServer(gRPCServer, questionHandler)

	go func() {
		logger.Info(fmt.Sprintf("Start server on %s\n", ":8084"))
		listener, err := net.Listen("tcp", ":8084")
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
	dbPool.CloseDBPool()

	return nil
}
