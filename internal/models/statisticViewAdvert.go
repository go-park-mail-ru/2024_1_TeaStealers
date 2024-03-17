package models

import (
	"time"

	"github.com/satori/uuid"
)

// StatisticViewAdvert represents a view statistic for an advert by a user.
type StatisticViewAdvert struct {
	// ID is the unique identifier for the view statistic.
	ID uuid.UUID `json:"id"`
	// UserID is the identifier of the user who viewed the advert.
	UserID uuid.UUID `json:"userId"`
	// AdvertID is the identifier of the advert that was viewed.
	AdvertID uuid.UUID `json:"advertId"`
	// DateCreation is the date when the view statistic was created.
	DateCreation time.Time `json:"-"`
}
