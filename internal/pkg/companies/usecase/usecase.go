package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// CompanyUsecase represents the usecase for company using.
type CompanyUsecase struct {
	repo   companies.CompanyRepo
	logger *zap.Logger
}

// NewCompanyUsecase creates a new instance of CompanyUsecase.
func NewCompanyUsecase(repo companies.CompanyRepo, logger *zap.Logger) *CompanyUsecase {
	return &CompanyUsecase{repo: repo, logger: logger}
}

// CreateCompany handles the company registration process.
func (u *CompanyUsecase) CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error) {
	newCompany := &models.Company{
		ID:          uuid.NewV4(),
		Name:        data.Name,
		Description: data.Description,
		Phone:       data.Phone,
		YearFounded: data.YearFounded,
	}

	company, err := u.repo.CreateCompany(ctx, newCompany)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (u *CompanyUsecase) UpdateCompanyPhoto(file io.Reader, fileType string, id uuid.UUID) (string, error) {
	newId := uuid.NewV4()
	newFileName := newId.String() + fileType
	subDirectory := "companies"
	directory := filepath.Join(os.Getenv("DOCKER_DIR"), subDirectory)
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", err
	}
	destination, err := os.Create(directory + "/" + newFileName)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		return "", err
	}
	fileName, err := u.repo.UpdateCompanyPhoto(id, subDirectory+"/"+newFileName)
	if err != nil {
		return "", nil
	}
	return fileName, nil
}

// GetCompanyById handles the getting company advert process.
func (u *CompanyUsecase) GetCompanyById(ctx context.Context, id uuid.UUID) (foundCompanyData *models.CompanyData, err error) {

	if foundCompanyData, err = u.repo.GetCompanyById(ctx, id); err != nil {
		return nil, err
	}

	return foundCompanyData, nil
}
