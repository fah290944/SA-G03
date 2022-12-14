package controller

import (

        	"github.com/fah290944/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"	

)

// POST /doctor //เพิ่มข้อมูลใน DB

func CreateDoctor(c *gin.Context) {

	var doctor entity.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}


	if err := entity.DB().Create(&doctor).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}

	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

// GET /doctor/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetDoctor(c *gin.Context) {

	var doctor entity.Doctor

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM doctors WHERE id = ?", id).Scan(&doctor).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

// GET /doctor ดึงทั้งหมดใน DB ของ หมอ
func ListDoctors(c *gin.Context) {

	var doctor []entity.Doctor

	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctor).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": doctor})

}

// DELETE /doctor/:id

func DeleteDoctor(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM doctors WHERE id = ?", id); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "doctors not found"})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateDoctor(c *gin.Context) {

	var doctor entity.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	if tx := entity.DB().Where("id = ?", doctor.ID).First(&doctor); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		   return

	}



	if err := entity.DB().Save(&doctor).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": doctor})

}