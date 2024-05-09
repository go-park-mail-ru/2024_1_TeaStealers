package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
)

// UserRepo represents a repository for user.
type UserRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of UserRepo.
func NewRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, first_name, surname, birthdate, phone, email, photo FROM user_data WHERE id=$1`
	res := r.db.QueryRow(query, id)
	var firstname, secondname, photo sql.NullString
	var dateBirthday sql.NullTime
	if err := res.Scan(&user.ID, &firstname, &secondname, &dateBirthday, &user.Phone, &user.Email, &photo); err != nil {
		return nil, err
	}
	user.FirstName = firstname.String
	user.SecondName = secondname.String
	user.Photo = photo.String
	user.DateBirthday = dateBirthday.Time

	return user, nil
}

func (r *UserRepo) UpdateUserPhoto(ctx context.Context, id int64, fileName string) (string, error) {
	query := `UPDATE user_data SET photo = $1 WHERE id = $2`
	if _, err := r.db.Query(query, fileName, id); err != nil {
		return "", err
	}
	return fileName, nil
}

func (r *UserRepo) DeleteUserPhoto(ctx context.Context, id int64) error {
	query := `UPDATE user_data SET photo = '' WHERE id = $1`
	if _, err := r.db.Query(query, id); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateUserInfo(ctx context.Context, id int64, data *models.UserUpdateData) (*models.User, error) {
	query := `UPDATE user_data SET first_name = $1, surname = $2, phone = $3, email = $4 WHERE id = $5`

	if _, err := r.db.Exec(query, data.FirstName, data.SecondName, data.Phone, data.Email, id); err != nil {
		return nil, err
	}
	user := &models.User{}
	querySelect := `SELECT id, first_name, surname, phone, email FROM user_data WHERE id = $1`
	res := r.db.QueryRow(querySelect, id)
	if err := res.Scan(&user.ID, &user.FirstName, &user.SecondName, &user.Phone, &user.Email); err != nil {
		return nil, err
	}

	return user, nil
}
