package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Title           string `gorm:"not null"`
	Description     *string
	CategoryID      uint
	Category        Category
	CreatorID       uint
	Creator         User
	IsPublic        *bool `gorm:"default:true"`
	TimeLimit       *int  `gorm:"default:60"`
	AttemptsAllowed *int  `gorm:"default:1"`
	Questions       []Question
}
