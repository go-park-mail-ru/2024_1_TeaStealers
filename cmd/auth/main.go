package main

import (
	genAuth "2024_1_TeaStealers/internal/pkg/auth/delivery/grpc/gen"
	authR "2024_1_TeaStealers/internal/pkg/auth/repo"
	authUc "2024_1_TeaStealers/internal/pkg/auth/usecase"

	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcAuth "2024_1_TeaStealers/internal/pkg/auth/delivery/grpc"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	// здесь сетап для слоев

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

	authRepo := authR.NewRepository(db, logger)
	authUsecase := authUc.NewAuthUsecase(authRepo, logger)
	// слой grpc
	authHandler := grpcAuth.NewServerAuthHandler(authUsecase, logger)
	gRPCServer := grpc.NewServer()
	// регистрация сервиса
	genAuth.RegisterAuthServer(gRPCServer, authHandler)

	// запуск grpc сервиса
	go func() {
		listener, err := net.Listen("tcp", "PORT") // todo порт
		if err != nil {
			// log.Error(errr)
		}
		if err := gRPCServer.Serve(listener); err != nil {
			// log.Error(errr)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	gRPCServer.GracefulStop()
	return nil
}
