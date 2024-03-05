package adverts

import (
	"2024_1_TeaStealers/internal/models"
	"context"

	"github.com/satori/uuid"
)

// AdvertUsecase represents the usecase interface for manage adverts.
type AdvertUsecase interface {
	CreateAdvert(ctx context.Context, data *models.AdvertCreateData) (*models.Advert, error)
	GetAdvertById(ctx context.Context, id uuid.UUID) (findAdvert *models.Advert, err error)
	GetAdvertsList(ctx context.Context) (findAdverts []*models.Advert, err error)
	DeleteAdvertById(ctx context.Context, id uuid.UUID) (err error)
	UpdateAdvertById(ctx context.Context, body map[string]interface{}, id uuid.UUID) (err error)
}

// AdvertRepo represents the repository interface for manage adverts.
type AdvertRepo interface {
	CreateAdvert(ctx context.Context, advert *models.Advert) error
	GetAdvertById(ctx context.Context, id uuid.UUID) (*models.Advert, error)
	GetAdvertsList(ctx context.Context) ([]*models.Advert, error)
	DeleteAdvertById(ctx context.Context, id uuid.UUID) error
	UpdateAdvertById(ctx context.Context, body map[string]interface{}, id uuid.UUID) (err error)
}
