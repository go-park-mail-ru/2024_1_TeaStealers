package models

import (
	"html"
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

func (comp *Company) Sanitize() {
	comp.Photo = html.EscapeString(comp.Photo)
	comp.Name = html.EscapeString(comp.Name)
	comp.Phone = html.EscapeString(comp.Phone)
	comp.Description = html.EscapeString(comp.Description)
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

func (compCrDat *CompanyCreateData) Sanitize() {
	compCrDat.Name = html.EscapeString(compCrDat.Name)
	compCrDat.Phone = html.EscapeString(compCrDat.Phone)
	compCrDat.Description = html.EscapeString(compCrDat.Description)
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

func (compDat *CompanyData) Sanitize() {
	compDat.Photo = html.EscapeString(compDat.Photo)
	compDat.Name = html.EscapeString(compDat.Name)
	compDat.Phone = html.EscapeString(compDat.Phone)
	compDat.Description = html.EscapeString(compDat.Description)
}
