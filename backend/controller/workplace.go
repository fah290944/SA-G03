package controller

import (

        	"github.com/fah290944/sa-65-example/entity"

           "github.com/gin-gonic/gin"

           "net/http"	

)

// POST /WorkPlace//เพิ่มข้อมูลใน DB

func CreateWorkPlace(c *gin.Context) {

	var workPlace entity.WorkPlace

	if err := c.ShouldBindJSON(&workPlace); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}


	if err := entity.DB().Create(&workPlace).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}

	c.JSON(http.StatusOK, gin.H{"data": workPlace})

}

// GET /WorkPlace/:id ดึงข้อมูลเฉพาะตัวที่ต้องการ

func GetWorkPlace(c *gin.Context) {

	var workPlace entity.WorkPlace

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM work_places WHERE id = ?", id).Scan(&workPlace).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": workPlace})

}

// GET /WorkPlace ดึงทั้งหมดใน DB ของ WorkPlace
func ListWorkPlaces(c *gin.Context) {

	var workPlace []entity.WorkPlace //[] อาเรย์

	if err := entity.DB().Raw("SELECT * FROM work_places").Scan(&workPlace).Error;
		   err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": workPlace})

}

// DELETE /doctor/:id

func DeleteWorkPlace(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM work_places WHERE id = ?", id); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "work_places not found"})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /doctor

func UpdateWorkPlace(c *gin.Context) {

	var workPlace entity.WorkPlace

	if err := c.ShouldBindJSON(&workPlace); err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	if tx := entity.DB().Where("id = ?", workPlace.ID).First(&workPlace); tx.RowsAffected == 0 {

		   c.JSON(http.StatusBadRequest, gin.H{"error": "schedule not found"})

		   return

	}



	if err := entity.DB().Save(&workPlace).Error; err != nil {

		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		   return

	}



	c.JSON(http.StatusOK, gin.H{"data": workPlace})

}