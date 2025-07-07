package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model

	QuizID            uint
	Text              string `gorm:"type:text;not null"`
	Type              string `gorm:"not null"`
	ScoreValue        int    `gorm:"default:1"`
	Answers           []Answer
	CorrectAnswerText *string `gorm:"type:text"`
}
