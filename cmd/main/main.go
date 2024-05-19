package main

import (
	advertsH "2024_1_TeaStealers/internal/pkg/adverts/delivery/http"
	advertsR "2024_1_TeaStealers/internal/pkg/adverts/repo"
	advertsUc "2024_1_TeaStealers/internal/pkg/adverts/usecase"
	authH "2024_1_TeaStealers/internal/pkg/auth/delivery/http"
	complexH "2024_1_TeaStealers/internal/pkg/complexes/delivery/http"
	complexR "2024_1_TeaStealers/internal/pkg/complexes/repo"
	complexUc "2024_1_TeaStealers/internal/pkg/complexes/usecase"
	"2024_1_TeaStealers/internal/pkg/config"
	imageH "2024_1_TeaStealers/internal/pkg/images/delivery/http"
	imageR "2024_1_TeaStealers/internal/pkg/images/repo"
	imageUc "2024_1_TeaStealers/internal/pkg/images/usecase"
	"2024_1_TeaStealers/internal/pkg/middleware"
	statsH "2024_1_TeaStealers/internal/pkg/questionnaire/delivery/http"
	statsR "2024_1_TeaStealers/internal/pkg/questionnaire/repo"
	statsUc "2024_1_TeaStealers/internal/pkg/questionnaire/usecase"
	userH "2024_1_TeaStealers/internal/pkg/users/delivery/http"
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

	//http.Handle("/metrics", promhttp.Handler())

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
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnAuth.Close()

	grcpConnQuestion, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.QuestionContainerIP, cfg.GRPC.QuestionPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnQuestion.Close()

	grcpConnComplex, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GRPC.ComplexContainerIP, cfg.GRPC.ComplexPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
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
	auth.Handle("/signup", csrfMd.SetCSRFToken(http.HandlerFunc(authHandler.SignUp))).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/login", csrfMd.SetCSRFToken(http.HandlerFunc(authHandler.Login))).Methods(http.MethodPost, http.MethodOptions)
	auth.Handle("/logout", jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.Logout))).Methods(http.MethodGet, http.MethodOptions)
	auth.Handle("/check_auth", jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.CheckAuth))).Methods(http.MethodGet, http.MethodOptions)

	statRepo := statsR.NewRepository(db, logger)
	statUsecase := statsUc.NewQuestionnaireUsecase(statRepo, logger)
	statHandler := statsH.NewQuestionnaireClientHandler(grcpConnQuestion, statUsecase, logger)
	stat := r.PathPrefix("/stat").Subrouter()
	stat.Handle("/answer", jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.UploadAnswer))).Methods(http.MethodPost, http.MethodOptions)
	stat.Handle("/theme", jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.GetAnswerStatistics))).Methods(http.MethodGet, http.MethodOptions)
	stat.Handle("/{theme}/questions", jwtMd.JwtMiddleware(http.HandlerFunc(statHandler.GetQuestionsByTheme))).Methods(http.MethodGet, http.MethodOptions)

	advertRepo := advertsR.NewRepository(db, logger)
	advertUsecase := advertsUc.NewAdvertUsecase(advertRepo, logger)
	advertHandler := advertsH.NewAdvertsClientHandler(grcpConnAdverts, grcpConnComplex, advertUsecase, logger)

	imageRepo := imageR.NewRepository(db, logger)
	imageUsecase := imageUc.NewImageUsecase(imageRepo, logger)
	imageHandler := imageH.NewImageHandler(imageUsecase, logger)

	advert := r.PathPrefix("/adverts").Subrouter()
	advert.Handle("/{id}", jwtMd.StatMiddleware(http.HandlerFunc(advertHandler.GetAdvertById))).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.UpdateAdvertById))).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.DeleteAdvertById))).Methods(http.MethodDelete, http.MethodOptions)
	advert.Handle("/{id}/like", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.LikeAdvert))).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/{id}/dislike", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.DislikeAdvert))).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/houses/", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.CreateHouseAdvert))).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/building/", advertHandler.GetExistBuildingByAddress).Methods(http.MethodPost, http.MethodOptions)
	advert.Handle("/flats/", http.HandlerFunc(advertHandler.CreateFlatAdvert)).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/squarelist/", advertHandler.GetSquareAdvertsList).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/rectanglelist/", jwtMd.StatMiddleware(http.HandlerFunc(advertHandler.GetRectangleAdvertsList))).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/image/", http.HandlerFunc(imageHandler.UploadImage)).Methods(http.MethodPost, http.MethodOptions)
	advert.HandleFunc("/{id}/image", imageHandler.GetAdvertImages).Methods(http.MethodGet, http.MethodOptions)
	advert.Handle("/{id}/image", jwtMd.JwtMiddleware(http.HandlerFunc(imageHandler.DeleteImage))).Methods(http.MethodDelete, http.MethodOptions)

	userRepo := userR.NewRepository(db)
	userUsecase := userUc.NewUserUsecase(userRepo)
	userHandler := userH.NewClientUserHandler(grcpConnUsers)
	userHandlerPhoto := userH.NewUserHandlerPhoto(userUsecase)

	user := r.PathPrefix("/users").Subrouter()
	user.Handle("/me", jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.GetCurUser))).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/avatar", jwtMd.JwtMiddleware(http.HandlerFunc(userHandlerPhoto.UpdateUserPhoto))).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/avatar", jwtMd.JwtMiddleware(http.HandlerFunc(userHandlerPhoto.DeleteUserPhoto))).Methods(http.MethodDelete, http.MethodOptions)
	user.Handle("/info", jwtMd.JwtMiddleware(http.HandlerFunc(userHandler.UpdateUserInfo))).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/password", jwtMd.JwtMiddleware(http.HandlerFunc(authHandler.UpdateUserPassword))).Methods(http.MethodPost, http.MethodOptions)
	user.Handle("/myadverts", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetUserAdverts))).Methods(http.MethodGet, http.MethodOptions)
	user.Handle("/likedadverts", jwtMd.JwtMiddleware(http.HandlerFunc(advertHandler.GetLikedUserAdverts))).Methods(http.MethodGet, http.MethodOptions)

	complexRepo := complexR.NewRepository(db, logger)
	complexUsecase := complexUc.NewComplexUsecase(complexRepo, logger)
	complexHandler := complexH.NewClientComplexHandler(grcpConnComplex, complexUsecase, logger)

	complexRoute := r.PathPrefix("/complexes").Subrouter()
	complexRoute.HandleFunc("/", complexHandler.CreateComplex).Methods(http.MethodPost, http.MethodOptions)
	complexRoute.HandleFunc("/{id}", complexHandler.GetComplexById).Methods(http.MethodGet, http.MethodOptions)
	complexRoute.HandleFunc("/{id}/rectanglelist/", advertHandler.GetComplexAdverts).Methods(http.MethodGet, http.MethodOptions)
	complexRoute.HandleFunc("/images/{id}", complexHandler.UpdateComplexPhoto).Methods(http.MethodPost, http.MethodOptions)

	company := r.PathPrefix("/companies").Subrouter()
	company.HandleFunc("/", complexHandler.CreateCompany).Methods(http.MethodPost, http.MethodOptions)
	company.HandleFunc("/{id}", complexHandler.GetCompanyById).Methods(http.MethodGet, http.MethodOptions)
	company.HandleFunc("/images/{id}", complexHandler.UpdateCompanyPhoto).Methods(http.MethodPost, http.MethodOptions)

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
