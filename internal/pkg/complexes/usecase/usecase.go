package usecase

import (
	"2024_1_TeaStealers/internal/models"
	complex "2024_1_TeaStealers/internal/pkg/complexes"
	"context"

	"go.uber.org/zap"
)

// ComplexUsecase represents the usecase for complex using.
type ComplexUsecase struct {
	repo   complex.ComplexRepo
	logger *zap.Logger
}

// NewComplexUsecase creates a new instance of ComplexUsecase.
func NewComplexUsecase(repo complex.ComplexRepo, logger *zap.Logger) *ComplexUsecase {
	return &ComplexUsecase{repo: repo, logger: logger}
}

// CreateComplex handles the complex registration process.
func (u *ComplexUsecase) CreateComplex(ctx context.Context, data *models.ComplexCreateData) (*models.Complex, error) {
	newComplex := &models.Complex{
		CompanyId:              data.CompanyId,
		Name:                   data.Name,
		Description:            data.Description,
		Address:                data.Address,
		DateBeginBuild:         data.DateBeginBuild,
		DateEndBuild:           data.DateEndBuild,
		WithoutFinishingOption: data.WithoutFinishingOption,
		FinishingOption:        data.FinishingOption,
		PreFinishingOption:     data.PreFinishingOption,
		ClassHousing:           data.ClassHousing,
		Parking:                data.Parking,
		Security:               data.Security,
	}

	complex, err := u.repo.CreateComplex(ctx, newComplex)
	if err != nil {
		return nil, err
	}

	return complex, nil
}

func (u *ComplexUsecase) UpdateComplexPhoto(ctx context.Context, id int64, filename string) (string, error) {
	fileName, err := u.repo.UpdateComplexPhoto(ctx, id, filename)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

// GetComplexById handles the getting complex advert process.
func (u *ComplexUsecase) GetComplexById(ctx context.Context, id int64) (foundComplexData *models.ComplexData, err error) {

	if foundComplexData, err = u.repo.GetComplexById(ctx, id); err != nil {
		return nil, err
	}

	return foundComplexData, nil
}

// CreateCompany handles the company registration process.
func (u *ComplexUsecase) CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error) {
	newCompany := &models.Company{
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

func (u *ComplexUsecase) UpdateCompanyPhoto(ctx context.Context, id int64, filename string) (string, error) {
	fileName, err := u.repo.UpdateCompanyPhoto(ctx, id, filename)
	if err != nil {
		return "", nil
	}
	return fileName, nil
}

// GetCompanyById handles the getting company advert process.
func (u *ComplexUsecase) GetCompanyById(ctx context.Context, id int64) (foundCompanyData *models.CompanyData, err error) {

	if foundCompanyData, err = u.repo.GetCompanyById(ctx, id); err != nil {
		return nil, err
	}

	return foundCompanyData, nil
}
