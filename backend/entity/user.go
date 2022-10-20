package entity

import (

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name  string
	Email string `gorm:"uniqueIndex"`
}

