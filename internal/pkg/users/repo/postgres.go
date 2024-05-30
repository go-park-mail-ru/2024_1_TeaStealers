package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	"2024_1_TeaStealers/internal/pkg/metrics"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// UserRepo represents a repository for user.
type UserRepo struct {
	db       *pgxpool.Pool
	metricsC metrics.MetricsHTTP
}

// NewRepository creates a new instance of UserRepo.
func NewRepository(metrics metrics.MetricsHTTP) *UserRepo {
	return &UserRepo{db: dbPool.GetDBPool(), metricsC: metrics}
}

func (r *UserRepo) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, phone, email FROM user_data WHERE id=$1`
	queryProfile := `SELECT first_name, surname, photo, birthdate FROM user_profile_data WHERE user_id=$1`

	var dur time.Duration
	start := time.Now()
	res := r.db.QueryRow(ctx, query, id)
	resProfile := r.db.QueryRow(ctx, queryProfile, id)
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("GetUserById", "select user_profile_data", dur)

	var firstname, secondname, photo sql.NullString
	var dateBirthday sql.NullTime

	start = time.Now()
	if err := res.Scan(&user.ID, &user.Phone, &user.Email); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("GetUserById", "select user_data", dur)
		r.metricsC.IncreaseExtSystemErr("database", "select")

		return nil, err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("GetUserById", "select user_data", dur)

	if err := resProfile.Scan(&firstname, &secondname, &photo, &dateBirthday); err != nil {
		r.metricsC.IncreaseExtSystemErr("database", "select")
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
	var dur time.Duration
	start := time.Now()
	if _, err := r.db.Exec(ctx, query, id, fileName); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("UpdateUserPhoto", "select user_data", dur)
		r.metricsC.IncreaseExtSystemErr("database", "insert")
		return "", err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("UpdateUserPhoto", "select user_data", dur)

	return fileName, nil
}

func (r *UserRepo) DeleteUserPhoto(ctx context.Context, id int64) error {
	query := `UPDATE user_data SET photo = '' WHERE id = $1`
	var dur time.Duration
	start := time.Now()
	if _, err := r.db.Query(ctx, query, id); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("DeleteUserPhoto", "update user_data", dur)
		r.metricsC.IncreaseExtSystemErr("database", "update")
		return err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("DeleteUserPhoto", "update user_data", dur)
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

	var dur time.Duration
	start := time.Now()
	if _, err := r.db.Exec(ctx, query, id, data.FirstName, data.SecondName); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("UpdateUserInfo", "insert user_profile_data", dur)
		r.metricsC.IncreaseExtSystemErr("database", "insert")
		fmt.Println(err.Error())
		return nil, err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("UpdateUserInfo", "insert user_profile_data", dur)

	updateQuery := `
        UPDATE user_data
        SET phone = $1, email = $2
        WHERE id = $3;
    `

	start = time.Now()
	_, err := r.db.Exec(ctx, updateQuery, data.Phone, data.Email, id)
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("UpdateUserInfo", "update user_data", dur)

	if err != nil {
		r.metricsC.IncreaseExtSystemErr("database", "update")
		fmt.Println("Error executing update query:", err)
		return nil, err
	}

	return r.GetUserById(ctx, id)
}
