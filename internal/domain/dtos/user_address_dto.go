package dtos

import (
	"eccom-api/internal/domain/entity"
	"eccom-api/internal/domain/repository"
	"time"
)

type CreateUserAddressDTO struct {
	StatedId       uint   `json:"state_id"`
	CityId         uint   `json:"city_id"`
	NeighborhoodId uint   `json:"neighborhood_id"`
	PlaceId        uint   `json:"place_id"`
	Logradouro     string `json:"logradouro"`
	Numero         string `json:"numero"`
	Complemento    string `json:"complemento"`
}

type UserAddressResponseDTO struct {
	ID             uint          `json:"id"`
	StateID        uint          `json:"state_id"`
	CityID         uint          `json:"city_id"`
	NeighborhoodID uint          `json:"neighborhood_id"`
	PlaceID        uint          `json:"place_id"`
	Numero         string        `json:"numero"`
	Complemento    string        `json:"complemento"`
	Status         entity.Status `json:"status"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

func NewUserAddressResponseDTO(userAddress *entity.UserAddress) *UserAddressResponseDTO {
	return &UserAddressResponseDTO{
		ID:             userAddress.ID,
		StateID:        userAddress.StateID,
		CityID:         userAddress.CityID,
		NeighborhoodID: userAddress.NeighborhoodID,
		PlaceID:        userAddress.PlaceID,
		Numero:         userAddress.Numero,
		Complemento:    userAddress.Complemento,
		Status:         userAddress.Status,
		CreatedAt:      userAddress.CreatedAt,
		UpdatedAt:      userAddress.UpdatedAt,
	}
}

type PageUserAddressResponseDTO struct {
	Content          []UserAddressResponseDTO `json:"content"`
	Empty            bool                     `json:"empty"`
	First            bool                     `json:"first"`
	Last             bool                     `json:"last"`
	Number           int                      `json:"number"`
	NumberOfElements int                      `json:"numberOfElements"`
	Pageable         repository.PageableInfo  `json:"pageable"`
	Size             int                      `json:"size"`
	Sort             repository.SortInfo      `json:"sort"`
	TotalElements    int64                    `json:"totalElements"`
	TotalPages       int                      `json:"totalPages"`
}

func NewPageUserAddressResponseDTO(page *repository.PageResponse[entity.UserAddress]) *PageUserAddressResponseDTO {

	content := make([]UserAddressResponseDTO, 0)
	for _, neighborhood := range page.Content {
		content = append(content, *NewUserAddressResponseDTO(&neighborhood))
	}

	return &PageUserAddressResponseDTO{
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
