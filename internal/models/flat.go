package models

import (
	"time"

	"github.com/google/uuid"
)

// Flat represents a flat entity.
type Flat struct {
	// ID is the unique identifier for the flat.
	ID uuid.UUID `json:"id"`
	// BuildingID is the identifier of the building to which the flat belongs.
	BuildingID uuid.UUID `json:"buildingId"`
	// AdvertTypeID is the identifier of the advert type of the flat.
	AdvertTypeID uuid.UUID `json:"advertTypeId"`
	// Floor is the floor of the flat.
	Floor int `json:"floor"`
	// CeilingHeight is the ceiling height of the flat.
	CeilingHeight float64 `json:"ceilingHeight"`
	// SquareGeneral is the general square of the flat.
	SquareGeneral float64 `json:"squareGeneral"`
	// SquareResidential is the residential square of the flat.
	SquareResidential float64 `json:"squareResidential"`
	// Apartment indicates if the flat is an apartment.
	Apartment bool `json:"apartment"`
	// DateCreation is the date when the flat was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the flat is deleted.
	IsDeleted bool `json:"-"`
}
