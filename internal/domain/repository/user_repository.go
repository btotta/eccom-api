package repository

import (
	"eccom-api/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
	SaveUser(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {

	var user entity.User

	query := r.db.Where("email = ?", email)
	query = r.filterByNotDeleted(query)

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) SaveUser(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) filterByNotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("status != ?", entity.Deleted)
}
