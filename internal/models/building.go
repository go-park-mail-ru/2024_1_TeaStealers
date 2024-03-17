package models

import (
	"time"

	"github.com/satori/uuid"
)

// MaterialBuilding represents the material of a building.
type MaterialBuilding string

const (
	// MaterialBrick represents a brick material.
	MaterialBrick MaterialBuilding = "Brick"
	// MaterialMonolithic represents a monolithic material.
	MaterialMonolithic MaterialBuilding = "Monolithic"
	// MaterialWood represents a wood material.
	MaterialWood MaterialBuilding = "Wood"
	// MaterialPanel represents a panel material.
	MaterialPanel MaterialBuilding = "Panel"
	// MaterialStalinsky represents a stalinsky material.
	MaterialStalinsky MaterialBuilding = "Stalinsky"
	// MaterialBlock represents a block material.
	MaterialBlock MaterialBuilding = "Block"
	// MaterialMonolithicBlock represents a monolithic block material.
	MaterialMonolithicBlock MaterialBuilding = "MonolithicBlock"
	// MaterialFrame represents a frame material.
	MaterialFrame MaterialBuilding = "Frame"
	// MaterialAeratedConcreteBlock represents an aerated concrete block material.
	MaterialAeratedConcreteBlock MaterialBuilding = "AeratedConcreteBlock"
	// MaterialGasSilicateBlock represents a gas silicate block material.
	MaterialGasSilicateBlock MaterialBuilding = "GasSilicateBlock"
	// MaterialFoamConcreteBlock represents a foam concrete block material.
	MaterialFoamConcreteBlock MaterialBuilding = "FoamConcreteBlock"
)

// Building represents a building entity.
type Building struct {
	// ID is the unique identifier for the building.
	ID uuid.UUID `json:"id"`
	// ComplexID is the identifier of the complex to which the building belongs.
	ComplexID uuid.UUID `json:"complexId"`
	// Floor is the number of floors in the building.
	Floor int `json:"floor"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
	// Address is the address of the building.
	Address string `json:"address"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"addressPoint"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// DateCreation is the date when the building was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the building is deleted.
	IsDeleted bool `json:"-"`
}
