package main

import (
	advertsH "2024_1_TeaStealers/internal/pkg/adverts/delivery/http"
	advertsR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertsUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	authH "2024_1_TeaStealers/internal/pkg/auth/delivery/http"
	complexH "2024_1_TeaStealers/internal/pkg/complexes/delivery/http"
	"2024_1_TeaStealers/internal/pkg/config"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	imageH "2024_1_TeaStealers/internal/pkg/images/delivery/http"
	imageR "2024_1_TeaStealers/internal/pkg/images/repo"
	imageUc "2024_1_TeaStealers/internal/pkg/images/usecase"
	metricsMw "2024_1_TeaStealers/internal/pkg/metrics/middleware"
	"2024_1_TeaStealers/internal/pkg/middleware"
	statsH "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/http"
	statsR "2024_1_TeaStealers/internal/pkg/questionnaire/repo"
	statsUc "2024_1_TeaStealers/internal/pkg/questionnaire/usecase"
	userH "2024_1_TeaStealers/internal/pkg/users/delivery/http"
	userR "2024_1_TeaStealers/internal/pkg/users/repo"
	userUc "2024_1_TeaStealers/internal/pkg/users/usecase"
	"context"
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
	maxConns := int32(10)
	_ = godotenv.Load()
	logger := zap.Must(zap.NewDevelopment())
	dbPool.InitDatabasePool(fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.Database.DBUser,
		cfg.Database.DBPass,
		cfg.Database.DBHost,
		cfg.Database.DBPort,
		cfg.Database.DBName), maxConns)
	pool := dbPool.GetDBPool()
	if err := pool.Ping(context.Background()); err != nil {
		log.Println("fail ping postgres")
		err = fmt.Errorf("error happened in db.Ping: %w", err)
		log.Println(err)
	}
	metricmW := metricsMw.Create()
	metricmW.RegisterMetrics()
	go metricmW.UpdatePSS()

	// http.Handle("/metrics", promhttp.Handler())

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	r.Use(middleware.CORSMiddleware, middleware.AccessLogMiddleware)
	r.HandleFunc("/ping", pingPongHandler).Methods(http.MethodGet)
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/metrics").Handler(promhttp.Handler())
	grcpConnAuth, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.AuthContainerIP, cfg.GRPC.AuthPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("cant connect to grpc")
	}
	defer grcpConnAuth.Close()

	grcpConnQuestion, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.QuestionContainerIP, cfg.GRPC.QuestionPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	defer grcpConnQuestion.Close()
	if err != nil {
		log.Println("cant connect to grpc")
	}

	grcpConnComplex, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.ComplexContainerIP, cfg.GRPC.ComplexPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("cant connect to grpc")
	}
	defer grcpConnComplex.Close()

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
	auth.Handle("/signup", metricmW.MetricsMiddleware(csrfMd.SetCSRFToken(http.HandlerFunc(authHandler.SignUp)), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/login", metricmW.MetricsMiddleware(csrfMd.SetCSRFToken(http.HandlerFunc(authHandler.Login)), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/logout", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.Logout), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	auth.Handle("/check_auth", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.CheckAuth), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)

	statRepo := statsR.NewRepository(logger)
	statUsecase := statsUc.NewQuestionnaireUsecase(statRepo, logger)
	statHandler := statsH.NewQuestionnaireClientHandler(grcpConnQuestion, statUsecase, logger)
	stat := r.PathPrefix("/stat").Subrouter()
	stat.Handle("/answer", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.UploadAnswer), metricmW), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	stat.Handle("/theme", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.GetAnswerStatistics), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	stat.Handle("/{theme}/questions", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.GetQuestionsByTheme), metricmW), 3, "theme")).Methods(http.MethodGet, http.MethodOptions)

	advertRepo := advertsR.NewRepository(logger, metricmW)
	advertUsecase := advertsUc.NewAdvertUsecase(advertRepo, logger)
	advertHandler := advertsH.NewAdvertsClientHandler(grcpConnAdverts, grcpConnComplex, advertUsecase, logger)

	imageRepo := imageR.NewRepository(logger)
	imageUsecase := imageUc.NewImageUsecase(imageRepo, logger)
	imageHandler := imageH.NewImageHandler(imageUsecase, logger)

	advert := r.PathPrefix("/adverts").Subrouter()
	advert.Handle("/{id}", metricmW.MetricsMiddleware(jwtMd.StatMiddleware(http.HandlerFunc(advertHandler.GetAdvertById)), 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.UpdateAdvertById), metricmW), 3, "id")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.DeleteAdvertById), metricmW), 3, "id")).Methods(http.MethodDelete, http.MethodOptions)
	advert.Handle("/{id}/like", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.LikeAdvert), metricmW)).Methods(http.MethodPost, http.MethodOptions)

	advert.Handle("/{id}/dislike", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.DislikeAdvert), metricmW), 3, "id")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/houses/", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.CreateHouseAdvert), metricmW), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/building/", metricmW.MetricsMiddleware(http.HandlerFunc(advertHandler.GetExistBuildingByAddress), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/flats/", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.CreateFlatAdvert), metricmW), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/squarelist/", metricmW.MetricsMiddleware(http.HandlerFunc(advertHandler.GetSquareAdvertsList), 0, "")).Methods(http.MethodGet, http.MethodOptions)

	advert.Handle("/rectanglelist/", metricmW.MetricsMiddleware(jwtMd.StatMiddleware(http.HandlerFunc(advertHandler.GetRectangleAdvertsList)), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/image/", metricmW.MetricsMiddleware(http.HandlerFunc(imageHandler.UploadImage), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}/image", metricmW.MetricsMiddleware(http.HandlerFunc(imageHandler.GetAdvertImages), 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}/image", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(imageHandler.DeleteImage), metricmW), 3, "id")).Methods(http.MethodDelete, http.MethodOptions)
	advert.Handle("/{id}/donate", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.UpdatePriority), metricmW), 3, "id")).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}/rating", metricmW.MetricsMiddleware(http.HandlerFunc(advertHandler.GetPriority), 0, "")).Methods(http.MethodGet, http.MethodOptions)

	userRepo := userR.NewRepository(metricmW)
	userUsecase := userUc.NewUserUsecase(userRepo)
	userHandler := userH.NewClientUserHandler(grcpConnUsers)
	userHandlerPhoto := userH.NewUserHandlerPhoto(userUsecase)

	user := r.PathPrefix("/users").Subrouter()
	user.Handle("/me", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.GetCurUser), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/avatar", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandlerPhoto.UpdateUserPhoto), metricmW), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/avatar", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandlerPhoto.DeleteUserPhoto), metricmW), 0, "")).Methods(http.MethodDelete, http.MethodOptions)
	user.Handle("/info", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.UpdateUserInfo), metricmW), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/password", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.UpdateUserPassword), metricmW), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/savedadverts", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetUserAdverts), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/myadverts", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetUserAdverts), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/likedadverts", metricmW.MetricsMiddleware(jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetLikedUserAdverts), metricmW), 0, "")).Methods(http.MethodGet, http.MethodOptions)

	complexHandler := complexH.NewClientComplexHandler(grcpConnComplex, logger)

	complexRoute := r.PathPrefix("/complexes").Subrouter()
	complexRoute.Handle("/", metricmW.MetricsMiddleware(http.HandlerFunc(complexHandler.CreateComplex), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	complexRoute.Handle("/{id}", metricmW.MetricsMiddleware(http.HandlerFunc(complexHandler.GetComplexById), 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	complexRoute.Handle("/{id}/rectanglelist/", metricmW.MetricsMiddleware(http.HandlerFunc(advertHandler.GetComplexAdverts), 3, "id")).Methods(http.MethodGet, http.MethodOptions)
	complexRoute.Handle("/images/{id}", metricmW.MetricsMiddleware(http.HandlerFunc(complexHandler.UpdateComplexPhoto), 4, "id")).Methods(http.MethodPost, http.MethodOptions)

	company := r.PathPrefix("/companies").Subrouter()
	company.Handle("/", metricmW.MetricsMiddleware(http.HandlerFunc(complexHandler.CreateCompany), 0, "")).Methods(http.MethodPost, http.MethodOptions)
	company.Handle("/{id}", metricmW.MetricsMiddleware(http.HandlerFunc(complexHandler.GetCompanyById), 0, "")).Methods(http.MethodGet, http.MethodOptions)
	company.Handle("/images/{id}", metricmW.MetricsMiddleware(http.HandlerFunc(complexHandler.UpdateCompanyPhoto), 0, "")).Methods(http.MethodPost, http.MethodOptions)

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
