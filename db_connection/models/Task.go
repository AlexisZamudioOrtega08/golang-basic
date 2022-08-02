package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID          int    `gorm:"primary_key"`
	Title       string `gorm:"not null; type:varchar(100)"`
	Description string
	Completed   bool `gorm:"default:false"`
	UserID      uint `gorm:"not null"`
}
