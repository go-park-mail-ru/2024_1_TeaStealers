package models

import (
	"html"
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

func (adv *Advert) Sanitize() {
	adv.Title = html.EscapeString(adv.Title)
	adv.Description = html.EscapeString(adv.Description)
	adv.Phone = html.EscapeString(adv.Phone)
	adv.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(adv.AdvertTypeSale)))
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

func (advFlCr *AdvertFlatCreateData) Sanitize() {
	advFlCr.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(advFlCr.AdvertTypeSale)))
	advFlCr.AdvertTypePlacement = AdvertTypeAdvert(html.EscapeString(string(advFlCr.AdvertTypePlacement)))
	advFlCr.Title = html.EscapeString(advFlCr.Title)
	advFlCr.Description = html.EscapeString(advFlCr.Description)
	advFlCr.Phone = html.EscapeString(advFlCr.Phone)
	advFlCr.Address = html.EscapeString(advFlCr.Address)
	advFlCr.AddressPoint = html.EscapeString(advFlCr.AddressPoint)
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

func (complAdvFlCrDat *ComplexAdvertFlatCreateData) Sanitize() {
	complAdvFlCrDat.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(complAdvFlCrDat.AdvertTypeSale)))
	complAdvFlCrDat.AdvertTypePlacement = AdvertTypeAdvert(html.EscapeString(string(complAdvFlCrDat.AdvertTypePlacement)))
	complAdvFlCrDat.Title = html.EscapeString(complAdvFlCrDat.Title)
	complAdvFlCrDat.Description = html.EscapeString(complAdvFlCrDat.Description)
	complAdvFlCrDat.Phone = html.EscapeString(complAdvFlCrDat.Phone)
	complAdvFlCrDat.Address = html.EscapeString(complAdvFlCrDat.Address)
	complAdvFlCrDat.AddressPoint = html.EscapeString(complAdvFlCrDat.AddressPoint)
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
	CeilingHeight float64 `json:"ceilingHeight"`
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

func (advHousCrDat *AdvertHouseCreateData) Sanitize() {
	advHousCrDat.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(advHousCrDat.AdvertTypeSale)))
	advHousCrDat.AdvertTypePlacement = AdvertTypeAdvert(html.EscapeString(string(advHousCrDat.AdvertTypePlacement)))
	advHousCrDat.Title = html.EscapeString(advHousCrDat.Title)
	advHousCrDat.Description = html.EscapeString(advHousCrDat.Description)
	advHousCrDat.Phone = html.EscapeString(advHousCrDat.Phone)
	advHousCrDat.Address = html.EscapeString(advHousCrDat.Address)
	advHousCrDat.AddressPoint = html.EscapeString(advHousCrDat.AddressPoint)
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
	CeilingHeight float64 `json:"ceilingHeight"`
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

func (complAdvHousCrDat *ComplexAdvertHouseCreateData) Sanitize() {
	complAdvHousCrDat.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(complAdvHousCrDat.AdvertTypeSale)))
	complAdvHousCrDat.AdvertTypePlacement = AdvertTypeAdvert(html.EscapeString(string(complAdvHousCrDat.AdvertTypePlacement)))
	complAdvHousCrDat.Title = html.EscapeString(complAdvHousCrDat.Title)
	complAdvHousCrDat.Description = html.EscapeString(complAdvHousCrDat.Description)
	complAdvHousCrDat.Phone = html.EscapeString(complAdvHousCrDat.Phone)
	complAdvHousCrDat.Address = html.EscapeString(complAdvHousCrDat.Address)
	complAdvHousCrDat.AddressPoint = html.EscapeString(complAdvHousCrDat.AddressPoint)
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

func (advSqDat *AdvertSquareData) Sanitize() {
	advSqDat.TypeAdvert = html.EscapeString(advSqDat.TypeAdvert)
	advSqDat.Photo = html.EscapeString(advSqDat.Photo)
	advSqDat.TypeSale = html.EscapeString(advSqDat.TypeSale)
	advSqDat.Address = html.EscapeString(advSqDat.Address)
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

func (advRectDat *AdvertRectangleData) Sanitize() {
	advRectDat.TypeAdvert = html.EscapeString(advRectDat.TypeAdvert)
	advRectDat.Photo = html.EscapeString(advRectDat.Photo)
	advRectDat.TypeSale = html.EscapeString(advRectDat.TypeSale)
	advRectDat.Address = html.EscapeString(advRectDat.Address)
}

// AdvertData represents the structure of the JSON data for advert.
type AdvertData struct {
	// ID is the unique identifier for the advert.
	ID uuid.UUID `json:"advertId"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	AdvertType string `json:"advertType"`
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
	// ComplexProperties represents residential complex information.
	ComplexProperties *ComplexAdvertProperties `json:"complexProperties,omitempty"`
	// DateCreation is the date when the advert was created.
	DateCreation time.Time `json:"dateCreation"`
}

func (advDat *AdvertData) Sanitize() {
	advDat.AdvertType = html.EscapeString(advDat.AdvertType)
	advDat.TypeSale = html.EscapeString(advDat.TypeSale)
	advDat.Title = html.EscapeString(advDat.Title)
	advDat.Description = html.EscapeString(advDat.Description)
	advDat.Address = html.EscapeString(advDat.Address)
	advDat.AddressPoint = html.EscapeString(advDat.AddressPoint)
	advDat.HouseProperties.Sanitize()
	advDat.Material = MaterialBuilding(html.EscapeString(string(advDat.Material)))
	advDat.ComplexProperties.Sanitize()
}

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
	Price float64 `json:"price"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Address is the address of the advertisement.
	Address string `json:"adress"`
	// AddressPoint is the address of the advertisement.
	AddressPoint string `json:"adressPoint"`
	// HouseProperties contains additional properties for house.
	HouseProperties *HouseProperties `json:"houseProperties,omitempty"`
	// FlatProperties contains additional properties for flat.
	FlatProperties *FlatProperties `json:"flatProperties,omitempty"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// Material is the material of the building.
	Material MaterialBuilding `json:"material"`
}

func (advDat *AdvertUpdateData) Sanitize() {
	advDat.TypeAdvert = html.EscapeString(advDat.TypeAdvert)
	advDat.TypeSale = html.EscapeString(advDat.TypeSale)
	advDat.Title = html.EscapeString(advDat.Title)
	advDat.Description = html.EscapeString(advDat.Description)
	advDat.Address = html.EscapeString(advDat.Address)
	advDat.AddressPoint = html.EscapeString(advDat.AddressPoint)
	advDat.HouseProperties.Sanitize()
	advDat.Material = MaterialBuilding(html.EscapeString(string(advDat.Material)))
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

func (advFilter *AdvertFilter) Sanitize() {
	advFilter.AdvertType = html.EscapeString(advFilter.AdvertType)
	advFilter.DealType = html.EscapeString(advFilter.DealType)
	advFilter.Address = html.EscapeString(advFilter.Address)
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

func (advDatPage *AdvertDataPage) Sanitize() {
	for _, v := range advDatPage.Adverts {
		v.Sanitize()
	}
}
