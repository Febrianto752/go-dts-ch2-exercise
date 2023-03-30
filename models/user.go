package models

import "time"

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(100)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
