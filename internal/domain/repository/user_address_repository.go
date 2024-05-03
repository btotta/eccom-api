package repository

import (
	"eccom-api/internal/domain/entity"

	"gorm.io/gorm"
)

type UserAddressRepository interface {
	CreateUserAddress(userAddress *entity.UserAddress) (*entity.UserAddress, error)
	FindUserAddressByIDAndUserID(userAddressID, userID uint) (*entity.UserAddress, error)
	SaveUserAddress(userAddress *entity.UserAddress) (*entity.UserAddress, error)
	GetUserAddresses(user uint, page, limit int, sort []string) (*PageResponse[entity.UserAddress], error)
}

type userAddressRepository struct {
	db *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) UserAddressRepository {
	return &userAddressRepository{
		db: db,
	}
}

func (r *userAddressRepository) CreateUserAddress(userAddress *entity.UserAddress) (*entity.UserAddress, error) {
	if err := r.db.Create(userAddress).Error; err != nil {
		return nil, err
	}
	return userAddress, nil
}

func (r *userAddressRepository) FindUserAddressByIDAndUserID(userAddressID, userID uint) (*entity.UserAddress, error) {

	var userAddress entity.UserAddress

	if err := r.db.Where("id = ? AND user_id = ?", userAddressID, userID).First(&userAddress).Error; err != nil {
		return nil, err
	}

	return &userAddress, nil
}

func (r *userAddressRepository) SaveUserAddress(userAddress *entity.UserAddress) (*entity.UserAddress, error) {
	if err := r.db.Save(userAddress).Error; err != nil {
		return nil, err
	}
	return userAddress, nil
}

func (r *userAddressRepository) GetUserAddresses(user uint, page, limit int, sort []string) (*PageResponse[entity.UserAddress], error) {

	query := r.db.Model(&entity.UserAddress{})
	query = query.Where("user_id = ?", user)
	query = r.filterByNotDeleted(query)

	resp, err := Paginate[entity.UserAddress](query, page, limit, sort)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *userAddressRepository) filterByNotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("status != ?", entity.Deleted)
}
