package main

import (
	advertH "2024_1_TeaStealers/internal/pkg/adverts/delivery"
	advertR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	authH "2024_1_TeaStealers/internal/pkg/auth/delivery"
	authR "2024_1_TeaStealers/internal/pkg/auth/repo"
	authUc "2024_1_TeaStealers/internal/pkg/auth/usecase"
	buildingH "2024_1_TeaStealers/internal/pkg/buildings/delivery"
	buildingR "2024_1_TeaStealers/internal/pkg/buildings/repo"
	buildingUc "2024_1_TeaStealers/internal/pkg/buildings/usecase"
	companyH "2024_1_TeaStealers/internal/pkg/companies/delivery"
	companyR "2024_1_TeaStealers/internal/pkg/companies/repo"
	companyUc "2024_1_TeaStealers/internal/pkg/companies/usecase"
	imageH "2024_1_TeaStealers/internal/pkg/images/delivery"
	imageR "2024_1_TeaStealers/internal/pkg/images/repo"
	imageUc "2024_1_TeaStealers/internal/pkg/images/usecase"
	"2024_1_TeaStealers/internal/pkg/middleware"
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
	auth.Handle("/logout", middleware.JwtMiddleware(http.HandlerFunc(autHandler.Logout))).Methods(http.MethodGet, http.MethodOptions)
	auth.HandleFunc("/check_auth", autHandler.CheckAuth).Methods(http.MethodGet, http.MethodOptions)

	companyRepo := companyR.NewRepository(db)
	companyUsecase := companyUc.NewCompanyUsecase(companyRepo)
	companyHandler := companyH.NewCompanyHandler(companyUsecase)

	companyApi := r.PathPrefix("/companies").Subrouter()
	companyApi.HandleFunc("/", companyHandler.CreateCompany).Methods(http.MethodPost, http.MethodOptions)
	companyApi.HandleFunc("/{id}", companyHandler.GetCompanyById).Methods(http.MethodGet, http.MethodOptions)
	companyApi.HandleFunc("/list/", companyHandler.GetCompaniesList).Methods(http.MethodGet, http.MethodOptions)
	companyApi.HandleFunc("/{id}", companyHandler.DeleteCompanyById).Methods(http.MethodDelete, http.MethodOptions)
	companyApi.HandleFunc("/{id}", companyHandler.UpdateCompanyById).Methods(http.MethodPost, http.MethodOptions)

	buildingRepo := buildingR.NewRepository(db)
	buildingUsecase := buildingUc.NewBuildingUsecase(buildingRepo)
	buildingHandler := buildingH.NewBuildingHandler(buildingUsecase)

	buildingApi := r.PathPrefix("/buildings").Subrouter()
	buildingApi.HandleFunc("/", buildingHandler.CreateBuilding).Methods(http.MethodPost, http.MethodOptions)
	buildingApi.HandleFunc("/{id}", buildingHandler.GetBuildingById).Methods(http.MethodGet, http.MethodOptions)
	buildingApi.HandleFunc("/list/", buildingHandler.GetBuildingsList).Methods(http.MethodGet, http.MethodOptions)
	buildingApi.HandleFunc("/{id}", buildingHandler.DeleteBuildingById).Methods(http.MethodDelete, http.MethodOptions)
	buildingApi.HandleFunc("/{id}", buildingHandler.UpdateBuildingById).Methods(http.MethodPost, http.MethodOptions)

	advertRepo := advertR.NewRepository(db)
	advertUsecase := advertUc.NewAdvertUsecase(advertRepo)
	advertHandler := advertH.NewAdvertHandler(advertUsecase)

	advertApi := r.PathPrefix("/adverts").Subrouter()
	advertApi.HandleFunc("/", advertHandler.CreateAdvert).Methods(http.MethodPost, http.MethodOptions)
	advertApi.HandleFunc("/{id}", advertHandler.GetAdvertById).Methods(http.MethodGet, http.MethodOptions)
	advertApi.HandleFunc("/list/", advertHandler.GetAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advertApi.HandleFunc("/{id}", advertHandler.DeleteAdvertById).Methods(http.MethodDelete, http.MethodOptions)
	advertApi.HandleFunc("/{id}", advertHandler.UpdateAdvertById).Methods(http.MethodPost, http.MethodOptions)

	imageRepo := imageR.NewRepository(db)
	imageUsecase := imageUc.NewImageUsecase(imageRepo)
	imageHandler := imageH.NewImageHandler(imageUsecase)

	imageApi := r.PathPrefix("/images").Subrouter()
	imageApi.HandleFunc("/", imageHandler.CreateImage).Methods(http.MethodPost, http.MethodOptions)
	imageApi.HandleFunc("/list/by/{advert_id}", imageHandler.GetImagesByAdvertId).Methods(http.MethodGet, http.MethodOptions)
	imageApi.HandleFunc("/{id}", imageHandler.DeleteImageById).Methods(http.MethodDelete, http.MethodOptions)

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
