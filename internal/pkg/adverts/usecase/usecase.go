package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"context"

	"time"

	"github.com/satori/uuid"
)

// AdvertUsecase represents the usecase for manage adverts.
type AdvertUsecase struct {
	repo adverts.AdvertRepo
}

// NewAdvertUsecase creates a new instance of AdvertUsecase.
func NewAdvertUsecase(repo adverts.AdvertRepo) *AdvertUsecase {
	return &AdvertUsecase{repo: repo}
}

// CreateAdvert handles the advert creation process.
func (u *AdvertUsecase) CreateAdvert(ctx context.Context, data *models.AdvertCreateData) (*models.Advert, error) {
	newAdvert := &models.Advert{
		ID:           uuid.NewV4(),
		UserId:       data.UserId,
		Phone:        data.Phone,
		Descpription: data.Descpription,
		BuildingId:   data.BuildingId,
		CompanyId:    data.CompanyId,
		Price:        data.Price,
		Location:     data.Location,
		DataCreation: time.Now(),
		IsDeleted:    false,
	}

	if err := u.repo.CreateAdvert(ctx, newAdvert); err != nil {
		return nil, err
	}

	return newAdvert, nil
}

// GetBuildingById handles the building getting process.
func (u *AdvertUsecase) GetAdvertById(ctx context.Context, id uuid.UUID) (findAdvert *models.Advert, err error) {
	if findAdvert, err = u.repo.GetAdvertById(ctx, id); err != nil {
		return nil, err
	}

	return findAdvert, nil
}

// GetAdvertsList handles the adverts getting process.
func (u *AdvertUsecase) GetAdvertsList(ctx context.Context) (findAdverts []*models.Advert, err error) {
	if findAdverts, err = u.repo.GetAdvertsList(ctx); err != nil {
		return nil, err
	}

	return findAdverts, nil
}
