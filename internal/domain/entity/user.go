package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// user role enum
type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"index"`
	Email    string `json:"email" gorm:"index,unique"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	Status   Status `json:"status"`
}

func NewUser(name, email, password string, role Role) (*User, error) {

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error hashing password")
	}

	return &User{
		Model:    gorm.Model{CreatedAt: time.Now()},
		Name:     name,
		Email:    email,
		Password: string(hashPwd),
		Role:     role,
		Status:   Active,
	}, nil
}

func (u *User) ComparePassword(password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		return errors.New("credenciais inválidas")
	}

	return nil
}
