package models

import (
	"time"

	"github.com/satori/uuid"
)

// Image represents an image associated with an advert.
type Image struct {
	// ID is the unique identifier for the image.
	ID uuid.UUID `json:"id"`
	// AdvertID is the identifier of the advert to which the image belongs.
	AdvertID uuid.UUID `json:"advertId"`
	// Photo is the filename of the image.
	Photo string `json:"photo"`
	// Priority is the priority of the image.
	Priority int `json:"priority"`
	// DateCreation is the date when the image was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the image is deleted.
	IsDeleted bool `json:"-"`
}
