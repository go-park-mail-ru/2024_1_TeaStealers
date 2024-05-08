package models

import (
	"time"
)

// Flat represents a flat entity.
type Flat struct {
	// ID is the unique identifier for the flat.
	ID int64 `json:"id"`
	// BuildingID is the identifier of the building to which the flat belongs.
	BuildingID int64 `json:"buildingId"`
	// Floor is the floor of the flat.
	Floor int `json:"floor"`
	// CeilingHeight is the ceiling height of the flat.
	CeilingHeight float64 `json:"ceilingHeight"`
	// RoomCount is the how many rooms in flat
	RoomCount int `json:"roomCount"`
	// SquareGeneral is the general square of the flat.
	SquareGeneral float64 `json:"squareGeneral"`
	// SquareResidential is the residential square of the flat.
	SquareResidential float64 `json:"squareResidential"`
	// Apartment indicates if the flat is an apartment.
	Apartment bool `json:"apartment"`
	// DateCreation is the date when the flat was published.
	DateCreation time.Time `json:"-"`
	// IsDeleted is a flag indicating whether the flat is deleted.
	IsDeleted bool `json:"-"`
}

// FlatProperties represents a flat properties.
type FlatProperties struct {
	// Floor is the floor of the flat.
	Floor int `json:"floor"`
	// FloorGeneral is the General foor of the building.
	FloorGeneral int `json:"floorGeneral"`
	// CeilingHeight is the ceiling height of the flat.
	CeilingHeight float64 `json:"ceilingHeight"`
	// RoomCount is the how many rooms in flat
	RoomCount int `json:"roomCount"`
	// SquareGeneral is the general square of the flat.
	SquareGeneral float64 `json:"squareGeneral"`
	// SquareResidential is the residential square of the flat.
	SquareResidential float64 `json:"squareResidential"`
	// Apartment indicates if the flat is an apartment.
	Apartment bool `json:"apartment"`
}

// FlatSquareProperties represents a flat properties in square advert.
type FlatSquareProperties struct {
	// Floor is the floor of the flat.
	Floor int `json:"floor"`
	// FloorGeneral is the General foor of the building.
	FloorGeneral int `json:"floorGeneral"`
	// RoomCount is the how many rooms in flat
	RoomCount int `json:"roomCount"`
	// SquareGeneral is the general square of the flat.
	SquareGeneral float64 `json:"squareGeneral"`
}

// FlatRectangleProperties represents a flat properties in rectangle advert.
type FlatRectangleProperties struct {
	// Floor is the floor of the flat.
	Floor int `json:"floor"`
	// FloorGeneral is the General foor of the building.
	FloorGeneral int `json:"floorGeneral"`
	// RoomCount is the how many rooms in flat
	RoomCount int `json:"roomCount"`
	// SquareGeneral is the general square of the flat.
	SquareGeneral float64 `json:"squareGeneral"`
}
