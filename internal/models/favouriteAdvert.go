package models

import (
	"github.com/satori/uuid"
)

// FavouriteAdvert represents a favourite advert for a user.
type FavouriteAdvert struct {
	// ID is the unique identifier for the favourite advert.
	ID uuid.UUID `json:"id"`
	// UserID is the identifier of the user who favourited the advert.
	UserID uuid.UUID `json:"userId"`
	// AdvertID is the identifier of the advert that was favourited.
	AdvertID uuid.UUID `json:"advertId"`
	// IsDeleted is a flag indicating whether the favourite advert is deleted.
	IsDeleted bool `json:"-"`
}
