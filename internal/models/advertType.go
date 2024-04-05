package models

import (
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
