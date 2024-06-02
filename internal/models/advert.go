package models

import (
	"html"
	"time"
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
	ID int64 `json:"id"`
	// UserID is the identifier of the user who created the advert.
	UserID int64 `json:"userId"`
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

	Likes int `json:"likes"`
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
	UserID int64 `json:"userId"`
	// AdvertTypePlacement is the sale type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
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
	// RoomCount is how many rooms in flat
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
	Address AddressData `json:"address"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
	// DateCreation is the date when the building was published.
}

type AddressData struct {
	// Province is the name of the province.
	Province string `json:"province"`
	// Town is the name of the streer.
	Town string `json:"town"`
	// Street is the name of the street.
	Street string `json:"street"`
	// House is the name of the house.
	House string `json:"house"`
	// Metro is the name of the metro.
	Metro string `json:"metro"`
	// AddressPoint is the geographical point of the building's address.
	AddressPoint string `json:"addressPoint"`
}

func (advFlCr *AdvertFlatCreateData) Sanitize() {
	advFlCr.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(advFlCr.AdvertTypeSale)))
	advFlCr.Title = html.EscapeString(advFlCr.Title)
	advFlCr.Description = html.EscapeString(advFlCr.Description)
	advFlCr.Phone = html.EscapeString(advFlCr.Phone)
}

// ComplexAdvertFlatCreateData represents a data for creation advertisement.
type ComplexAdvertFlatCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID int64 `json:"userId"`
	// BuildingID is the identifier of the building to which the flat belongs.
	BuildingID int64 `json:"buildingId"`
	// AdvertTypePlacement is the sale type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
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
}

func (complAdvFlCrDat *ComplexAdvertFlatCreateData) Sanitize() {
	complAdvFlCrDat.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(complAdvFlCrDat.AdvertTypeSale)))
	complAdvFlCrDat.Title = html.EscapeString(complAdvFlCrDat.Title)
	complAdvFlCrDat.Description = html.EscapeString(complAdvFlCrDat.Description)
	complAdvFlCrDat.Phone = html.EscapeString(complAdvFlCrDat.Phone)
}

// AdvertHouseCreateData represents a data for creation advertisement.
type AdvertHouseCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID int64 `json:"userId"`
	// AdvertTypePlacement is the sale type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
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
	Address AddressData `json:"address"`
	// YearCreation is the year when the building was created.
	YearCreation int `json:"yearCreation"`
}

func (advHousCrDat *AdvertHouseCreateData) Sanitize() {
	advHousCrDat.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(advHousCrDat.AdvertTypeSale)))
	advHousCrDat.Title = html.EscapeString(advHousCrDat.Title)
	advHousCrDat.Description = html.EscapeString(advHousCrDat.Description)
	advHousCrDat.Phone = html.EscapeString(advHousCrDat.Phone)
}

// ComplexAdvertHouseCreateData represents a data for creation advertisement.
type ComplexAdvertHouseCreateData struct {
	// UserID is the identifier of the user who created the advert.
	UserID int64 `json:"userId"`
	// BuildingID is the identifier of the building to which the house belongs.
	BuildingID int64 `json:"buildingId"`
	// AdvertTypePlacement is the sale type of the advert (Sale/Rent).
	AdvertTypeSale TypePlacementAdvert `json:"advertTypeSale"`
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
}

func (complAdvHousCrDat *ComplexAdvertHouseCreateData) Sanitize() {
	complAdvHousCrDat.AdvertTypeSale = TypePlacementAdvert(html.EscapeString(string(complAdvHousCrDat.AdvertTypeSale)))
	complAdvHousCrDat.Title = html.EscapeString(complAdvHousCrDat.Title)
	complAdvHousCrDat.Description = html.EscapeString(complAdvHousCrDat.Description)
	complAdvHousCrDat.Phone = html.EscapeString(complAdvHousCrDat.Phone)
}

// AdvertSquareData represents the structure of the JSON data for square advert.
type AdvertSquareData struct {
	// ID is the unique identifier for the advert.
	ID int64 `json:"advertId"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	TypeAdvert string `json:"typeAdvert"`
	// Photo is the filename of the photo.
	Photo string `json:"photo"`
	// TypeSale represents the sale type of the advertisement (Sale/Rent).
	TypeSale string `json:"typeSale"`
	// Address is the address of the advertisement.
	Address string `json:"address"`
	// Metro is the metro of the advertisement.
	Metro string `json:"metro"`
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
	ID int64 `json:"advertId"`
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
	// AddressPoint is the address of the advertisement.
	AddressPoint string `json:"adressPoint"`
	// Rating is the metro of the advertisement.
	Rating string `json:"rating"`
	// IsLiked indicates whether the advert is liked.
	IsLiked bool `json:"isLiked"`
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
	// Rating is the rating of the advert.
	// Rating string `json:"rating"`
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
	ID int64 `json:"advertId"`
	// TypeAdvert represents the type of the advertisement (House/Flat).
	AdvertType string `json:"advertType"`
	// TypeSale represents the sale type of the advertisement (Sale/Rent).
	TypeSale string `json:"typeSale"`
	// Title is the title of the advert.
	Title string `json:"title"`
	// Description is the description of the advert.
	Description string `json:"description"`
	// CountViews is the count of views the advertisement.
	CountViews int64 `json:"countViews"`
	// CountLikes is the count of likes the advertisement.
	CountLikes int64 `json:"countLikes"`
	// Price is the price of the advertisement.
	Price int64 `json:"price"`
	// Phone is the phone number associated with the advert.
	Phone string `json:"phone"`
	// IsLiked indicates whether the advert is liked.
	IsLiked bool `json:"isLiked"`
	// IsAgent indicates whether the advert is posted by an agent.
	IsAgent bool `json:"isAgent"`
	// Metro is the metro of the advertisement.
	Metro string `json:"metro"`
	// Address is the address of the advertisement.
	Address string `json:"adress"`
	// AddressPoint is the address of the advertisement.
	AddressPoint string `json:"adressPoint"`
	// PriceChange contains changes of price for advert.
	PriceChange []*PriceChangeData `json:"priceHistory"`
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
	ID int64 `json:"-"`
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
	// Address is the address of the building.
	Address AddressData `json:"address"`
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
