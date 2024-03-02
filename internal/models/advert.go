package models

import (
	"time"

	"github.com/satori/uuid"
)

// Advert represents advert information
type Advert struct {
	// ID uniquely identifies the advert.
	ID uuid.UUID `json:"id"`
	// UserId uniquely identifies who owns the advert.
	UserId uuid.UUID `json:"userid"`
	// Phone is the phone of the owner advert.
	Phone int `json:"phone"`
	// Description is the description of the advert.
	Descpription string `json:"description"`
	// BuildingId is the id of the building to which the advert belongs.
	BuildingId uuid.UUID `json:"buildingid"`
	// CompanyId is the id of the company to which the advert belongs.
	CompanyId uuid.UUID `json:"companyid"`
	// Price is the price of the object in advert.
	Price float64 `json:"price"`
	// Location is the location of the object in advert.
	Location string `json:"location"`
	// DataCreation is the time of adding a record to the database.
	DataCreation time.Time `json:"datacreation"`
	// isDeleted means is the advert deleted?.
	IsDeleted bool `json:"isdeleted"`
}

// AdvertCreateData represents building information for userid, phone, description, buildingid, companyid, price and location
type AdvertCreateData struct {
	// UserId uniquely identifies who owns the advert.
	UserId uuid.UUID `json:"userid"`
	// Phone is the phone of the owner advert.
	Phone int `json:"phone"`
	// Description is the description of the advert.
	Descpription string `json:"description"`
	// BuildingId is the id of the building to which the advert belongs.
	BuildingId uuid.UUID `json:"buildingid"`
	// CompanyId is the id of the company to which the advert belongs.
	CompanyId uuid.UUID `json:"companyid"`
	// Price is the price of the object in advert.
	Price float64 `json:"price"`
	// Location is the location of the object in advert.
	Location string `json:"location"`
}
