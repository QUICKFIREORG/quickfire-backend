package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    *string
	LastName     *string
	Email        string `gorm:"unique;not null"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"default:'user'"`
}
