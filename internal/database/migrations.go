package database

import (
	"eccom-api/internal/domain/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {

	db.AutoMigrate(&entity.User{},
		&entity.State{},
		&entity.City{},
		&entity.Neighborhood{},
		&entity.Place{},
		&entity.UserAddress{},
	)

}
