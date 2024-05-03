package entity

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	UserID         uint         `json:"user_id"`
	User           User         `json:"user" gorm:"foreignKey:UserID"`
	StateID        uint         `json:"state_id"`
	State          State        `json:"state" gorm:"foreignKey:StateID"`
	CityID         uint         `json:"city_id"`
	City           City         `json:"city" gorm:"foreignKey:CityID"`
	NeighborhoodID uint         `json:"neighborhood_id"`
	Neighborhood   Neighborhood `json:"neighborhood" gorm:"foreignKey:NeighborhoodID"`
	PlaceID        uint         `json:"place_id"`
	Place          Place        `json:"place" gorm:"foreignKey:PlaceID"`
	Status         Status       `json:"status"`
	Numero         string       `json:"numero"`
	Complemento    string       `json:"complemento"`
}
