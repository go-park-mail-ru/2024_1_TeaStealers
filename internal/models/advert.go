package models

import (
	"time"

	"github.com/satori/uuid"
)

// TypePlacementAdvert represents the type of placement for an advert.
type TypePlacementAdvert string

const (
	// TypePlacementSale represents a sale placement.
	TypePlacementSale TypePlacementAdvert = "Sale"
	// TypePlacementRent represents a rent placement.
	TypePlacementRent TypePlacementAdvert = "Rent"
)

// Advert represents an advertisement.
type Advert struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"id"`
	// UserID is the identifier of the user who created the advert.
	UserID uuid.UUID `json:"userId"`
	// AdvertTypeID is the identifier of the advert type.
	AdvertTypeID uuid.UUID `json:"advertTypeId"`
	// AdvertTypePlacement is the placement type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Priority is the priority level of the advert.
	Priority int `json:"priority"`
	// DateCreation is the date when the advert was created.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the advert is deleted.
	IsDeleted bool `json:"-"`
}

// AdvertCreateData represents a data for creation advertisement.
type AdvertFlatCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID uuid.UUID `json:"userId"`
	// AdvertTypePlacement is the sale type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
	// AdvertTypePlacement is the placement type of the advert (House/Flat).
	AdvertTypePlacement AdvertTypeAdvert `json:"advertTypePlacement"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Floor is the floor of the flat.
	Floor int `json:"floor"`
	// CeilingHeight is the ceiling height of the flat.
	CeilingHeight float64 `json:"ceilingHeight"`
	// SquareGeneral is the general square of the flat.
	SquareGeneral float64 `json:"squareGeneral"`
	// SquareResidential is the residential square of the flat.
	SquareResidential float64 `json:"squareResidential"`
	// Apartment indicates if the flat is an apartment.
	Apartment bool `json:"apartment"`
	// Price is the price of the advert.
	Price int64 `json:"price"`
	// Floor is the number of floors in the building.
	FloorGeneral int `json:"floorGeneral"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
	// Address is the address of the building.
	Address string `json:"address"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"addressPoint"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// DateCreation is the date when the building was published.
}

// AdvertCreateData represents a data for creation advertisement.
type AdvertHouseCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID uuid.UUID `json:"userId"`
	// AdvertTypePlacement is the sale type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
	// AdvertTypePlacement is the placement type of the advert (House/Flat).
	AdvertTypePlacement AdvertTypeAdvert `json:"advertTypePlacement"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
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
	// Price is the price of the advert.
	Price int64 `json:"price"`
	// Floor is the number of floors in the building.
	FloorGeneral int `json:"floorGeneral"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
	// Address is the address of the building.
	Address string `json:"address"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"addressPoint"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
}
