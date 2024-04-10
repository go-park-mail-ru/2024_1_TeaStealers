package main

import (
	advertsH "2024_1_TeaStealers/internal/pkg/adverts/delivery"
	advertsR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertsUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	authH "2024_1_TeaStealers/internal/pkg/auth/delivery"
	authR "2024_1_TeaStealers/internal/pkg/auth/repo"
	authUc "2024_1_TeaStealers/internal/pkg/auth/usecase"
	companyH "2024_1_TeaStealers/internal/pkg/companies/delivery"
	companyR "2024_1_TeaStealers/internal/pkg/companies/repo"
	companyUc "2024_1_TeaStealers/internal/pkg/companies/usecase"
	complexH "2024_1_TeaStealers/internal/pkg/complexes/delivery"
	complexR "2024_1_TeaStealers/internal/pkg/complexes/repo"
	complexUc "2024_1_TeaStealers/internal/pkg/complexes/usecase"
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
	"go.uber.org/zap"
)

// @title Sample Project API
// @version 1.0
// @description This is a sample server Tean server.

// @host 0.0.0.0:8080
// @BasePath /api
// @schemes http https
func main() {
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
	r.Use(middleware.CORSMiddleware)
	r.HandleFunc("/ping", pingPongHandler).Methods(http.MethodGet)
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	authRepo := authR.NewRepository(db)
	authUsecase := authUc.NewAuthUsecase(authRepo)
	autHandler := authH.NewAuthHandler(authUsecase)

	jwtMd := middleware.NewAuthMiddleware(authUsecase)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", autHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	auth.HandleFunc("/login", autHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/logout", jwtMd.JwtTMiddleware(http.HandlerFunc(autHandler.Logout))).Methods(http.MethodGet, http.MethodOptions)
	auth.Handle("/check_auth", jwtMd.JwtTMiddleware(http.HandlerFunc(autHandler.CheckAuth))).Methods(http.MethodGet, http.MethodOptions)

	advertRepo := advertsR.NewRepository(db, logger)
	advertUsecase := advertsUc.NewAdvertUsecase(advertRepo, logger)
	advertHandler := advertsH.NewAdvertHandler(advertUsecase, logger)

	imageRepo := imageR.NewRepository(db)
	imageUsecase := imageUc.NewImageUsecase(imageRepo)
	imageHandler := imageH.NewImageHandler(imageUsecase)

	advert := r.PathPrefix("/adverts").Subrouter()
	advert.HandleFunc("/{id}", advertHandler.GetAdvertById).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}", jwtMd.JwtTMiddleware(http.HandlerFunc(advertHandler.UpdateAdvertById))).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}", jwtMd.JwtTMiddleware(http.HandlerFunc(advertHandler.DeleteAdvertById))).Methods(http.MethodDelete, http.MethodOptions)
	advert.Handle("/houses/", jwtMd.JwtTMiddleware(http.HandlerFunc(advertHandler.CreateHouseAdvert))).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/buildings/", advertHandler.GetExistBuildingsByAddress).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/flats/", jwtMd.JwtTMiddleware(http.HandlerFunc(advertHandler.CreateFlatAdvert))).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/squarelist/", advertHandler.GetSquareAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.HandleFunc("/rectanglelist/", advertHandler.GetRectangeAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/image/", jwtMd.JwtTMiddleware(http.HandlerFunc(imageHandler.UploadImage))).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/{id}/image", imageHandler.GetAdvertImages).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}/image", jwtMd.JwtTMiddleware(http.HandlerFunc(imageHandler.DeleteImage))).Methods(http.MethodDelete, http.MethodOptions)

	userRepo := userR.NewRepository(db)
	userUsecase := userUc.NewUserUsecase(userRepo)
	userHandler := userH.NewUserHandler(userUsecase)

	user := r.PathPrefix("/users").Subrouter()
	user.Handle("/me", jwtMd.JwtTMiddleware(http.HandlerFunc(userHandler.GetCurUser))).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/avatar", jwtMd.JwtTMiddleware(http.HandlerFunc(userHandler.UpdateUserPhoto))).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/avatar", jwtMd.JwtTMiddleware(http.HandlerFunc(userHandler.DeleteUserPhoto))).Methods(http.MethodDelete, http.MethodOptions)
	user.Handle("/info", jwtMd.JwtTMiddleware(http.HandlerFunc(userHandler.UpdateUserInfo))).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/password", jwtMd.JwtTMiddleware(http.HandlerFunc(userHandler.UpdateUserPassword))).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/myadverts", jwtMd.JwtTMiddleware(http.HandlerFunc(advertHandler.GetUserAdverts))).Methods(http.MethodGet, http.MethodOptions)

	companyRepo := companyR.NewRepository(db)
	companyUsecase := companyUc.NewCompanyUsecase(companyRepo)
	companyHandler := companyH.NewCompanyHandler(companyUsecase)

	company := r.PathPrefix("/companies").Subrouter()
	company.HandleFunc("/", companyHandler.CreateCompany).Methods(http.MethodPost, http.MethodOptions)
	company.HandleFunc("/{id}", companyHandler.GetCompanyById).Methods(http.MethodGet, http.MethodOptions)
	company.HandleFunc("/images/{id}", companyHandler.UpdateCompanyPhoto).Methods(http.MethodPost, http.MethodOptions)

	complexRepo := complexR.NewRepository(db)
	complexUsecase := complexUc.NewComplexUsecase(complexRepo)
	complexHandler := complexH.NewComplexHandler(complexUsecase)

	complex := r.PathPrefix("/complexes").Subrouter()
	complex.HandleFunc("/", complexHandler.CreateComplex).Methods(http.MethodPost, http.MethodOptions)
	complex.HandleFunc("/{id}", complexHandler.GetComplexById).Methods(http.MethodGet, http.MethodOptions)
	complex.HandleFunc("/{id}/rectanglelist/", advertHandler.GetComplexAdverts).Methods(http.MethodGet, http.MethodOptions)
	complex.HandleFunc("/houses", complexHandler.CreateHouseAdvert).Methods(http.MethodPost, http.MethodOptions)
	complex.HandleFunc("/flats", complexHandler.CreateFlatAdvert).Methods(http.MethodPost, http.MethodOptions)
	complex.HandleFunc("/buildings", complexHandler.CreateBuilding).Methods(http.MethodPost, http.MethodOptions)
	complex.HandleFunc("/images/{id}", complexHandler.UpdateComplexPhoto).Methods(http.MethodPost, http.MethodOptions)

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
		logger.Info(fmt.Sprintf("Start server on %s\n", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	sig := <-signalCh
	logger.Info(fmt.Sprintf("Received signal: %v\n", sig))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Server shutdown failed: %s\n", err))
	}
}

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "pong")
}
