package entity

import (

	"gorm.io/gorm"
)

type WorkPlace struct {
	gorm.Model //ฟังก์ชันสำเร็จ
	Name  string
	// Address string

	Schedule []Schedule  `gorm:"foreignKey:WorkPlaceID"`
}

