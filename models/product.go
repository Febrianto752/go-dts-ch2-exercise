package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(150)"`
	Brand     string `gorm:"not null;type:varchar(50)"`
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}

	return
}
