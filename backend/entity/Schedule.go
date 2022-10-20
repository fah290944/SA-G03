package entity

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Time  time.Time

	DoctorID *uint 	//อ้างอิง ID
	Doctor Doctor	 `gorm:"references:ID"` //อ้างอิงอะไรก็ได้ //gorm ไม่จำเป็น ยกเว้นจะอ้าง

	WorkPlaceID *uint 	//อ้างอิง ID
	WorkPlace WorkPlace 	`gorm:"references:ID"` //หน้าเรียกตัวแปร หลังคือ type

	MedActivityID *uint
	MedActivity MedActivity `gorm:"references:ID"`
}