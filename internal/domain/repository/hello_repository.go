package repository

import "gorm.io/gorm"

type HelloRepository interface {
	Health() error
}

type helloRepository struct {
	db *gorm.DB
}

func NewHelloRepository(db *gorm.DB) HelloRepository {
	return &helloRepository{
		db: db,
	}
}

func (r *helloRepository) Health() error {

	return r.db.Exec("SELECT 1").Error
}
