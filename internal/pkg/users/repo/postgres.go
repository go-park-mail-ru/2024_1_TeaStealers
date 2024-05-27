package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"fmt"
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
	query := `SELECT id, phone, email FROM user_data WHERE id=$1`
	queryProfile := `SELECT first_name, surname, photo, birthdate FROM user_profile_data WHERE user_id=$1`

	res := r.db.QueryRow(query, id)
	resProfile := r.db.QueryRow(queryProfile, id)

	var firstname, secondname, photo sql.NullString
	var dateBirthday sql.NullTime
	if err := res.Scan(&user.ID, &user.Phone, &user.Email); err != nil {
		return nil, err
	}
	if err := resProfile.Scan(&firstname, &secondname, &photo, &dateBirthday); err != nil {
		fmt.Println("У пользователя нет  информации профиля")
		fmt.Println(err.Error())
	}
	user.FirstName = firstname.String
	user.SecondName = secondname.String
	user.Photo = photo.String
	user.DateBirthday = dateBirthday.Time

	return user, nil
}

func (r *UserRepo) UpdateUserPhoto(ctx context.Context, id int64, fileName string) (string, error) {
	query := `INSERT INTO user_profile_data (user_id, first_name, surname, birthdate, photo)  VALUES ($1, ' ', ' ', '0001-01-01T00:00:00Z', $2) ON CONFLICT (user_id) DO UPDATE SET photo = EXCLUDED.photo;`
	if _, err := r.db.Exec(query, id, fileName); err != nil {
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
	if data.FirstName == "" {
		data.FirstName = " "
	}
	if data.SecondName == "" {
		data.SecondName = " "
	}
	query := `INSERT INTO user_profile_data (user_id, first_name, surname, birthdate, photo)  VALUES ($1, $2, $3, '0001-01-01T00:00:00Z', 'avatar/defaultAvatar.png') ON CONFLICT (user_id) DO UPDATE SET first_name = excluded.first_name, surname = excluded.surname;`

	if _, err := r.db.Exec(query, id, data.FirstName, data.SecondName); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	updateQuery := `
        UPDATE user_data
        SET phone = $1, email = $2
        WHERE id = $3;
    `
	_, err := r.db.ExecContext(ctx, updateQuery, data.Phone, data.Email, id)
	if err != nil {
		fmt.Println("Error executing update query:", err)
		return nil, err
	}

	return r.GetUserById(ctx, id)
}
