package controller

import (
	"github.com/fah290944/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /schedule //เพิ่มข้อมูลใน DB

func CreateSchedule(c *gin.Context) {

	var schedule entity.Schedule
	var medActivity entity.MedActivity
	var workPlace entity.WorkPlace
	var doctor entity.Doctor

	if err := c.ShouldBindJSON(&schedule); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// ค้นหา medActivity ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.MedActivityID).First(&medActivity); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedActivity not found"})
		return
	}

	// ค้นหา workPlace ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.WorkPlaceID).First(&workPlace); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "WorkPlace not found"})
		return
	}

	//ค้นหา doctor ด้วย id
	if tx := entity.DB().Where("id = ?", schedule.DoctorID).First(&doctor); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Doctor not found"})
			return
	}

		// 12: สร้าง schedule
		sd := entity.Schedule{
			Doctor:  doctor,             // โยงความสัมพันธ์กับ Entity doctor
			WorkPlace: workPlace,                  // โยงความสัมพันธ์กับ workPlace
			MedActivity:    medActivity,               // โยงความสัมพันธ์กับ Entity medactivity
			Time: schedule.Time,
		}
	

	if err := entity.DB().Create(&sd).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": sd})

}

// GET /schedule/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetSchedule(c *gin.Context) {

	var schedule entity.Schedule

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM schedules WHERE id = ?", id).Scan(&schedule).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": schedule})

}

// GET /schedule ดึงทั้งหมดใน DB ของตารางเวลา
func ListSchedules(c *gin.Context) {

	var schedule []entity.Schedule

	if err := entity.DB().Preload("Doctor").Preload("WorkPlace").Preload("MedActivity").Raw("SELECT * FROM schedules").Find(&schedule).Error; err != nil {
//ดึงตารางย่อยมา .preload
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": schedule})

}

// DELETE /schedule/:id

func DeleteSchedule(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM schedules WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Schedule not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateSchedule(c *gin.Context) {

	var schedule entity.Schedule

	if err := c.ShouldBindJSON(&schedule); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", schedule.ID).First(&schedule); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "schedules not found"})

		return

	}

	if err := entity.DB().Save(&schedule).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": schedule})

}
