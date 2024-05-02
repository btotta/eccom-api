package repository

import (
	"eccom-api/internal/domain/entity"
	"strings"

	"gorm.io/gorm"
)

type AddressRepository interface {
	SaveState(state *entity.State) (*entity.State, error)
	FindStateByID(stateID uint) (*entity.State, error)
	FindStateByUF(uf string) (*entity.State, error)
	FindAllStates(page, size int, orderBy []string, uf, name string) (*PageResponse[entity.State], error)

	SaveCity(city *entity.City) (*entity.City, error)
	FindCityByID(cityID uint) (*entity.City, error)
	FindCityByName(name string) (*entity.City, error)
	FindAllCities(page, size int, orderBy []string, stateID uint, name string) (*PageResponse[entity.City], error)

	SaveNeighborhood(neighborhood *entity.Neighborhood) (*entity.Neighborhood, error)
	FindNeighborhoodByID(neighborhoodID uint) (*entity.Neighborhood, error)
	FindNeighborhoodByName(name string) (*entity.Neighborhood, error)
	FindAllNeighborhoods(page, size int, orderBy []string, cityID uint, name string) (*PageResponse[entity.Neighborhood], error)

	SavePlace(place *entity.Place) (*entity.Place, error)
	FindPlaceByID(placeID uint) (*entity.Place, error)
	FindPlaceByCEP(cep string) (*entity.Place, error)
	FindAllPlaces(page, size int, orderBy []string, neighborhoodID uint, name string) (*PageResponse[entity.Place], error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{
		db: db,
	}
}

// State

func (r *addressRepository) SaveState(state *entity.State) (*entity.State, error) {
	if err := r.db.Create(state).Error; err != nil {
		return nil, err
	}
	return state, nil
}

func (r *addressRepository) FindStateByID(stateID uint) (*entity.State, error) {

	var state entity.State

	if err := r.db.First(&state, stateID).Error; err != nil {
		return nil, err
	}

	return &state, nil
}

func (r *addressRepository) FindStateByUF(uf string) (*entity.State, error) {

	var state entity.State

	query := r.db.Model(&entity.State{})
	query = r.filterByStateUf(query, uf)
	query = r.filterByStatusNotDeleted(query)

	if err := query.First(&state).Error; err != nil {
		return nil, err
	}

	return &state, nil
}

func (r *addressRepository) FindAllStates(page, size int, orderBy []string, uf, name string) (*PageResponse[entity.State], error) {

	query := r.db.Model(&entity.State{})
	query = r.filterByStatusNotDeleted(query)
	query = r.filterByStateUf(query, uf)
	query = r.filterByStateName(query, name)

	resp, err := Paginate[entity.State](query, page, size, orderBy)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (r *addressRepository) filterByStateUf(query *gorm.DB, uf string) *gorm.DB {
	if uf != "" {
		query = query.Where("uf = ?", strings.ToUpper(uf))
	}
	return query

}

func (r *addressRepository) filterByStateName(query *gorm.DB, name string) *gorm.DB {
	if name != "" {
		query = query.Where("lower(name) LIKE ?", "%"+strings.TrimSpace(strings.ToLower(name))+"%")
	}

	return query
}

// City

func (r *addressRepository) SaveCity(city *entity.City) (*entity.City, error) {
	if err := r.db.Create(city).Error; err != nil {
		return nil, err
	}
	return city, nil
}

func (r *addressRepository) FindCityByID(cityID uint) (*entity.City, error) {
	var city entity.City

	if err := r.db.First(&city, cityID).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (r *addressRepository) FindCityByName(name string) (*entity.City, error) {
	var city entity.City

	query := r.db.Where("lower(name) = ?", strings.TrimSpace(strings.ToLower(name)))
	query = r.filterByStatusNotDeleted(query)

	if err := query.First(&city).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (r *addressRepository) FindAllCities(page, size int, orderBy []string, stateID uint, name string) (*PageResponse[entity.City], error) {

	query := r.db.Model(&entity.City{})
	query = r.filterByStatusNotDeleted(query)
	query = r.filterByCityStateID(query, stateID)
	query = r.filterByCityName(query, name)

	resp, err := Paginate[entity.City](query, page, size, orderBy)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *addressRepository) filterByCityStateID(query *gorm.DB, stateID uint) *gorm.DB {
	if stateID != 0 {
		query = query.Where("state_id = ?", stateID)
	}

	return query
}

func (r *addressRepository) filterByCityName(query *gorm.DB, name string) *gorm.DB {
	if name != "" {
		query = query.Where("lower(name) LIKE ?", "%"+strings.TrimSpace(strings.ToLower(name))+"%")
	}

	return query
}

// Neighborhood

func (r *addressRepository) SaveNeighborhood(neighborhood *entity.Neighborhood) (*entity.Neighborhood, error) {
	if err := r.db.Create(neighborhood).Error; err != nil {
		return nil, err
	}
	return neighborhood, nil
}

func (r *addressRepository) FindNeighborhoodByID(neighborhoodID uint) (*entity.Neighborhood, error) {
	var neighborhood entity.Neighborhood

	if err := r.db.First(&neighborhood, neighborhoodID).Error; err != nil {
		return nil, err
	}

	return &neighborhood, nil
}

func (r *addressRepository) FindNeighborhoodByName(name string) (*entity.Neighborhood, error) {
	var neighborhood entity.Neighborhood

	query := r.db.Where("lower(name) = ?", strings.TrimSpace(strings.ToLower(name)))
	query = r.filterByStatusNotDeleted(query)

	if err := query.First(neighborhood).Error; err != nil {
		return nil, err
	}

	return &neighborhood, nil
}

func (r *addressRepository) FindAllNeighborhoods(page, size int, orderBy []string, cityID uint, name string) (*PageResponse[entity.Neighborhood], error) {

	query := r.db.Model(&entity.Neighborhood{})
	query = r.filterByStatusNotDeleted(query)
	query = r.filterByNeighborhoodCityID(query, cityID)
	query = r.filterByNeighborhoodName(query, name)

	resp, err := Paginate[entity.Neighborhood](query, page, size, orderBy)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *addressRepository) filterByNeighborhoodCityID(query *gorm.DB, cityID uint) *gorm.DB {
	if cityID != 0 {
		query = query.Where("city_id = ?", cityID)
	}

	return query
}

func (r *addressRepository) filterByNeighborhoodName(query *gorm.DB, name string) *gorm.DB {
	if name != "" {
		query = query.Where("lower(name) LIKE ?", "%"+strings.TrimSpace(strings.ToLower(name))+"%")
	}

	return query
}

// Place

func (r *addressRepository) SavePlace(place *entity.Place) (*entity.Place, error) {
	if err := r.db.Create(place).Error; err != nil {
		return nil, err
	}
	return place, nil
}

func (r *addressRepository) FindPlaceByID(placeID uint) (*entity.Place, error) {
	var place entity.Place

	if err := r.db.First(&place, placeID).Error; err != nil {
		return nil, err
	}

	return &place, nil
}

func (r *addressRepository) FindPlaceByCEP(cep string) (*entity.Place, error) {
	var place entity.Place

	query := r.db.Where("cep = ?", cep)
	query = r.filterByStatusNotDeleted(query)

	if err := query.First(place).Error; err != nil {
		return nil, err
	}

	return &place, nil
}

func (r *addressRepository) FindAllPlaces(page, size int, orderBy []string, neighborhoodID uint, name string) (*PageResponse[entity.Place], error) {

	query := r.db.Model(&entity.Place{})
	query = r.filterByStatusNotDeleted(query)
	query = r.filterByPlaceNeighborhoodID(query, neighborhoodID)
	query = r.filterByPlaceName(query, name)

	resp, err := Paginate[entity.Place](query, page, size, orderBy)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *addressRepository) filterByPlaceNeighborhoodID(query *gorm.DB, neighborhoodID uint) *gorm.DB {
	if neighborhoodID != 0 {
		query = query.Where("neighborhood_id = ?", neighborhoodID)
	}

	return query
}

func (r *addressRepository) filterByPlaceName(query *gorm.DB, name string) *gorm.DB {
	if name != "" {
		query = query.Where("lower(name) LIKE ?", "%"+strings.TrimSpace(strings.ToLower(name))+"%")
	}

	return query
}

// Helpers

func (r *addressRepository) filterByStatusNotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("status != ?", entity.Deleted)
}
