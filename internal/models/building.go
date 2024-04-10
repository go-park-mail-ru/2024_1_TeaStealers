package models

import (
	"html"
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
	Address string `json:"adress"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"adressPoint"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// DateCreation is the date when the building was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the building is deleted.
	IsDeleted bool `json:"-"`
}

func (building *Building) Sanitize() {
	building.Address = html.EscapeString(building.Address)
	building.AddressPoint = html.EscapeString(building.AddressPoint)
	building.Material = MaterialBuilding(html.EscapeString(string(building.Material)))
}

// BuildingCreateData represents a data for creation building.
type BuildingCreateData struct {
	// ComplexID is the identifier of the complex to which the building belongs.
	ComplexID uuid.UUID `json:"complexId"`
	// Floor is the number of floors in the building.
	Floor int `json:"floor"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
	// Address is the address of the building.
	Address string `json:"adress"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"adressPoint"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
}

func (buildCrDat *BuildingCreateData) Sanitize() {
	buildCrDat.Address = html.EscapeString(buildCrDat.Address)
	buildCrDat.AddressPoint = html.EscapeString(buildCrDat.AddressPoint)
	buildCrDat.Material = MaterialBuilding(html.EscapeString(string(buildCrDat.Material)))
}

// BuildingData represents an exists buildings with concrete adress.
type BuildingData struct {
	// ID is the unique identifier for the building.
	ID uuid.UUID `json:"id"`
	// ComplexName is the name of the complex to which the building belongs.
	ComplexName string `json:"complexName"`
	// Floor is the number of floors in the building.
	Floor int `json:"floor"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
	// Address is the address of the building.
	Address string `json:"adress"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"adressPoint"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
}

func (buildDat *BuildingData) Sanitize() {
	buildDat.ComplexName = html.EscapeString(buildDat.ComplexName)
	buildDat.Address = html.EscapeString(buildDat.Address)
	buildDat.AddressPoint = html.EscapeString(buildDat.AddressPoint)
	buildDat.Material = MaterialBuilding(html.EscapeString(string(buildDat.Material)))
}
