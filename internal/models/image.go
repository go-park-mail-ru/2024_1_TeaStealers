package models

import (
	"time"

	"github.com/satori/uuid"
)

// Image represents image information
type Image struct {
	// ID uniquely identifies the images.
	ID uuid.UUID `json:"id"`
	// Path is the path of the image.
	Path string `json:"path"`
	// AdvertId is the id of the advert to which the image belongs.
	AdvertId uint `json:"advertid"`
	// Priority is the priority image locations.
	Priority int `json:"priority"`
	// DataCreation is the time of adding a record to the database.
	DataCreation time.Time `json:"datacreation"`
	// isDeleted means is the image deleted?.
	IsDeleted bool `json:"isdeleted"`
}
