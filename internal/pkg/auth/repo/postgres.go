package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

// AuthRepo represents a repository for authentication.
type AuthRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of AuthRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *AuthRepo {
	return &AuthRepo{db: db, logger: logger}
}

func (r *AuthRepo) BeginTx(ctx context.Context) (models.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, auth.BeginTxMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, auth.BeginTxMethod)
	return tx, nil
}

// CreateUser creates a new user in the database.
func (r *AuthRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	insert := `INSERT INTO user_data (email, phone, password_hash, first_name, surname, photo) VALUES ($1, $2, $3, '', '', '') RETURNING id`
	var lastInsertID int64
	if err := r.db.QueryRowContext(ctx, insert, user.Email, user.Phone, user.PasswordHash).Scan(&lastInsertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, auth.CreateUserMethod, err)
		return nil, err
	}

	query := `SELECT email, phone, password_hash, level_update FROM user_data WHERE id = $1`

	res := r.db.QueryRow(query, lastInsertID)

	newUser := &models.User{ID: lastInsertID}
	if err := res.Scan(&newUser.Email, &newUser.Phone, &newUser.PasswordHash, &newUser.LevelUpdate); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, auth.CreateUserMethod, err)
		return nil,
			err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, auth.CreateUserMethod)
	return newUser, nil
}

// GetUserByLogin retrieves a user from the database by their login.
func (r *AuthRepo) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	query := `SELECT id, email, phone, password_hash, level_update FROM user_data WHERE email = $1 OR phone = $1`

	res := r.db.QueryRowContext(ctx, query, login)

	user := &models.User{}
	if err := res.Scan(&user.ID, &user.Email, &user.Phone, &user.PasswordHash, &user.LevelUpdate); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserByLoginMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserByLoginMethod)
	return user, nil
}

// CheckUser checks if the user with the given login and password hash exists in the database.
func (r *AuthRepo) CheckUser(ctx context.Context, login string, passwordHash string) (*models.User, error) {
	user, err := r.GetUserByLogin(ctx, login)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.CheckUserMethod, err)
		return nil, err
	}

	if user.PasswordHash != passwordHash {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.CheckUserMethod, errors.New("Password hash not equal"))
		return nil, errors.New("wrong password")
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.CheckUserMethod)
	return user, nil
}

func (r *AuthRepo) GetUserLevelById(ctx context.Context, id int64) (int, error) {
	query := `SELECT level_update FROM user_data WHERE id = $1`

	res := r.db.QueryRow(query, id)

	level := 0
	if err := res.Scan(&level); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserLevelByIdMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, auth.GetUserLevelByIdMethod)
	return level, nil
}
