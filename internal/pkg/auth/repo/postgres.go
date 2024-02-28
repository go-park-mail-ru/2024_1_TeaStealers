package repo

import (
	"2024_1_TeaStealers/internal/models"
	"database/sql"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Create(user *models.User) error {
	insert := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := r.db.Exec(insert, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	return err
}

func (r *RepositoryImpl) UpdateInfo(user *models.User) error {
	update := `UPDATE users SET (email, password, FirstName, LastName) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(update, user.Email, user.PasswordHash, user.FirstName, user.LastName)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r *RepositoryImpl) Delete(user *models.User) error {

}

func (r *RepositoryImpl) GetByJWT(user *models.User) error {

}

func (r *RepositoryImpl) GetByEmail(user *models.User) error {

}
