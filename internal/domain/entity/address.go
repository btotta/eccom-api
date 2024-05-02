package entity

import (
	"gorm.io/gorm"
)

type Type string

const (
	INCOMPLETE Type = "INCOMPLETE"
	COMPLETE   Type = "COMPLETE"
)

type State struct {
	gorm.Model
	Name   string `json:"name" gorm:"index"`
	UF     string `json:"uf" gorm:"index,unique"`
	Status Status `json:"status" gorm:"not null"`
}

type City struct {
	gorm.Model
	Name    string `json:"name" gorm:"index"`
	StateID uint   `json:"state_id" gorm:"index"`
	State   State  `json:"state" gorm:"foreignKey:StateID"`
	Status  Status `json:"status" gorm:"not null"`
}

type Neighborhood struct {
	gorm.Model
	Name   string `json:"name" gorm:"index"`
	CityID uint   `json:"city_id"`
	City   *City  `json:"city" gorm:"foreignKey:CityID"`
	Status Status `json:"status" gorm:"not null"`
}

type Place struct {
	gorm.Model
	Name           string        `json:"name" gorm:"index"`
	PostalCode     string        `json:"postal_code" gorm:"not null;index"`
	Latitude       float64       `json:"latitude" gorm:"not null;type:decimal(10,8)"`
	Longitude      float64       `json:"longitude" gorm:"not null;type:decimal(11,8)"`
	NeighborhoodId uint          `json:"neighborhood_id" gorm:"not null;type:uuid;index"`
	Neighborhood   *Neighborhood `gorm:"foreignKey:NeighborhoodId"`
	Status         Status        `json:"status" gorm:"not null"`
	Type           Type          `json:"type" gorm:"not null;type:varchar(255);index"`
}
