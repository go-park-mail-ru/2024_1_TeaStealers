package models

import (
	"time"

	"github.com/google/uuid"
)

// AdvertTypeAdvert represents the type of advertisement.
type AdvertTypeAdvert string

const (
	// AdvertTypeHouse represents a house advertisement.
	AdvertTypeHouse AdvertTypeAdvert = "house"
	// AdvertTypeFlat represents a flat advertisement.
	AdvertTypeFlat AdvertTypeAdvert = "flat"
)

// AdvertType represents an advertisement type.
type AdvertType struct {
	// ID of the advert type.
	ID uuid.UUID `json:"id"`
	// AdvertType is the type of advertisement (house/flat).
	AdvertType AdvertTypeAdvert `json:"advertType"`
	// DateCreation is the date of creation of the advert type.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating if the advert type is deleted.
	IsDeleted bool `json:"-"`
}
