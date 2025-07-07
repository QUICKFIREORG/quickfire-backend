package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model

	QuestionID uint
	Text       *string `gorm:"type:text;not null"`
	ImageURL   *string `gorm:"type:text"` // For image-based aswers
	IsCorrect  bool    `gorm:"default:false"`
}
