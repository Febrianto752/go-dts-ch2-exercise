package routers

import (
	"tmp_latihan/controllers"

	_ "tmp_latihan/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"
)

// @title Car API
// @version 1.0
// @description This is a sample for managing cars
// @termsOfService http://swagger.io/terms/
// @Contact.name API Support
// @Contact.email saberkoder@swagger.io
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	

	return router
}
