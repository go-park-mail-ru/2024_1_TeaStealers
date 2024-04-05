package repo_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/users/repo"
	"database/sql"
	"errors"
	"testing"
	"time"

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

func (suite *UserRepoTestSuite) TestGetUserById() {
	type args struct {
		userId uuid.UUID
	}
	type want struct {
		user *models.User
		err  error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get user",
			args: args{
				userId: uuid.NewV4(),
			},
			want: want{
				user: &models.User{
					ID:           uuid.NewV4(),
					FirstName:    "Maksim",
					SecondName:   "Shagaev",
					DateBirthday: time.Date(1990, 11, 4, 12, 20, 10, 0, time.Local),
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					Photo:        "/path/to/photo/test.jpg",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetUserByID(tt.args.userId, tt.want.user)
			rep := repo.NewRepository(suite.db)
			gotUser, gotErr := rep.GetUserById(tt.args.userId)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotUser)
			} else {
				suite.Assert().Equal(tt.want.user, gotUser)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockGetUserByID(userID uuid.UUID, wantUser *models.User) {
	rows := sqlmock.NewRows([]string{"id", "firstName", "secondName", "dateBirthday", "phone", "email", "photo"})
	rows = rows.AddRow(wantUser.ID, wantUser.FirstName,
		wantUser.SecondName, wantUser.DateBirthday, wantUser.Phone, wantUser.Email, wantUser.Photo)
	suite.mock.ExpectQuery(`SELECT id, firstname, secondname, datebirthday, phone, email, photo FROM users WHERE id=\$1`).
		WithArgs(userID).
		WillReturnRows(rows)
}

func TestImageRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) TestUpdateUserInfo() {
	type args struct {
		// data   *models.UserUpdateData
		userId uuid.UUID
	}
	type want struct {
		user *models.User
		err  error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful update user",
			args: args{
				userId: uuid.NewV4(),
			},
			want: want{
				user: &models.User{
					ID:         uuid.NewV4(),
					FirstName:  "Maksim",
					SecondName: "Shagaev",
					Phone:      "+79003249325",
					Email:      "my@mail.ru",
					Photo:      "/path/to/photo/test.jpg",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockUpdateUserInfo(tt.args.userId, tt.want.user)
			rep := repo.NewRepository(suite.db)
			gotUser, gotErr := rep.UpdateUserInfo(tt.args.userId, &models.UserUpdateData{})
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotUser)
			} else {
				suite.Assert().Equal(tt.want.user, gotUser)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockUpdateUserInfo(userID uuid.UUID, wantUser *models.User) {
	rows := sqlmock.NewRows([]string{"id", "firstName", "secondName", "dateBirthday", "phone", "email", "photo"})
	rows = rows.AddRow(wantUser.ID, wantUser.FirstName,
		wantUser.SecondName, wantUser.DateBirthday, wantUser.Phone, wantUser.Email, wantUser.Photo)
	suite.mock.ExpectExec(`UPDATE users SET firstname = \$1, secondname = \$2, datebirthday = \$3, phone = \$4, email = \$5 WHERE id = \$6`).
		WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectQuery(`SELECT id, firstname, secondname, datebirthday, phone, email, photo FROM users WHERE id = \$1`).
		WithArgs(userID).WillReturnRows(rows)
}

func (suite *UserRepoTestSuite) TestCheckUserPassword() {
	type args struct {
		userId       uuid.UUID
		passHash     string
		passHashMock string
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful check user password",
			args: args{
				userId:       uuid.NewV4(),
				passHash:     "passwordhash1",
				passHashMock: "passwordhash1",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "different user passwords",
			args: args{
				userId:       uuid.NewV4(),
				passHash:     "HASH111",
				passHashMock: "ANOTHERHASH",
			},
			want: want{
				err: errors.New("passwords don't match"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCheckUserPassword(tt.args.userId, tt.args.passHashMock)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CheckUserPassword(tt.args.userId, tt.args.passHash)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCheckUserPassword(userID uuid.UUID, hash string) {
	rows := sqlmock.NewRows([]string{"hash"})
	rows = rows.AddRow(hash)
	suite.mock.ExpectQuery(`SELECT passwordhash FROM users WHERE id = \$1`).
		WithArgs(userID).WillReturnRows(rows)
}

func (suite *UserRepoTestSuite) TestUpdateUserPassword() {
	type args struct {
		userId      uuid.UUID
		newpassHash string
		// prevlevel   int
	}
	type want struct {
		err      error
		newlevel int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful update user password",
			args: args{
				userId:      uuid.NewV4(),
				newpassHash: "passwordhash1",
			},
			want: want{
				err:      nil,
				newlevel: 2,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockUpdateUserPassword(tt.args.userId, rune(tt.want.newlevel))
			rep := repo.NewRepository(suite.db)
			gotLevel, gotErr := rep.UpdateUserPassword(tt.args.userId, tt.args.newpassHash)
			suite.Assert().Equal(tt.want.newlevel, gotLevel)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockUpdateUserPassword(userID uuid.UUID, levelupdate rune) {
	rows := sqlmock.NewRows([]string{"hash"})
	rows = rows.AddRow(levelupdate)
	suite.mock.ExpectExec(`UPDATE users SET passwordhash=\$1, levelupdate = levelupdate\+1 WHERE id = \$2`).
		WillReturnError(nil).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectQuery(`SELECT levelupdate FROM users WHERE id = \$1`).
		WithArgs(userID).WillReturnRows(rows)
}
