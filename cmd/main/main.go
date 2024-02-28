package main

import (
	"2024_1_TeaStealers/internal/pkg/auth/delivery"
	"2024_1_TeaStealers/internal/pkg/auth/repo"
	"2024_1_TeaStealers/internal/pkg/auth/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.HandleFunc("/ping", pingPongHandler).Methods(http.MethodGet)

	repo := repo.NewRepository()
	usecase := usecase.NewUsecase(repo)
	authHandler := delivery.NewHandler(usecase)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Start server on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Panic("Server shutdown failed: ", err, '\n')

	}
}

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}
