package models

import (
	"time"
)

// PriceChange represents a change in the price of an advert.
type PriceChange struct {
	// ID is the unique identifier for the price change.
	ID int64 `json:"id"`
	// AdvertID is the identifier of the advert for which the price changed.
	AdvertID int64 `json:"advertId"`
	// Price is the new price of the advert.
	Price int64 `json:"price"`
	// DateCreation is the date when the price change was created.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the price change is deleted.
	IsDeleted bool `json:"-"`
}

// PriceChangeData represents a change in the price of an advert.
type PriceChangeData struct {
	// Price is the new price of the advert.
	Price int64 `json:"price"`
	// DateCreation is the date when the price change was created.
	DateCreation time.Time `json:"data"`
}
