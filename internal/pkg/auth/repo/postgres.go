package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/satori/uuid"
)

// AuthRepo represents a repository for authentication.
type AuthRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of AuthRepo.
func NewRepository(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

// CreateUser creates a new user in the database.
func (r *AuthRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	insert := `INSERT INTO users (id, email, phone, passwordhash) VALUES ($1, $2, $3, $4)`
	if _, err := r.db.ExecContext(ctx, insert, user.ID, user.Email, user.Phone, user.PasswordHash); err != nil {
		return nil, err
	}
	query := `SELECT id, email, phone, passwordhash, levelupdate FROM users WHERE id = $1`

	res := r.db.QueryRow(query, user.ID)

	newUser := &models.User{}
	if err := res.Scan(&user.ID, &user.Email, &user.Phone, &user.PasswordHash, &user.LevelUpdate); err != nil {
		return nil, err
	}

	return newUser, nil
}

// GetUserByLogin retrieves a user from the database by their login.
func (r *AuthRepo) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	query := `SELECT id, email, phone, passwordhash, levelupdate FROM users WHERE email = $1 OR phone = $1`

	res := r.db.QueryRowContext(ctx, query, login)

	user := &models.User{}
	if err := res.Scan(&user.ID, &user.Email, &user.Phone, &user.PasswordHash, &user.LevelUpdate); err != nil {
		return nil, err
	}

	return user, nil
}

// CheckUser checks if the user with the given login and password hash exists in the database.
func (r *AuthRepo) CheckUser(ctx context.Context, login string, passwordHash string) (*models.User, error) {
	user, err := r.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash != passwordHash {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (r *AuthRepo) GetUserLevelById(id uuid.UUID) (int, error) {
	query := `SELECT levelupdate FROM users WHERE id = $1`

	res := r.db.QueryRow(query, id)

	level := 0
	if err := res.Scan(&level); err != nil {
		return 0, err
	}
	fmt.Println(level)
	return level, nil
}
