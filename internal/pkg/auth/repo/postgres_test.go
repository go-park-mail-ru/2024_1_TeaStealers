package repo_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth/repo"
	"context"
	"database/sql"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (suite *UserRepoTestSuite) SetupTest() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	suite.Require().NoError(err)
}

func (suite *UserRepoTestSuite) TearDownTest() {
	suite.mock.ExpectClose()
	suite.Require().NoError(suite.db.Close())
}

func (suite *UserRepoTestSuite) TestCreateUser() {
	type args struct {
		user   *models.User
		userId uuid.UUID
		//	ctx    context.Context
	}
	type want struct {
		user *models.User
		err  error
	}
	id1 := uuid.NewV4()
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create user",
			args: args{
				userId: id1,
				user: &models.User{
					ID:           id1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
			},
			want: want{
				user: &models.User{
					ID:           id1,
					LevelUpdate:  1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCreateUser(tt.args.userId, tt.want.user)
			rep := repo.NewRepository(suite.db)
			newUser, gotErr := rep.CreateUser(context.Background(), tt.want.user)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newUser)
			} else {
				suite.Assert().Equal(tt.want.user, newUser)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

// rows := sqlmock.NewRows([]string{"id", "firstName", "secondName", "dateBirthday", "phone", "email", "photo"})
// rows = rows.AddRow(wantUser.ID, wantUser.FirstName,
// wantUser.SecondName, wantUser.DateBirthday, wantUser.Phone, wantUser.Email, wantUser.Photo)
func (suite *UserRepoTestSuite) setupMockCreateUser(userID uuid.UUID, wantUser *models.User) {
	rows := sqlmock.NewRows([]string{"id", "passwordhash", "phone", "email", "levelupdate"})
	rows = rows.AddRow(wantUser.ID, wantUser.Email, wantUser.Phone, wantUser.PasswordHash, wantUser.LevelUpdate)
	suite.mock.ExpectExec(`INSERT INTO users \(id, email, phone, passwordhash\) VALUES \(\$1, \$2, \$3, \$4\)`).
		WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectQuery(`SELECT id, email, phone, passwordhash, levelupdate FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnRows(rows)

}

func TestImageRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
func (suite *UserRepoTestSuite) TestGetUserByLogin() {
	type args struct {
		login  string
		user   *models.User
		userId uuid.UUID
		//	ctx    context.Context
	}
	type want struct {
		user *models.User
		err  error
	}
	id1 := uuid.NewV4()
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get user",
			args: args{
				userId: id1,
				login:  "+79003249325",
				user: &models.User{
					ID:           id1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
			},
			want: want{
				user: &models.User{
					ID:           id1,
					LevelUpdate:  1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetUserByLogin(tt.args.login, tt.want.user)
			rep := repo.NewRepository(suite.db)
			newUser, gotErr := rep.GetUserByLogin(context.Background(), tt.args.login)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newUser)
			} else {
				suite.Assert().Equal(tt.want.user, newUser)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

// rows := sqlmock.NewRows([]string{"id", "firstName", "secondName", "dateBirthday", "phone", "email", "photo"})
// rows = rows.AddRow(wantUser.ID, wantUser.FirstName,
// wantUser.SecondName, wantUser.DateBirthday, wantUser.Phone, wantUser.Email, wantUser.Photo)
func (suite *UserRepoTestSuite) setupMockGetUserByLogin(login string, wantUser *models.User) {
	rows := sqlmock.NewRows([]string{"id", "passwordhash", "phone", "email", "levelupdate"})
	rows = rows.AddRow(wantUser.ID, wantUser.Email, wantUser.Phone, wantUser.PasswordHash, wantUser.LevelUpdate)
	suite.mock.ExpectQuery(`SELECT id, email, phone, passwordhash, levelupdate FROM users WHERE email = \$1 OR phone = \$1`).
		WithArgs(login).
		WillReturnRows(rows)

}

func (suite *UserRepoTestSuite) TestGetUserLevelById() {
	type args struct {
		user   *models.User
		userId uuid.UUID
		//	ctx    context.Context
	}
	type want struct {
		user *models.User
		err  error
	}
	id1 := uuid.NewV4()
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create user",
			args: args{
				userId: id1,
				user: &models.User{
					ID:           id1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
			},
			want: want{
				user: &models.User{
					ID:           id1,
					LevelUpdate:  2,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetUserLevelById(tt.args.userId, tt.want.user)
			rep := repo.NewRepository(suite.db)
			level, gotErr := rep.GetUserLevelById(tt.want.user.ID)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Zero(level)
			} else {
				suite.Assert().Equal(tt.want.user.LevelUpdate, level)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockGetUserLevelById(userID uuid.UUID, wantUser *models.User) {
	rows := sqlmock.NewRows([]string{"levelupdate"})
	rows = rows.AddRow(wantUser.LevelUpdate)
	suite.mock.ExpectQuery(`SELECT levelupdate FROM users WHERE id = \$1`).
		WithArgs(userID).
		WillReturnRows(rows)
}
