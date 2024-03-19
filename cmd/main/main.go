package main

import (
	advertsH "2024_1_TeaStealers/internal/pkg/adverts/delivery"
	advertsR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertsUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	authH "2024_1_TeaStealers/internal/pkg/auth/delivery"
	authR "2024_1_TeaStealers/internal/pkg/auth/repo"
	authUc "2024_1_TeaStealers/internal/pkg/auth/usecase"
	imageH "2024_1_TeaStealers/internal/pkg/images/delivery/http"
	imageR "2024_1_TeaStealers/internal/pkg/images/repo"
	imageUc "2024_1_TeaStealers/internal/pkg/images/usecase"
	"2024_1_TeaStealers/internal/pkg/middleware"
	userH "2024_1_TeaStealers/internal/pkg/users/delivery"
	userR "2024_1_TeaStealers/internal/pkg/users/repo"
	userUc "2024_1_TeaStealers/internal/pkg/users/usecase"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "2024_1_TeaStealers/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Sample Project API
// @version 1.0
// @description This is a sample server Tean server.

// @host 0.0.0.0:8080
// @BasePath /api
// @schemes http https
func main() {
	_ = godotenv.Load()
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
	r.Use(middleware.CORSMiddleware)
	r.HandleFunc("/ping", pingPongHandler).Methods(http.MethodGet)
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	authRepo := authR.NewRepository(db)
	authUsecase := authUc.NewAuthUsecase(authRepo)
	autHandler := authH.NewAuthHandler(authUsecase)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", autHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	auth.HandleFunc("/login", autHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/logout", middleware.JwtMiddleware(http.HandlerFunc(autHandler.Logout), authRepo)).Methods(http.MethodGet, http.MethodOptions)
	auth.HandleFunc("/check_auth", autHandler.CheckAuth).Methods(http.MethodGet, http.MethodOptions)

	advertRepo := advertsR.NewRepository(db)
	advertUsecase := advertsUc.NewAdvertUsecase(advertRepo)
	advertHandler := advertsH.NewAdvertHandler(advertUsecase)

	imageRepo := imageR.NewRepository(db)
	imageUsecase := imageUc.NewImageUsecase(imageRepo)
	imageHandler := imageH.NewImageHandler(imageUsecase)

	advert := r.PathPrefix("/adverts").Subrouter()
	advert.HandleFunc("/{id}", advertHandler.GetAdvertById).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/houses", advertHandler.CreateHouseAdvert).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/flats", advertHandler.CreateFlatAdvert).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/squarelist/", advertHandler.GetSquareAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/houses/squarelist/", advertHandler.GetHouseSquareAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/flats/squarelist/", advertHandler.GetFlatSquareAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/flats/rectanglelist/", advertHandler.GetFlatRectangleAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/houses/rectanglelist/", advertHandler.GetHouseRectangleAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/image", imageHandler.UploadImage).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/{id}/image", imageHandler.GetAdvertImages).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/{id}/image", imageHandler.DeleteImage).Methods(http.MethodDelete, http.MethodOptions)

	userRepo := userR.NewRepository(db)
	userUsecase := userUc.NewUserUsecase(userRepo)
	userHandler := userH.NewUserHandler(userUsecase)

	user := r.PathPrefix("/user").Subrouter()
	user.Handle("/me", middleware.JwtMiddleware(http.HandlerFunc(userHandler.GetCurUser), authRepo)).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/avatar", middleware.JwtMiddleware(http.HandlerFunc(userHandler.UpdateUserPhoto), authRepo)).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/avatar", middleware.JwtMiddleware(http.HandlerFunc(userHandler.DeleteUserPhoto), authRepo)).Methods(http.MethodDelete, http.MethodOptions)
	user.Handle("/info", middleware.JwtMiddleware(http.HandlerFunc(userHandler.UpdateUserInfo), authRepo)).Methods(http.MethodPost, http.MethodOptions)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
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
