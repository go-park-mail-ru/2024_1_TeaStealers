package models

import (
	"time"

	"github.com/satori/uuid"
)

// Building represents company information
type Building struct {
	// ID uniquely identifies the building.
	ID uuid.UUID `json:"id"`
	// Location is the location of the building.
	Location string `json:"location"`
	// Description is the description of the building.
	Descpription string `json:"description"`
	// DataCreation is the time of adding a record to the database.
	DataCreation time.Time `json:"datacreation"`
	// isDeleted means is the building deleted?.
	IsDeleted bool `json:"isdeleted"`
}

// BuildingCreateData represents building information for location and description
type BuildingCreateData struct {
	// Location is the location of the building.
	Location string `json:"location"`
	// Description is the description of the building.
	Descpription string `json:"description"`
}
