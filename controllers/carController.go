package controllers

import (
	"fmt"
	"net/http"
	"tmp_latihan/database"
	"tmp_latihan/models"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

func GetAllCars(c *gin.Context){
	var db = database.GetDB()

	var cars []models.Car 
	err := db.Find(&cars).Error 

	if err != nil{
		fmt.Println("Error getting car data :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

func GetOneCars(c *gin.Context){
	var db = database.GetDB()

	var car models.Car 

	err := db.First(&car, "id = ?", c.Param("id")).Error 

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"data car": car})
}


func CreateCars(c *gin.Context){
	var db = database.GetDB()

	var input models.Car 

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	fmt.Println(input)
	// carInput := models.Car{Pemilik: input.Pemilik, Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars}
	err := db.Create(&input).Error 

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateCars(c *gin.Context){
	var db = database.GetDB()

	var car models.Car 
	err := db.First(&car, "id = ?", c.Param("id")).Error 

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return 
	}

	// validate input
	var input models.Car 
	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err = db.Model(&car).Where("id = ?", car.Id).Updates(models.Car{Pemilik: input.Pemilik, Merk: input.Merk, Harga: input.Harga, Typecars: input.Typecars}).Error
			

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated car"})
}

func DeleteCars(c *gin.Context){
	var db = database.GetDB()

	var carDelete models.Car 

	err := db.First(&carDelete, "id = ?", c.Param("id")).Error 

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return 
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"message": "successfuly to deleted car"})
}

// var CarDatas = []Car{}

// func CreateCar(ctx *gin.Context) {
// 	var newCar Car

// 	if err := ctx.ShouldBindJSON(&newCar); err != nil {
// 		// mengirim status error dan pesan errorny
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
// 	CarDatas = append(CarDatas, newCar)

// 	// mengirim response status dan data reponse ke client
// 	ctx.JSON(http.StatusCreated, gin.H{
// 		"car": newCar,
// 	})
// }

// func UpdateCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var updatedCar Car

// 	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	for i, car := range CarDatas {
// 		if carID == car.CarID {
// 			condition = true
// 			CarDatas[i] = updatedCar
// 			CarDatas[i].CarID = carID
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data Not Found",
// 			"error_message": fmt.Sprintf("Car with id %v not found", carID),
// 		})

// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
// 	})
// }

// func GetCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var carData Car

// 	for i, car := range CarDatas {
// 		if carID == car.CarID {
// 			condition = true
// 			carData = CarDatas[i]
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data Not Found",
// 			"error_message": fmt.Sprintf("Car with id %v not found", carID),
// 		})

// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"car": carData,
// 	})
// }

// func DeleteCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var carIndex int

// 	for i, car := range CarDatas {
// 		if carID == car.CarID {
// 			condition = true
// 			carIndex = i
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data Not Found",
// 			"error_message": fmt.Sprintf("car with id %v not found", carID),
// 		})

// 		return
// 	}

// 	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
// 	CarDatas[len(CarDatas)-1] = Car{}
// 	CarDatas = CarDatas[:len(CarDatas)-1]

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
// 	})

// }
