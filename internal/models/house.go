package models

import (
	"time"

	"github.com/google/uuid"
)

// StatusAreaHouse represents the status area of a house.
type StatusAreaHouse string

const (
	// StatusAreaIHC represents an IHC status area.
	StatusAreaIHC StatusAreaHouse = "IHC"
	// StatusAreaDNP represents a DNP status area.
	StatusAreaDNP StatusAreaHouse = "DNP"
	// StatusAreaG represents a G status area.
	StatusAreaG StatusAreaHouse = "G"
	// StatusAreaF represents a F status area.
	StatusAreaF StatusAreaHouse = "F"
	// StatusAreaPSP represents a PSP status area.
	StatusAreaPSP StatusAreaHouse = "PSP"
)

// StatusHomeHouse represents the status home of a house.
type StatusHomeHouse string

const (
	// StatusHomeLive represents a Live status home.
	StatusHomeLive StatusHomeHouse = "Live"
	// StatusHomeRepairNeed represents a RepairNeed status home.
	StatusHomeRepairNeed StatusHomeHouse = "RepairNeed"
	// StatusHomeCompleteNeed represents a CompleteNeed status home.
	StatusHomeCompleteNeed StatusHomeHouse = "CompleteNeed"
	// StatusHomeRenovation represents a Renovation status home.
	StatusHomeRenovation StatusHomeHouse = "Renovation"
)

// House represents a house entity.
type House struct {
	// ID is the unique identifier for the house.
	ID uuid.UUID `json:"id"`
	// BuildingID is the identifier of the building to which the house belongs.
	BuildingID uuid.UUID `json:"buildingId"`
	// AdvertTypeID is the identifier of the advert type of the house.
	AdvertTypeID uuid.UUID `json:"advertTypeId"`
	// CeilingHeight is the ceiling height of the house.
	CeilingHeight int `json:"ceilingHeight"`
	// SquareArea is the square area of the house.
	SquareArea float64 `json:"squareArea"`
	// SquareHouse is the square area of the house.
	SquareHouse float64 `json:"squareHouse"`
	// BedroomCount is the number of bedrooms in the house.
	BedroomCount int `json:"bedroomCount"`
	// StatusArea is the status area of the house.
	StatusArea StatusAreaHouse `json:"statusArea"`
	// Cottage indicates if the house is a cottage.
	Cottage bool `json:"cottage"`
	// StatusHome is the status home of the house.
	StatusHome StatusHomeHouse `json:"statusHome"`
	// DateCreation is the date when the house was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the house is deleted.
	IsDeleted bool `json:"-"`
}
