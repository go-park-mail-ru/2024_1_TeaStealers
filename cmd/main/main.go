package main

import (
	advertsH "2024_1_TeaStealers/internal/pkg/adverts/delivery/http"
	advertsR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertsUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	authH "2024_1_TeaStealers/internal/pkg/auth/delivery/http"
	companyH "2024_1_TeaStealers/internal/pkg/companies/delivery"
	companyR "2024_1_TeaStealers/internal/pkg/companies/repo"
	companyUc "2024_1_TeaStealers/internal/pkg/companies/usecase"
	complexH "2024_1_TeaStealers/internal/pkg/complexes/delivery"
	complexR "2024_1_TeaStealers/internal/pkg/complexes/repo"
	complexUc "2024_1_TeaStealers/internal/pkg/complexes/usecase"
	"2024_1_TeaStealers/internal/pkg/config"
	imageH "2024_1_TeaStealers/internal/pkg/images/delivery/http"
	imageR "2024_1_TeaStealers/internal/pkg/images/repo"
	imageUc "2024_1_TeaStealers/internal/pkg/images/usecase"
	metricsMw "2024_1_TeaStealers/internal/pkg/metrics/middleware"
	"2024_1_TeaStealers/internal/pkg/middleware"
	statsH "2024_1_TeaStealers/internal/pkg/questionnaire/delivery"
	statsR "2024_1_TeaStealers/internal/pkg/questionnaire/repo"
	statsUc "2024_1_TeaStealers/internal/pkg/questionnaire/usecase"
	http2 "2024_1_TeaStealers/internal/pkg/users/delivery/http"
	userR "2024_1_TeaStealers/internal/pkg/users/repo"
	userUc "2024_1_TeaStealers/internal/pkg/users/usecase"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "2024_1_TeaStealers/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	cfg := config.MustLoad()
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

	http.Handle("/metrics", promhttp.Handler())
	metricsMd := metricsMw.Create()
	// metricsMd.ServerMetricsMiddleware()
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(middleware.CORSMiddleware)
	r.Handle("/ping", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(pingPongHandler), 0, 0, "")).Methods(http.MethodGet)
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	grcpConnAuth, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.AuthContainerIP, cfg.GRPC.AuthPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnAuth.Close()

	grcpConnUsers, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.UsersContainerIP, cfg.GRPC.UserPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("cant connect to grpc")
		return
	}
	defer grcpConnUsers.Close()

	grcpConnAdverts, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.AdvertContainerIP, cfg.GRPC.AdvertPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("cant connect to grpc")
		return
	}
	defer grcpConnAdverts.Close()

	authHandler := authH.NewClientAuthHandler(grcpConnAuth, logger)
	jwtMd := middleware.NewAuthMiddleware(grcpConnAuth, logger)
	csrfMd := middleware.NewCsrfMiddleware()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.Handle("/signup", metricsMd.ServerMetricsMiddleware(csrfMd.SetCSRFToken(http.HandlerFunc(authHandler.SignUp)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/login", metricsMd.ServerMetricsMiddleware(csrfMd.SetCSRFToken(http.HandlerFunc(authHandler.Login)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/logout", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.Logout)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)
	auth.Handle("/check_auth", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.CheckAuth)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)

	statRepo := statsR.NewRepository(db, logger)
	statUsecase := statsUc.NewQuestionnaireUsecase(statRepo, logger)
	statHandler := statsH.NewQuestionnaireHandler(statUsecase, logger)
	stat := r.PathPrefix("/stat").Subrouter()
	stat.Handle("/answer", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.UploadAnswer)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	stat.Handle("/theme", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.GetAnswerStatistics)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)
	stat.Handle("/{theme}/questions", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.GetQuestionsByTheme)), 0, 3, "theme")).Methods(http.MethodGet, http.MethodOptions)

	advertRepo := advertsR.NewRepository(db, logger)
	advertUsecase := advertsUc.NewAdvertUsecase(advertRepo, logger)
	advertHandler := advertsH.NewAdvertsClientHandler(grcpConnAdverts, advertUsecase, logger)

	imageRepo := imageR.NewRepository(db, logger)
	imageUsecase := imageUc.NewImageUsecase(imageRepo, logger)
	imageHandler := imageH.NewImageHandler(imageUsecase, logger)

	advert := r.PathPrefix("/adverts").Subrouter()
	advert.Handle("/{id}", metricsMd.ServerMetricsMiddleware(jwtMd.StatMiddleware(http.HandlerFunc(advertHandler.GetAdvertById)), 0, 3, "getAdvert")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.UpdateAdvertById)), 0, 3, "updateAdvert")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.DeleteAdvertById)), 0, 3, "deleteAdvert")).Methods(http.MethodDelete, http.MethodOptions)
	advert.Handle("/{id}/like", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.LikeAdvert)), 0, 3, "id")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}/dislike", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.DislikeAdvert)), 0, 3, "id")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/houses/", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.CreateHouseAdvert)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/building/", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(advertHandler.GetExistBuildingByAddress), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/flats/", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.CreateFlatAdvert)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/squarelist/", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(advertHandler.GetSquareAdvertsList), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/rectanglelist/", metricsMd.ServerMetricsMiddleware(jwtMd.StatMiddleware(http.HandlerFunc(advertHandler.GetRectangeAdvertsList)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/image/", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(imageHandler.UploadImage)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}/image", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(imageHandler.GetAdvertImages), 0, 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}/image", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(imageHandler.DeleteImage)), 0, 3, "deleteId")).Methods(http.MethodDelete, http.MethodOptions)

	userRepo := userR.NewRepository(db)
	userUsecase := userUc.NewUserUsecase(userRepo)
	userHandler := http2.NewClientUserHandler(grcpConnUsers)
	userHandlerPhoto := http2.NewUserHandlerPhoto(userUsecase)

	user := r.PathPrefix("/users").Subrouter()
	user.Handle("/me", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.GetCurUser)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/avatar", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandlerPhoto.UpdateUserPhoto)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/avatar", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandlerPhoto.DeleteUserPhoto)), 0, 0, "")).Methods(http.MethodDelete, http.MethodOptions)
	user.Handle("/info", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.UpdateUserInfo)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/password", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.UpdateUserPassword)), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/myadverts", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetUserAdverts)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/likedadverts", metricsMd.ServerMetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetLikedUserAdverts)), 0, 0, "")).Methods(http.MethodGet, http.MethodOptions)

	companyRepo := companyR.NewRepository(db, logger)
	companyUsecase := companyUc.NewCompanyUsecase(companyRepo, logger)
	companyHandler := companyH.NewCompanyHandler(companyUsecase, logger)

	company := r.PathPrefix("/companies").Subrouter()
	company.Handle("/", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(companyHandler.CreateCompany), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	company.Handle("/{id}", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(companyHandler.GetCompanyById), 0, 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	company.Handle("/images/{id}", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(companyHandler.UpdateCompanyPhoto), 0, 4, "id")).Methods(http.MethodPost, http.MethodOptions)

	complexRepo := complexR.NewRepository(db, logger)
	complexUsecase := complexUc.NewComplexUsecase(complexRepo, logger)
	complexHandler := complexH.NewComplexHandler(complexUsecase, logger)

	complexRoute := r.PathPrefix("/complexes").Subrouter()
	complexRoute.Handle("/", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(complexHandler.CreateComplex), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	complexRoute.Handle("/{id}", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(complexHandler.GetComplexById), 0, 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	complexRoute.Handle("/{id}/rectanglelist/", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(advertHandler.GetComplexAdverts), 0, 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	complexRoute.Handle("/houses", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(complexHandler.CreateHouseAdvert), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	complexRoute.Handle("/flats", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(complexHandler.CreateFlatAdvert), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	complexRoute.Handle("/buildings", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(complexHandler.CreateBuilding), 0, 0, "")).Methods(http.MethodPost, http.MethodOptions)
	complexRoute.Handle("/images/{id}", metricsMd.ServerMetricsMiddleware(http.HandlerFunc(complexHandler.UpdateComplexPhoto), 0, 4, "id")).Methods(http.MethodPost, http.MethodOptions)
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
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
