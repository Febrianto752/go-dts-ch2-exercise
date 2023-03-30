package database

import (
	"fmt"
	"log"
	"tmp_latihan/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tmp struct {
	db *gorm.DB
}

var (
	host     = "localhost"
	user     = "postgres"
	password = "admin123"
	dbPort   = "5432"
	dbname   = "learning-gorm"
	DB       *gorm.DB
	err      error
)

// migration table
func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() (db *gorm.DB) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	return
}
