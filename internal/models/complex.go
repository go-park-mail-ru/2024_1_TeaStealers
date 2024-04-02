package models

import (
	"time"

	"github.com/satori/uuid"
)

// ClassHouse represents the class of a house in a complex.
type ClassHouse string

const (
	// ClassHouseEconom represents an Econom class house.
	ClassHouseEconom ClassHouse = "Econom"
	// ClassHouseComfort represents a Comfort class house.
	ClassHouseComfort ClassHouse = "Comfort"
	// ClassHouseBusiness represents a Business class house.
	ClassHouseBusiness ClassHouse = "Business"
	// ClassHousePremium represents a Premium class house.
	ClassHousePremium ClassHouse = "Premium"
	// ClassHouseElite represents an Elite class house.
	ClassHouseElite ClassHouse = "Elite"
)

// Complex represents a complex entity.
type Complex struct {
	// ID is the unique identifier for the complex.
	ID uuid.UUID `json:"id"`
	// CompanyID is the identifier of the company that owns the complex.
	CompanyId uuid.UUID `json:"companyId"`
	// Name is the name of the complex.
	Name string `json:"name"`
	// Address is the address of the complex.
	Address string `json:"address"`
	// Photo is the filename of the complex's photo.
	Photo string `json:"photo"`
	// Description is the description of the complex.
	Description string `json:"description"`
	// DateBeginBuild is the start date of the construction of the complex.
	DateBeginBuild time.Time `json:"dateBeginBuild"`
	// DateEndBuild is the end date of the construction of the complex.
	DateEndBuild time.Time `json:"dateEndBuild"`
	// WithoutFinishingOption indicates if the complex has an option without finishing.
	WithoutFinishingOption bool `json:"withoutFinishingOption"`
	// FinishingOption indicates if the complex has a finishing option.
	FinishingOption bool `json:"finishingOption"`
	// PreFinishingOption indicates if the complex has a pre-finishing option.
	PreFinishingOption bool `json:"preFinishingOption"`
	// ClassHousing is the class of housing in the complex.
	ClassHousing ClassHouse `json:"classHousing"`
	// Parking indicates if the complex has parking.
	Parking bool `json:"parking"`
	// Security indicates if the complex has security.
	Security bool `json:"security"`
	// DateCreation is the date when the complex was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the complex is deleted.
	IsDeleted bool `json:"-"`
}

// ComplexCreate represents a data for creation complex.
type ComplexCreateData struct {
	// CompanyID is the identifier of company for the complex.
	CompanyId uuid.UUID `json:"companyId"`
	// Name is the name of the complex.
	Name string `json:"name"`
	// Address is the address of the complex.
	Address string `json:"address"`
	// Description is the description of the complex.
	Description string `json:"description"`
	// DateBeginBuild is the start date of the construction of the complex.
	DateBeginBuild time.Time `json:"dateBeginBuild"`
	// DateEndBuild is the end date of the construction of the complex.
	DateEndBuild time.Time `json:"dateEndBuild"`
	// WithoutFinishingOption indicates if the complex has an option without finishing.
	WithoutFinishingOption bool `json:"withoutFinishingOption"`
	// FinishingOption indicates if the complex has a finishing option.
	FinishingOption bool `json:"finishingOption"`
	// PreFinishingOption indicates if the complex has a pre-finishing option.
	PreFinishingOption bool `json:"preFinishingOption"`
	// ClassHousing is the class of housing in the complex.
	ClassHousing ClassHouse `json:"classHousing"`
	// Parking indicates if the complex has parking.
	Parking bool `json:"parking"`
	// Security indicates if the complex has security.
	Security bool `json:"security"`
}

// ComplexData represents a complex information.
type ComplexData struct {
	// ID is the unique identifier for the complex.
	ID uuid.UUID `json:"id"`
	// CompanyID is the identifier of the company that owns the complex.
	CompanyId uuid.UUID `json:"companyId"`
	// Name is the name of the complex.
	Name string `json:"name"`
	// Address is the address of the complex.
	Address string `json:"address"`
	// Photo is the filename of the complex's photo.
	Photo string `json:"photo"`
	// Description is the description of the complex.
	Description string `json:"description"`
	// DateBeginBuild is the start date of the construction of the complex.
	DateBeginBuild time.Time `json:"dateBeginBuild"`
	// DateEndBuild is the end date of the construction of the complex.
	DateEndBuild time.Time `json:"dateEndBuild"`
	// WithoutFinishingOption indicates if the complex has an option without finishing.
	WithoutFinishingOption bool `json:"withoutFinishingOption"`
	// FinishingOption indicates if the complex has a finishing option.
	FinishingOption bool `json:"finishingOption"`
	// PreFinishingOption indicates if the complex has a pre-finishing option.
	PreFinishingOption bool `json:"preFinishingOption"`
	// ClassHousing is the class of housing in the complex.
	ClassHousing ClassHouse `json:"classHousing"`
	// Parking indicates if the complex has parking.
	Parking bool `json:"parking"`
	// Security indicates if the complex has security.
	Security bool `json:"security"`
}

// ComplexAdvertProperties represents a complex properties for advert.
type ComplexAdvertProperties struct {
	// ComplexId is the unique identifier for the complex.
	ComplexId string `json:"complexId"`
	// Name is the name of the complex.
	NameComplex string `json:"nameComplex"`
	// Photo is the filename of the company's photo.
	PhotoCompany string `json:"photoCompany"`
	// Name is the name of the company.
	NameCompany string `json:"nameCompany"`
}
