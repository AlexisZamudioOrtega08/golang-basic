package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        int    `gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Tasks     []Task `gorm:"foreignkey:UserID"`
}
