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

// AdvertFlatCreateData represents a data for creation advertisement.
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
	// RoomCount is the how many rooms in flat
	RoomCount int `json:"roomCount"`
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

// ComplexAdvertFlatCreateData represents a data for creation advertisement.
type ComplexAdvertFlatCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID uuid.UUID `json:"userId"`
	// BuildingID is the identifier of the building to which the flat belongs.
	BuildingID uuid.UUID `json:"buildingId"`
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
	// RoomCount is the how many rooms in flat
	RoomCount int `json:"roomCount"`
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

// AdvertHouseCreateData represents a data for creation advertisement.
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

// ComplexAdvertHouseCreateData represents a data for creation advertisement.
type ComplexAdvertHouseCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID uuid.UUID `json:"userId"`
	// BuildingID is the identifier of the building to which the house belongs.
	BuildingID uuid.UUID `json:"buildingId"`
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

// AdvertSquareData represents the structure of the JSON data for square advert.
type AdvertSquareData struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"advertId"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	TypeAdvert string `json:"typeAdvert"`
	// Photo is the filename of the photo.
	Photo string `json:"photo"`
	// TypeSale represents the sale type of the advertisement (Sale/Rent).
	TypeSale string `json:"typeSale"`
	// Address is the address of the advertisement.
	Address string `json:"adress"`
	// HouseProperties contains additional properties for houses.
	HouseProperties *HouseSquareProperties `json:"houseProperties,omitempty"`
	// FlatProperties contains additional properties for flats.
	FlatProperties *FlatSquareProperties `json:"flatProperties,omitempty"`
	// Price is the price of the advertisement.
	Price int `json:"price"`
	// DateCreation is the date when the advert was created.
	DateCreation time.Time `json:"dateCreation"`
}

// AdvertRectangleData represents the structure of the JSON data for Rectangle advert.
type AdvertRectangleData struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"advertId"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	TypeAdvert string `json:"typeAdvert"`
	// Photo is the filename of the photo.
	Photo string `json:"photo"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// TypeSale represents the sale type of the advertisement (Sale/Rent).
	TypeSale string `json:"typeSale"`
	// Address is the address of the advertisement.
	Address string `json:"adress"`
	// Complex represents residential complex information.
	// Complex map[string]interface{} `json:"complex"`
	// FlatProperties contains additional properties for flats.
	FlatProperties *FlatRectangleProperties `json:"flatProperties,omitempty"`
	// HouseProperties contains additional properties for flats.
	HouseProperties *HouseRectangleProperties `json:"houseProperties,omitempty"`
	// Price is the price of the advertisement.
	Price int `json:"price"`
	// DateCreation is the date when the advert was created.
	DateCreation time.Time `json:"dateCreation"`
}

// AdvertData represents the structure of the JSON data for advert.
type AdvertData struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"advertId"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	TypeAdvert string `json:"typeAdvert"`
	// TypeSale represents the sale type of the advertisement (Sale/Rent).
	TypeSale string `json:"typeSale"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// Price is the price of the advertisement.
	Price int64 `json:"price"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Address is the address of the advertisement.
	Address string `json:"adress"`
	// AddressPoint is the address of the advertisement.
	AddressPoint string `json:"adressPoint"`
	// Images contains filenames of photos for advert.
	Images []*ImageResp `json:"images"`
	// HouseProperties contains additional properties for house.
	HouseProperties *HouseProperties `json:"houseProperties,omitempty"`
	// FlatProperties contains additional properties for flat.
	FlatProperties *FlatProperties `json:"flatProperties,omitempty"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
	//ComplexProperties represents residential complex information.
	ComplexProperties *ComplexAdvertProperties `json:"complexProperties,omitempty"`
	// DateCreation is the date when the advert was created.
	DateCreation time.Time `json:"dateCreation"`
}

// AdvertDataWithImages represents the structure of the JSON data for advert.
//type AdvertUpdateData struct {

//}

// AdvertUpdateData represents the structure of the JSON data for update advert.
type AdvertUpdateData struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"-"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	TypeAdvert string `json:"typeAdvert"`
	// TypeSale represents the sale type of the advertisement (Sale/Rent).
	TypeSale string `json:"typeSale"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// Price is the price of the advertisement.
	Price int64 `json:"price"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Address is the address of the advertisement.
	Address string `json:"adress"`
	// AddressPoint is the address of the advertisement.
	AddressPoint string `json:"adressPoint"`
	// Properties contains additional properties for houses or flats.
	Properties map[string]interface{} `json:"properties"` //Одна структура properties для всех объявлений
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
}

type AdvertFilter struct {
	// MinPrice is the minimal price of the search advertisement.
	MinPrice int64 `json:"minPrice"`
	// MaxPrice is the maximum price of the search advertisement.
	MaxPrice int64 `json:"maxPrice"`
	// Offset is the offset search advertisement.
	Offset int `json:"offset"`
	// Page ...
	Page int `json:"page"`
	// RoomCount is the how many rooms need in advert
	RoomCount int `json:"roomCount"`
	// AdvertType represents the type of the search advertisement (House/Flat).
	AdvertType string `json:"advertType"`
	// DealType represents the deal type of the search advertisement (Sale/Rent).
	DealType string `json:"dealType"`
	// Address is the address of the search advertisement.
	Address string `json:"adress"`
}

type PageInfo struct {
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	CurrentPage   int `json:"currentPage"`
	PageSize      int `json:"pageSize"`
}

type AdvertDataPage struct {
	Adverts  []*AdvertRectangleData `json:"adverts"`
	PageInfo *PageInfo              `json:"pageInfo"`
}
