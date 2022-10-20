package entity

import (

	"gorm.io/gorm"
)

type MedActivity struct {
	gorm.Model
	Name  string

	Schedule []Schedule  `gorm:"foreignKey:MedActivityID"`
}