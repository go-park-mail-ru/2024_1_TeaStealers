package main

import (
	genUsers "2024_1_TeaStealers/internal/pkg/users/delivery/grpc/gen"
	UsersR "2024_1_TeaStealers/internal/pkg/users/repo"
	UsersUc "2024_1_TeaStealers/internal/pkg/users/usecase"

	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcUsers "2024_1_TeaStealers/internal/pkg/users/delivery/grpc"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	_ = godotenv.Load()
	// logger := zap.Must(zap.NewDevelopment())

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

	usersRepo := UsersR.NewRepository(db)
	usersUsecase := UsersUc.NewUserUsecase(usersRepo)
	usersHandler := grpcUsers.NewUserServerHandler(usersUsecase)
	gRPCServer := grpc.NewServer()
	genUsers.RegisterUsersServer(gRPCServer, usersHandler)

	go func() {
		listener, err := net.Listen("tcp", GateWayPort)
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
