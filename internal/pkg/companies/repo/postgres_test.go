package repo_test

/*
import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies/repo"
	"context"
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
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

func (suite *UserRepoTestSuite) TestCreateCompany() {
	type args struct {
		comp *models.Company
		// ctx context.Context
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		// user *models.User
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create company",
			args: args{
				comp: &models.Company{ // дата не возвращается!
					ID:          uuid.NewV4(),
					Name:        "Maxorella Company",
					YearFounded: 1999,
					Phone:       "+79091535823",
					Description: "best company ever description",
				},
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error insert",
			args: args{
				comp: &models.Company{ // дата не возвращается!
					ID:          uuid.NewV4(),
					Name:        "Maxorella Company",
					YearFounded: 1999,
					Phone:       "+79091535823",
					Description: "best company ever description",
				},
				errExec:  errors.New("error"),
				errQuery: nil,
				expExec:  true,
				expQuery: false,
			},
			want: want{
				err: errors.New("error"),
			},
		},
		{
			name: "error query",
			args: args{
				comp: &models.Company{ // дата не возвращается!
					ID:          uuid.NewV4(),
					Name:        "Maxorella Company",
					YearFounded: 1999,
					Phone:       "+79091535823",
					Description: "best company ever description",
				},
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCreateCompany(tt.args.comp, tt.args.errExec, tt.args.errQuery, tt.args.expExec, tt.args.expQuery)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			newCompany, gotErr := rep.CreateCompany(context.Background(), tt.args.comp)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newCompany)
			} else {
				suite.Assert().Equal(tt.args.comp, newCompany)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockCreateCompany(company *models.Company, errExec, errQuery error, expExec, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "name", "yearFounded", "phone", "description"})
	rows = rows.AddRow(company.ID, company.Name, company.YearFounded, company.Phone, company.Description)
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO companies \(id, name, photo, yearFounded, phone, description\) VALUES \(\$1, \$2, '', \$3, \$4, \$5\);`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(company.ID, company.Name, company.YearFounded, company.Phone, company.Description)
	}
	if epxQuery {
		suite.mock.ExpectQuery(`SELECT id, name, yearFounded, phone, description FROM companies WHERE id = \$1`).
			WithArgs(company.ID).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func (suite *UserRepoTestSuite) TestUpdatePhoto() {
	type args struct {
		id       uuid.UUID
		filename string
		errExec  error
		expExec  bool
	}
	type want struct {
		err error
	}
	id1 := uuid.NewV4()
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful update photo company",
			args: args{
				id:       id1,
				filename: "file/name",
				errExec:  nil,
				expExec:  true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail update photo company",
			args: args{
				id:       id1,
				filename: "file/name",
				errExec:  errors.New("some error"),
				expExec:  true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockUpdateCompanyPhoto(tt.args.id, tt.args.filename, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			newFilename, gotErr := rep.UpdateCompanyPhoto(tt.args.id, tt.args.filename)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newFilename)
			} else {
				suite.Assert().Equal(tt.args.filename, newFilename)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockUpdateCompanyPhoto(idi uuid.UUID, filen string, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`UPDATE companies SET photo = \$1 WHERE id = \$2`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(filen, idi)
	}
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) TestGetCompanyById() {
	type args struct {
		// ctx context.Context
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		// user *models.User
		err      error
		compData *models.CompanyData
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get company",
			args: args{

				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				compData: &models.CompanyData{ // дата не возвращается!
					ID:          uuid.NewV4(),
					Name:        "Maxorella Company",
					YearFounded: 1999,
					Phone:       "+79091535823",
					Description: "best company ever description",
				},
				err: nil,
			},
		},
		{
			name: "fail get company",
			args: args{
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				compData: &models.CompanyData{ // дата не возвращается!
					ID:          uuid.NewV4(),
					Name:        "Maxorella Company",
					YearFounded: 1999,
					Phone:       "+79091535823",
					Description: "best company ever description",
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetCompanyById(tt.want.compData, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			updCompany, gotErr := rep.GetCompanyById(context.Background(), tt.want.compData.ID)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(updCompany)
			} else {
				suite.Assert().Equal(tt.want.compData, updCompany)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockGetCompanyById(companyData *models.CompanyData, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "photo", "name", "yearFounded", "phone", "description"})
	rows = rows.AddRow(companyData.ID, companyData.Photo, companyData.Name, companyData.YearFounded, companyData.Phone, companyData.Description)

	if epxQuery {
		suite.mock.ExpectQuery(`SELECT id, photo, name, yearfounded, phone, description FROM companies WHERE id = \$1`).
			WithArgs(companyData.ID).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}
*/
