package controller

import (

        	"github.com/fah290944/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"	

)

// POST /medActivity //เพิ่มข้อมูลใน DB

func CreateMedActivity(c *gin.Context) {

	var medActivity entity.MedActivity

	if err := c.ShouldBindJSON(&medActivity); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}


	if err := entity.DB().Create(&medActivity).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}

	c.JSON(http.StatusOK, gin.H{"data": medActivity})

}

// GET /medActivity/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetMedActivity(c *gin.Context) {

	var medActivity entity.MedActivity

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM med_activities WHERE id = ?", id).Scan(&medActivity).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": medActivity})

}

// GET /doctor ดึงทั้งหมดใน DB ของ หมอ
func ListMedActivitys(c *gin.Context) {

	var medActivity []entity.MedActivity

	if err := entity.DB().Raw("SELECT * FROM med_activities").Scan(&medActivity).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": medActivity})

}

// DELETE /doctor/:id

func DeleteMedActivity(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM med_activities WHERE id = ?", id); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "med_activities not found"})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateMedActivity(c *gin.Context) {

	var medActivity entity.MedActivity

	if err := c.ShouldBindJSON(&medActivity); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	if tx := entity.DB().Where("id = ?", medActivity.ID).First(&medActivity); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "medactivity not found"})

		   return

	}



	if err := entity.DB().Save(&medActivity).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": medActivity})

}