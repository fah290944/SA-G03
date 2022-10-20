package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db
}
//เข้ารหัส
func SetupPasswordHash(pwd string) string {
    var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
    return string(password)
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		//ระบบ
		&Doctor{},
		&MedActivity{},
		&WorkPlace{},
		&Schedule{},
	)
	db = database
	//การเพิ่มข้อมูลตาราง Doctor
	Phonsak := Doctor{
		Name:  "Phonsak songsang",
		Email: "Phonsak@gmail.com",
		Passwordmed: SetupPasswordHash("Phonsak01"),
	}
	db.Model(&Doctor{}).Create(&Phonsak)

	Hanoi := Doctor{
		Name:  "Hanoi slotmachine",
		Email: "Hanoi@gmail.com",
		Passwordmed: SetupPasswordHash("Hanoiploy"),
	}
	db.Model(&Doctor{}).Create(&Hanoi)

	//การเพิ่มข้อมูลตาราง WorkPalce
	loca1 := WorkPlace{
		Name:    "Emergency and Accident Department",
		// Address: "Suranaree Building, 1st Floor",
	}
	db.Model(&WorkPlace{}).Create(&loca1)

	loca2 := WorkPlace{
		Name:    "Outpatient Department",
		// Address: "Suranaree Building, 1st Floor",
	}
	db.Model(&WorkPlace{}).Create(&loca2)
	
	//การเพิ่มข้อมูลตาราง MedActivity
	Activity1 := MedActivity{
		Name: "Operating Room",
	}
	db.Model(&MedActivity{}).Create(&Activity1)

	Activity2 := MedActivity{
		Name: "External Patient Examination",
	}
	db.Model(&MedActivity{}).Create(&Activity2)
	//การเพิ่มข้อมูลตาราง Schedule
	timeSchedule1 := time.Date(2022, 8, 30, 06, 00, 00, 00, time.Local)
	timeSchedule2 := time.Date(2022, 8, 30, 10, 00, 00, 00, time.Local)

	var Doctor1 Doctor
	db.Raw("SELECT * FROM Doctors WHERE Email = ?", "Phonsak@gmail.com").Scan(&Doctor1)

	var Doctor2 Doctor
	db.Raw("SELECT * FROM Doctors WHERE Email = ?", "Hanoi@gmail.com").Scan(&Doctor2)

	var Location1 WorkPlace
	var Location2 WorkPlace
	db.Raw("SELECT * FROM Work_Places WHERE Name = ?", "Emergency and Accident Department").Scan(&Location1)
	db.Raw("SELECT * FROM Work_Places WHERE Name = ?", "Outpatient Department").Scan(&Location2)

	var MA1 MedActivity
	var MA2 MedActivity
	db.Raw("SELECT * FROM Med_Activities WHERE Name = ?", "Operating Room").Scan(&MA1)
	db.Raw("SELECT * FROM Med_Activities WHERE Name = ?", "External Patient Examination").Scan(&MA2)

	db.Model(&Schedule{}).Create(&Schedule{
		Time:        timeSchedule1,
		Doctor:      Doctor1,
		WorkPlace:   loca1,
		MedActivity: Activity1,
	})

	db.Model(&Schedule{}).Create(&Schedule{
		Time:        timeSchedule2,
		Doctor:      Doctor2,
		WorkPlace:   loca2,
		MedActivity: Activity2,
	})

}
