package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"errors"
)

type AuthRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) CreateUser(ctx context.Context, user *models.User) error {
	insert := `INSERT INTO users (id, login, phone, passwordHash) VALUES ($1, $2, $3, $4)`

	if _, err := r.db.ExecContext(ctx, insert, user.ID, user.Login, user.Phone, user.PasswordHash); err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	query := `SELECT * FROM users WHERE login = $1`

	res := r.db.QueryRowContext(ctx, query, login)

	user := &models.User{
		Login: login,
	}
	if err := res.Scan(&user.ID, &user.Login, &user.Phone, &user.PasswordHash); err != nil {
		return nil, err
	}

	return user, nil
}

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
