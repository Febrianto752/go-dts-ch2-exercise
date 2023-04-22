package routers

import (
	"tmp_latihan/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Read
	router.GET("/cars/:id", controllers.GetOneCars)
	// Create
	router.POST("/cars", controllers.CreateCars)
	// Read All
	router.GET("/cars", controllers.GetAllCars)
	// Update
	router.PUT("/cars/:id", controllers.UpdateCars)
	// Delete 
	router.DELETE("/cars/:id", controllers.DeleteCars)
	

	return router
}
