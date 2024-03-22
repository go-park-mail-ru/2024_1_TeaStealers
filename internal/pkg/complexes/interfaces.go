package complexes

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"

	"github.com/satori/uuid"
)

// ComplexUsecase represents the usecase interface for complexes.
type ComplexUsecase interface {
	CreateComplex(ctx context.Context, data *models.ComplexCreateData) (*models.Complex, error)
	CreateBuilding(ctx context.Context, data *models.BuildingCreateData) (*models.Building, error)
	UpdateComplexPhoto(file io.Reader, fileType string, id uuid.UUID) (string, error)
	GetComplexById(ctx context.Context, id uuid.UUID) (foundComplex *models.ComplexData, err error)
}

// ComplexRepo represents the repository interface for complexes.
type ComplexRepo interface {
	CreateComplex(ctx context.Context, company *models.Complex) (*models.Complex, error)
	CreateBuilding(ctx context.Context, complex *models.Building) (*models.Building, error)
	UpdateComplexPhoto(id uuid.UUID, fileName string) (string, error)
	GetComplexById(ctx context.Context, complexId uuid.UUID) (*models.ComplexData, error)
}
