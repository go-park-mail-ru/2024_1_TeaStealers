package repo

import (
	"2024_1_TeaStealers/internal/models"
	"database/sql"
	"errors"
	"github.com/satori/uuid"
)

// UserRepo represents a repository for user.
type UserRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of UserRepo.
func NewRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserById(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, firstname, secondname, datebirthday, phone, email, photo FROM users WHERE id=$1`
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

func (r *UserRepo) UpdateUserPhoto(id uuid.UUID, fileName string) (string, error) {
	query := `UPDATE users SET photo = $1 WHERE id = $2`
	if _, err := r.db.Query(query, fileName, id); err != nil {
		return "", err
	}
	return fileName, nil
}

func (r *UserRepo) DeleteUserPhoto(id uuid.UUID) error {
	query := `UPDATE users SET photo = '' WHERE id = $1`
	if _, err := r.db.Query(query, id); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateUserInfo(id uuid.UUID, data *models.UserUpdateData) (*models.User, error) {
	query := `UPDATE users SET firstname = $1, secondname = $2, datebirthday = $3, phone = $4, email = $5 WHERE id = $6`
	if _, err := r.db.Query(query, data.FirstName, data.SecondName, data.DateBirthday, data.Phone, data.Email, id); err != nil {
		return nil, err
	}
	var firstname, secondname, photo sql.NullString
	var dateBirthday sql.NullTime
	user := &models.User{}
	querySelect := `SELECT id, firstname, secondname, datebirthday, phone, email, photo FROM users WHERE id = $1`
	res := r.db.QueryRow(querySelect, id)
	if err := res.Scan(&user.ID, &firstname, &secondname, &dateBirthday, &user.Phone, &user.Email, &photo); err != nil {
		return nil, err
	}
	user.FirstName = firstname.String
	user.SecondName = secondname.String
	user.DateBirthday = dateBirthday.Time
	user.Phone = photo.String
	return user, nil
}

func (r *UserRepo) UpdateUserPassword(id uuid.UUID, newPasswordHash string) (int, error) {
	query := `UPDATE users SET passwordhash=$1, levelupdate = levelupdate+1 WHERE id = $2`
	if _, err := r.db.Query(query, newPasswordHash, id); err != nil {
		return 0, err
	}
	querySelect := `SELECT levelupdate FROM users WHERE id = $1`
	level := 0
	res := r.db.QueryRow(querySelect, id)
	if err := res.Scan(&level); err != nil {
		return 0, err
	}
	return level, nil
}

func (r *UserRepo) CheckUserPassword(id uuid.UUID, passwordHash string) error {
	passwordHashCur := ""
	querySelect := `SELECT passwordhash FROM users WHERE id = $1`
	res := r.db.QueryRow(querySelect, id)
	if err := res.Scan(&passwordHashCur); err != nil {
		return err
	}
	if passwordHashCur != passwordHash {
		return errors.New("passwords don't match")
	}
	return nil
}

// ToDo: отдавать нового юзера, в юзкейсе генерировать новый жвт с новым уровнем, поменять функции
