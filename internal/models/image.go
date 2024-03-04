package models

import (
	"time"

	"github.com/satori/uuid"
)

// Image represents image information
type Image struct {
	// ID uniquely identifies the images.
	ID uuid.UUID `json:"id"`
	// Filename is the filename of the image.
	Filename string `json:"path"`
	// AdvertId is the id of the advert to which the image belongs.
	AdvertId uuid.UUID `json:"advertid"`
	// Priority is the priority image locations.
	Priority int `json:"priority"`
	// DataCreation is the time of adding a record to the database.
	DataCreation time.Time `json:"datacreation"`
	// isDeleted means is the image deleted?.
	IsDeleted bool `json:"isdeleted"`
}

// ImageCreateData represents image information for advertId, priority
type ImageCreateData struct {
	// AdvertId stands which id of advert this image standing
	AdvertId uuid.UUID `json:"advertid"`
	// Priority stands for company priority
	Priority int `json:"priority"`
}
