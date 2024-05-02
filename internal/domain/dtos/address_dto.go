package dtos

import (
	"eccom-api/internal/domain/entity"
	"eccom-api/internal/domain/repository"
	"strings"
	"time"
)

// State

type CreateStateDTO struct {
	Name string `json:"name" binding:"required"`
	UF   string `json:"uf" binding:"required"`
}

func (dto *CreateStateDTO) ToState() *entity.State {
	return &entity.State{
		Name:   strings.TrimSpace(dto.Name),
		UF:     strings.TrimSpace(strings.ToUpper(dto.UF)),
		Status: entity.Active,
	}
}

type StateResponseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	UF        string    `json:"uf"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewStateResponseDTO(state *entity.State) *StateResponseDTO {
	return &StateResponseDTO{
		ID:        state.ID,
		Name:      state.Name,
		UF:        state.UF,
		CreatedAt: state.CreatedAt,
		UpdatedAt: state.UpdatedAt,
	}
}

type PageStateResponseDTO struct {
	Content          []StateResponseDTO      `json:"content"`
	Empty            bool                    `json:"empty"`
	First            bool                    `json:"first"`
	Last             bool                    `json:"last"`
	Number           int                     `json:"number"`
	NumberOfElements int                     `json:"numberOfElements"`
	Pageable         repository.PageableInfo `json:"pageable"`
	Size             int                     `json:"size"`
	Sort             repository.SortInfo     `json:"sort"`
	TotalElements    int64                   `json:"totalElements"`
	TotalPages       int                     `json:"totalPages"`
}

func NewPageStateResponseDTO(page *repository.PageResponse[entity.State]) *PageStateResponseDTO {

	content := make([]StateResponseDTO, 0)
	for _, state := range page.Content {
		content = append(content, *NewStateResponseDTO(&state))
	}

	return &PageStateResponseDTO{
		Content:          content,
		Empty:            page.Empty,
		First:            page.First,
		Last:             page.Last,
		Number:           page.Number,
		NumberOfElements: page.NumberOfElements,
		Pageable:         page.Pageable,
		Size:             page.Size,
		Sort:             page.Sort,
		TotalElements:    page.TotalElements,
		TotalPages:       page.TotalPages,
	}
}

// City

type CreateCityDTO struct {
	Name    string `json:"name" binding:"required"`
	StateID uint   `json:"state_id" binding:"required"`
}

func (dto *CreateCityDTO) ToCity() *entity.City {
	return &entity.City{
		Name:   strings.TrimSpace(dto.Name),
		Status: entity.Active,
	}
}

type CityResponseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	StateID   uint      `json:"state_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCityResponseDTO(city *entity.City) *CityResponseDTO {
	return &CityResponseDTO{
		ID:        city.ID,
		Name:      city.Name,
		StateID:   city.StateID,
		CreatedAt: city.CreatedAt,
		UpdatedAt: city.UpdatedAt,
	}
}

type PageCityResponseDTO struct {
	Content          []CityResponseDTO       `json:"content"`
	Empty            bool                    `json:"empty"`
	First            bool                    `json:"first"`
	Last             bool                    `json:"last"`
	Number           int                     `json:"number"`
	NumberOfElements int                     `json:"numberOfElements"`
	Pageable         repository.PageableInfo `json:"pageable"`
	Size             int                     `json:"size"`
	Sort             repository.SortInfo     `json:"sort"`
	TotalElements    int64                   `json:"totalElements"`
	TotalPages       int                     `json:"totalPages"`
}

func NewPageCityResponseDTO(page *repository.PageResponse[entity.City]) *PageCityResponseDTO {

	content := make([]CityResponseDTO, 0)
	for _, city := range page.Content {
		content = append(content, *NewCityResponseDTO(&city))
	}

	return &PageCityResponseDTO{
		Content:          content,
		Empty:            page.Empty,
		First:            page.First,
		Last:             page.Last,
		Number:           page.Number,
		NumberOfElements: page.NumberOfElements,
		Pageable:         page.Pageable,
		Size:             page.Size,
		Sort:             page.Sort,
		TotalElements:    page.TotalElements,
		TotalPages:       page.TotalPages,
	}
}

// Neighborhood

type CreateNeighborhoodDTO struct {
	Name   string `json:"name" binding:"required"`
	CityID uint   `json:"city_id" binding:"required"`
}

func (dto *CreateNeighborhoodDTO) ToNeighborhood() *entity.Neighborhood {
	return &entity.Neighborhood{
		Name:   strings.TrimSpace(dto.Name),
		Status: entity.Active,
	}
}

type NeighborhoodResponseDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CityID    uint      `json:"city_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNeighborhoodResponseDTO(neighborhood *entity.Neighborhood) *NeighborhoodResponseDTO {
	return &NeighborhoodResponseDTO{
		ID:        neighborhood.ID,
		Name:      neighborhood.Name,
		CityID:    neighborhood.CityID,
		CreatedAt: neighborhood.CreatedAt,
		UpdatedAt: neighborhood.UpdatedAt,
	}
}

type PageNeighborhoodResponseDTO struct {
	Content          []NeighborhoodResponseDTO `json:"content"`
	Empty            bool                      `json:"empty"`
	First            bool                      `json:"first"`
	Last             bool                      `json:"last"`
	Number           int                       `json:"number"`
	NumberOfElements int                       `json:"numberOfElements"`
	Pageable         repository.PageableInfo   `json:"pageable"`
	Size             int                       `json:"size"`
	Sort             repository.SortInfo       `json:"sort"`
	TotalElements    int64                     `json:"totalElements"`
	TotalPages       int                       `json:"totalPages"`
}

func NewPageNeighborhoodResponseDTO(page *repository.PageResponse[entity.Neighborhood]) *PageNeighborhoodResponseDTO {

	content := make([]NeighborhoodResponseDTO, 0)
	for _, neighborhood := range page.Content {
		content = append(content, *NewNeighborhoodResponseDTO(&neighborhood))
	}

	return &PageNeighborhoodResponseDTO{
		Content:          content,
		Empty:            page.Empty,
		First:            page.First,
		Last:             page.Last,
		Number:           page.Number,
		NumberOfElements: page.NumberOfElements,
		Pageable:         page.Pageable,
		Size:             page.Size,
		Sort:             page.Sort,
		TotalElements:    page.TotalElements,
		TotalPages:       page.TotalPages,
	}
}

// Place

type CreatePlaceDTO struct {
	Name           string   `json:"name" binding:"required"`
	PostalCode     string   `json:"postal_code" binding:"required"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
	NeighborhoodID uint     `json:"neighborhood_id" binding:"required"`
}

func (dto *CreatePlaceDTO) ToPlace() *entity.Place {
	return &entity.Place{
		Name:       strings.TrimSpace(dto.Name),
		PostalCode: dto.PostalCode,
		Latitude:   *dto.Latitude,
		Longitude:  *dto.Longitude,
		Status:     entity.Active,
	}
}

type PlaceResponseDTO struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	PostalCode     string    `json:"postal_code"`
	Latitude       *float64  `json:"latitude"`
	Longitude      *float64  `json:"longitude"`
	NeighborhoodID uint      `json:"neighborhood_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewPlaceResponseDTO(place *entity.Place) *PlaceResponseDTO {
	return &PlaceResponseDTO{
		ID:             place.ID,
		Name:           place.Name,
		PostalCode:     place.PostalCode,
		Latitude:       &place.Latitude,
		Longitude:      &place.Longitude,
		NeighborhoodID: place.NeighborhoodId,
		CreatedAt:      place.CreatedAt,
		UpdatedAt:      place.UpdatedAt,
	}
}

type PagePlaceResponseDTO struct {
	Content          []PlaceResponseDTO      `json:"content"`
	Empty            bool                    `json:"empty"`
	First            bool                    `json:"first"`
	Last             bool                    `json:"last"`
	Number           int                     `json:"number"`
	NumberOfElements int                     `json:"numberOfElements"`
	Pageable         repository.PageableInfo `json:"pageable"`
	Size             int                     `json:"size"`
	Sort             repository.SortInfo     `json:"sort"`
	TotalElements    int64                   `json:"totalElements"`
	TotalPages       int                     `json:"totalPages"`
}

func NewPagePlaceResponseDTO(page *repository.PageResponse[entity.Place]) *PagePlaceResponseDTO {

	content := make([]PlaceResponseDTO, 0)
	for _, place := range page.Content {
		content = append(content, *NewPlaceResponseDTO(&place))
	}

	return &PagePlaceResponseDTO{
		Content:          content,
		Empty:            page.Empty,
		First:            page.First,
		Last:             page.Last,
		Number:           page.Number,
		NumberOfElements: page.NumberOfElements,
		Pageable:         page.Pageable,
		Size:             page.Size,
		Sort:             page.Sort,
		TotalElements:    page.TotalElements,
		TotalPages:       page.TotalPages,
	}
}
