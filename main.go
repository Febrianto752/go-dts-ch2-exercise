package main

import (
	"errors"
	"fmt"
	"tmp_latihan/database"
	"tmp_latihan/models"

	"gorm.io/gorm"
)

var PORT = ":8080"

func main() {
	// database.StartDB()

	// createUser("febrianto.bekasi3@gmail.com")
	// getUserById(1)
	// updateUserById(1, "email.baru@gmail.com")
	// createProduct(1, "Aukey", "keyboard rgb v1")
	// getUsersWithProducts()
	deleteProductById(10)
}

func createUser(email string) {
	db := database.GetDB()

	fmt.Println(db)
	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data :", err)
		return

	}

	fmt.Println("New User Data :", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user", err)
	}

	fmt.Printf("User Data : %+v \n", user)
}

func updateUserById(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	// atau
	// db.Table("users").Where("id = ?", id).Updates(map[string]interface{}{"email": email})

	if err != nil {
		fmt.Println("Error updateing user data :", err)
		return
	}

	fmt.Printf("Update user's email : %+v \n", user.Email)
}

// Product session
func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserId: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data :", err.Error())
		return
	}

	fmt.Println("New Product Data :", Product)

}
func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{Id: 1}

	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products :", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	affctedRows := db.Where("id = ?", id).Delete(&product).RowsAffected
	if affctedRows == 0 {
		fmt.Println("Error deleting product")
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}
