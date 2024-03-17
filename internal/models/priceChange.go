package models

import (
	"time"

	"github.com/satori/uuid"
)

// PriceChange represents a change in the price of an advert.
type PriceChange struct {
	// ID is the unique identifier for the price change.
	ID uuid.UUID `json:"id"`
	// AdvertID is the identifier of the advert for which the price changed.
	AdvertID uuid.UUID `json:"advertId"`
	// Price is the new price of the advert.
	Price int64 `json:"price"`
	// DateCreation is the date when the price change was created.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the price change is deleted.
	IsDeleted bool `json:"-"`
}
