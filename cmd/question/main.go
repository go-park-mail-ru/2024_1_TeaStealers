package main

import (
	genQuestion "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/grpc/gen"
	questionR "2024_1_TeaStealers/internal/pkg/questionnaire/repo"
	questionUc "2024_1_TeaStealers/internal/pkg/questionnaire/usecase"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"database/sql"
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

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("fail ping postgres")
		err = fmt.Errorf("error happened in db.Ping: %w", err)
		log.Println(err)
	}

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
	http.Handle("/", r)

	questionRepo := questionR.NewRepository(db, logger)
	questionUsecase := questionUc.NewQuestionnaireUsecase(questionRepo, logger)
	questionHandler := grpcQuestion.NewQuestionServerHandler(questionUsecase, logger)
	metricMw := metricsMw.Create()
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
	gRPCServer.GracefulStop()
	return nil
}
