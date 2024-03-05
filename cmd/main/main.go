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

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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
	r.HandleFunc("/ping", pingPongHandler).Methods(http.MethodGet)

	authRepo := authR.NewRepository(db)
	authUsecase := authUc.NewAuthUsecase(authRepo)
	autHandler := authH.NewAuthHandler(authUsecase)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", autHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	auth.HandleFunc("/login", autHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/logout", middleware.JwtMiddleware(http.HandlerFunc(autHandler.Logout))).Methods(http.MethodGet, http.MethodOptions)

	companyRepo := companyR.NewRepository(db)
	companyUsecase := companyUc.NewCompanyUsecase(companyRepo)
	companyHandler := companyH.NewCompanyHandler(companyUsecase)

	companyApi := r.PathPrefix("/company").Subrouter()
	companyApi.Handle("/create", middleware.JwtMiddleware(http.HandlerFunc(companyHandler.CreateCompany))).Methods(http.MethodPost)
	companyApi.HandleFunc("/get/by/id", companyHandler.GetCompanyById).Methods(http.MethodGet)
	companyApi.HandleFunc("/get/list", companyHandler.GetCompaniesList).Methods(http.MethodGet)
	companyApi.Handle("/delete/by/id", middleware.JwtMiddleware(http.HandlerFunc(companyHandler.DeleteCompanyById))).Methods(http.MethodDelete, http.MethodPost)
	companyApi.Handle("/update/by/id", middleware.JwtMiddleware(http.HandlerFunc(companyHandler.UpdateCompanyById))).Methods(http.MethodPost, http.MethodPut)

	buildingRepo := buildingR.NewRepository(db)
	buildingUsecase := buildingUc.NewBuildingUsecase(buildingRepo)
	buildingHandler := buildingH.NewBuildingHandler(buildingUsecase)

	buildingApi := r.PathPrefix("/building").Subrouter()
	buildingApi.Handle("/create", middleware.JwtMiddleware(http.HandlerFunc(buildingHandler.CreateBuilding))).Methods(http.MethodPost)
	buildingApi.HandleFunc("/get/by/id", buildingHandler.GetBuildingById).Methods(http.MethodGet)
	buildingApi.HandleFunc("/get/list", buildingHandler.GetBuildingsList).Methods(http.MethodGet)
	buildingApi.Handle("/delete/by/id", middleware.JwtMiddleware(http.HandlerFunc(buildingHandler.DeleteBuildingById))).Methods(http.MethodDelete, http.MethodPost)
	buildingApi.Handle("/update/by/id", middleware.JwtMiddleware(http.HandlerFunc(buildingHandler.UpdateBuildingById))).Methods(http.MethodPost, http.MethodPut)

	advertRepo := advertR.NewRepository(db)
	advertUsecase := advertUc.NewAdvertUsecase(advertRepo)
	advertHandler := advertH.NewAdvertHandler(advertUsecase)

	advertApi := r.PathPrefix("/advert").Subrouter()
	advertApi.Handle("/create", middleware.JwtMiddleware(http.HandlerFunc(advertHandler.CreateAdvert))).Methods(http.MethodPost)
	advertApi.HandleFunc("/get/by/id", advertHandler.GetAdvertById).Methods(http.MethodGet)
	advertApi.HandleFunc("/get/list", advertHandler.GetAdvertsList).Methods(http.MethodGet)
	advertApi.Handle("/delete/by/id", middleware.JwtMiddleware(http.HandlerFunc(advertHandler.DeleteAdvertById))).Methods(http.MethodDelete, http.MethodPost)
	advertApi.Handle("/update/by/id", middleware.JwtMiddleware(http.HandlerFunc(advertHandler.UpdateAdvertById))).Methods(http.MethodPost, http.MethodPut)

	imageRepo := imageR.NewRepository(db)
	imageUsecase := imageUc.NewImageUsecase(imageRepo)
	imageHandler := imageH.NewImageHandler(imageUsecase)

	imageApi := r.PathPrefix("/image").Subrouter()
	imageApi.Handle("/create", middleware.JwtMiddleware(http.HandlerFunc(imageHandler.CreateImage))).Methods(http.MethodPost)
	imageApi.HandleFunc("/get/list/by/advert/id", imageHandler.GetImagesByAdvertId).Methods(http.MethodGet)
	imageApi.Handle("/delete/by/id", middleware.JwtMiddleware(http.HandlerFunc(imageHandler.DeleteImageById))).Methods(http.MethodDelete, http.MethodPost)

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
