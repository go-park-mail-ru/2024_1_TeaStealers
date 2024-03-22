package models

import (
	"time"

	"github.com/satori/uuid"
)

// Company represents a company entity.
type Company struct {
	// ID is the unique identifier for the company.
	ID uuid.UUID `json:"id"`
	// Photo is the filename of the company's photo.
	Photo string `json:"photo"`
	// Name is the name of the company.
	Name string `json:"name"`
	// YearFounded is the year when the company was founded.
	YearFounded int `json:"yearFounded"`
	// Phone is the phone number of the company.
	Phone string `json:"phone"`
	// Description is the description of the company.
	Description string `json:"description"`
	// DateCreation is the date when the company was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the company is deleted.
	IsDeleted bool `json:"-"`
}

// CompanyCreateData represents a data for creation company.
type CompanyCreateData struct {
	// Name is the name of the company.
	Name string `json:"name"`
	// YearFounded is the year when the company was founded.
	YearFounded int `json:"yearFounded"`
	// Phone is the phone number of the company.
	Phone string `json:"phone"`
	// Description is the description of the company.
	Description string `json:"description"`
}

// Company represents a company information.
type CompanyData struct {
	// ID is the unique identifier for the company.
	ID uuid.UUID `json:"id"`
	// Photo is the filename of the company's photo.
	Photo string `json:"photo"`
	// Name is the name of the company.
	Name string `json:"name"`
	// YearFounded is the year when the company was founded.
	YearFounded int `json:"yearFounded"`
	// Phone is the phone number of the company.
	Phone string `json:"phone"`
	// Description is the description of the company.
	Description string `json:"description"`
}
