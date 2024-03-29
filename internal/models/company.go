package models

import (
	"time"

	"github.com/satori/uuid"
)

// Company represents company information
type Company struct {
	// ID uniquely identifies the company.
	ID uuid.UUID `json:"id"`
	// Name is the name of the company.
	Name string `json:"name"`
	// Phone is the phone of the company.
	Phone int `json:"phone"`
	// Description is the description of the company.
	Descpription string `json:"description"`
	// DataCreation is the time of adding a record to the database.
	DataCreation time.Time `json:"datacreation"`
	// isDeleted means is the company deleted?.
	IsDeleted bool `json:"isdeleted"`
}

// CompanyCreateData represents company information for name, phone and description
type CompanyCreateData struct {
	// Name stands for company name
	Name string `json:"name"`
	// Phone stands for company phone
	Phone int `json:"phone"`
	// Descpription stands for company description
	Descpription string `json:"description"`
}
