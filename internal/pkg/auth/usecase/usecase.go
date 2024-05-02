package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"errors"
	"time"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// AuthUsecase represents the usecase for authentication.
type AuthUsecase struct {
	repo   auth.AuthRepo
	logger *zap.Logger
}

// NewAuthUsecase creates a new instance of AuthUsecase.
func NewAuthUsecase(repo auth.AuthRepo, logger *zap.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, logger: logger}
}

// SignUp handles the user registration process.
func (u *AuthUsecase) SignUp(ctx context.Context, data *models.UserSignUpData) (*models.User, string, time.Time, error) {
	newUser := &models.User{
		ID:           uuid.NewV4(),
		Email:        data.Email,
		Phone:        data.Phone,
		PasswordHash: utils.GenerateHashString(data.Password),
		LevelUpdate:  1,
	}

	userResponse, err := u.repo.CreateUser(ctx, newUser)
	if err != nil {
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.SignUpMethod, err)
		return nil, "", time.Now(), err
	}

	token, exp, err := jwt.GenerateToken(newUser)
	if err != nil {
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.SignUpMethod, err)
		return nil, "", time.Now(), err
	}

	// utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.SignUpMethod)
	return userResponse, token, exp, nil
}

// Login handles the user login process.
func (u *AuthUsecase) Login(ctx context.Context, data *models.UserLoginData) (*models.User, string, time.Time, error) {
	user, err := u.repo.CheckUser(ctx, data.Login, utils.GenerateHashString(data.Password))
	if err != nil {
		u.logger.Error(err.Error())
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.LoginMethod, err)
		return nil, "", time.Now(), err
	}

	token, exp, err := jwt.GenerateToken(user)
	if err != nil {
		u.logger.Error(err.Error())
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.LoginMethod, err)
		return nil, "", time.Now(), err
	}

	u.logger.Info("success login usecase")
	// utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.LoginMethod)
	return user, token, exp, nil
}

// CheckAuth checking autorizing
func (u *AuthUsecase) CheckAuth(ctx context.Context, id uuid.UUID, jwtLevel int) error {
	level, err := u.repo.GetUserLevelById(ctx, id)
	if err != nil {
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserLevelByIdMethod, err)
		return err
	}
	if jwtLevel != level {
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserLevelByIdMethod, errors.New("jwt levels not equal"))
		return errors.New("levels don't match")
	}

	// utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserLevelByIdMethod)
	return nil
}
