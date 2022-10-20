package main

import (
	"github.com/fah290944/sa-65-example/controller"
	"github.com/fah290944/sa-65-example/middlewares"

	"github.com/fah290944/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Doctor Routes
			protected.GET("/doctors", controller.ListDoctors)
			protected.GET("/doctor/:id", controller.GetDoctor)
			protected.POST("/doctors", controller.CreateDoctor)

			//workplace
			protected.GET("/workPlace", controller.ListWorkPlaces)
			protected.GET("/workPlace/:id", controller.GetWorkPlace)

			//medactivity
			protected.GET("/medActivitys", controller.ListMedActivitys)
			protected.GET("/medActivity/:id", controller.GetMedActivity)

			//schedule
			protected.GET("/schedules", controller.ListSchedules)
			protected.GET("/schedule/:id", controller.GetSchedule)
			protected.POST("/saveschedule", controller.CreateSchedule)

		}
	}

	r.POST("/Login", controller.Login)

	// Run the server

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
