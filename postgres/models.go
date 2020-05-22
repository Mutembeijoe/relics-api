package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name" gorm:"type:varchar(50)"`
	CategorySlug string `json:"category_slug" gorm:"not null"`
	//Products     []Product
}

type Product struct {
	gorm.Model
	ProductName string  `json:"product_name" binding:"required" gorm:"type:varchar(50) not null"`
	ProductSlug string  `json:"product_slug" gorm:"not null"`
	Price       float32 `json:"price" binding:"required" gorm:"not null"`
	Details     string  `json:"details" gorm:"type:text"`
	ImageUrl    string  `json:"imageUrl" binding:"required" gorm:"type:text not null"`
	Options     postgres.Jsonb
	Category    Category `json:"category"`
	CategoryID  uint     `json:"-"`
}
