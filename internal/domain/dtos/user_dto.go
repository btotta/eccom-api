package dtos

import (
	"eccom-api/internal/domain/entity"
	"strings"
)

type UserResponseDTO struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

func NewUserResponseDTO(user *entity.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   string(user.Role),
		Status: string(user.Status),
	}
}

type CreateUserDTO struct {
	Name            string `json:"name" binding:"required,min=3,max=100"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

func (c *CreateUserDTO) ToEntity() (*entity.User, error) {
	return entity.NewUser(
		strings.TrimSpace(c.Name),
		strings.TrimSpace(strings.ToLower(c.Email)),
		strings.TrimSpace(c.Password),
		entity.Customer)
}
