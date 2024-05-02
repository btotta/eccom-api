package repository

import (
	"math"
	"strings"

	"gorm.io/gorm"
)

type SortInfo struct {
	Empty    bool `json:"empty"`
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
}

type PageableInfo struct {
	Offset     int      `json:"offset"`
	Paged      bool     `json:"paged"`
	PageNumber int      `json:"pageNumber"`
	PageSize   int      `json:"pageSize"`
	Sort       SortInfo `json:"sort"`
	Unpaged    bool     `json:"unpaged"`
}

type PageResponse[T any] struct {
	Content          []T          `json:"content"`
	Empty            bool         `json:"empty"`
	First            bool         `json:"first"`
	Last             bool         `json:"last"`
	Number           int          `json:"number"`
	NumberOfElements int          `json:"numberOfElements"`
	Pageable         PageableInfo `json:"pageable"`
	Size             int          `json:"size"`
	Sort             SortInfo     `json:"sort"`
	TotalElements    int64        `json:"totalElements"`
	TotalPages       int          `json:"totalPages"`
}

func Paginate[T any](query *gorm.DB, page int, size int, orderBy []string) (*PageResponse[T], error) {

	var data []T
	var totalElements int64

	for _, order := range orderBy {
		order := strings.Split(order, ",")

		if strings.TrimSpace(order[0]) == "" {
			continue
		}
		if len(order) == 1 {
			query = query.Order(order[0] + " ASC")
		} else {
			query = query.Order(order[0] + " " + order[1])
		}
	}

	result := query.Model(&data).Count(&totalElements)
	if result.Error != nil {
		return nil, result.Error
	}

	offset := page * size
	query = query.Offset(offset).Limit(size)

	if err := query.Find(&data).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalElements) / float64(size)))
	lastPage := page == totalPages-1

	response := PageResponse[T]{
		Content:          data,
		Empty:            len(data) == 0,
		First:            page == 0,
		Last:             lastPage,
		Number:           page,
		NumberOfElements: len(data),

		Pageable: PageableInfo{
			Offset:     offset,
			Paged:      true,
			PageNumber: page,
			PageSize:   size,

			Sort: SortInfo{
				Empty:    len(orderBy) == 0,
				Sorted:   len(orderBy) > 0,
				Unsorted: len(orderBy) == 0,
			},
			Unpaged: false,
		},

		Size: size,
		Sort: SortInfo{
			Empty:    len(orderBy) == 0,
			Sorted:   len(orderBy) > 0,
			Unsorted: len(orderBy) == 0,
		},
		TotalElements: totalElements,
		TotalPages:    totalPages,
	}

	return &response, nil
}
