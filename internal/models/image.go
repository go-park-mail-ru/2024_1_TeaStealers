package models

import (
	"html"
	"time"
)

// Image represents an image associated with an advert.
type Image struct {
	// ID is the unique identifier for the image.
	ID int64 `json:"id"`
	// AdvertID is the identifier of the advert to which the image belongs.
	AdvertID int64 `json:"advertId"`
	// Photo is the filename of the image.
	Photo string `json:"photo"`
	// Priority is the priority of the image.
	Priority int `json:"priority"`
	// DateCreation is the date when the image was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the image is deleted.
	IsDeleted bool `json:"-"`
}

func (imag *Image) Sanitize() {
	imag.Photo = html.EscapeString(imag.Photo)
}

// ImageResp represents an image response.
type ImageResp struct {
	// ID is the unique identifier for the image.
	ID int64 `json:"id"`
	// Photo is the filename of the image.
	Photo string `json:"photo"`
	// Priority is the priority of the image.
	Priority int `json:"priority"`
}

func (imagResp *ImageResp) Sanitize() {
	imagResp.Photo = html.EscapeString(imagResp.Photo)
}
