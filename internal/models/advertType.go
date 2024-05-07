package models

import (
	"html"
	"time"

	"github.com/satori/uuid"
)

// AdvertTypeAdvert represents the type of advertisement.
type AdvertTypeAdvert string

const (
	// AdvertTypeHouse represents a house advertisement.
	AdvertTypeHouse AdvertTypeAdvert = "House"
	// AdvertTypeFlat represents a flat advertisement.
	AdvertTypeFlat AdvertTypeAdvert = "Flat"
)

// AdvertType represents an advertisement type.
type AdvertType struct {
	// ID of the advert type.
	ID uuid.UUID `json:"id"`
	// AdvertType is the type of advertisement (House/Flat).
	AdvertType AdvertTypeAdvert `json:"advertType"`
	// DateCreation is the date of creation of the advert type.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating if the advert type is deleted.
	IsDeleted bool `json:"-"`
}

// AdvertType represents an advertisement type flat.
type HouseTypeAdvert struct {
	// HouseID of the advert type.
	HouseID int64 `json:"houseId"`
	// AdvertID of the advert type.
	AdvertID int64 `json:"advertId"`
	// IsDeleted is a flag indicating if the advert type is deleted.
	IsDeleted bool `json:"-"`
}

// AdvertType represents an advertisement type flat.
type FlatTypeAdvert struct {
	// FlatID of the advert type.
	FlatID int64 `json:"flatId"`
	// AdvertID of the advert type.
	AdvertID int64 `json:"advertId"`
	// IsDeleted is a flag indicating if the advert type is deleted.
	IsDeleted bool `json:"-"`
}

func (advType *AdvertType) Sanitize() {
	advType.AdvertType = AdvertTypeAdvert(html.EscapeString(string(advType.AdvertType)))
}
