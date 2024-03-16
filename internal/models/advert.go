package models

import (
	"time"

	"github.com/google/uuid"
)

// TypePlacementAdvert represents the type of placement for an advert.
type TypePlacementAdvert string

const (
	// TypePlacementSale represents a sale placement.
	TypePlacementSale TypePlacementAdvert = "Sale"
	// TypePlacementRent represents a rent placement.
	TypePlacementRent TypePlacementAdvert = "Rent"
)

// Advert represents an advertisement.
type Advert struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"id"`
	// UserID is the identifier of the user who created the advert.
	UserID uuid.UUID `json:"userId"`
	// AdvertTypeID is the identifier of the advert type.
	AdvertTypeID uuid.UUID `json:"advertTypeId"`
	// AdvertTypePlacement is the placement type of the advert (Sale/Rent).
	AdvertTypePlacement TypePlacementAdvert `json:"advertTypePlacement"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Priority is the priority level of the advert.
	Priority int `json:"priority"`
	// DateCreation is the date when the advert was created.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the advert is deleted.
	IsDeleted bool `json:"-"`
}
