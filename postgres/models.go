package postgres

import "github.com/jinzhu/gorm"

type Category struct{
	gorm.Model
	CategoryName string `json:"category_name" binding:"required" gorm:"type:varchar(50)"`
	Products []Product
}

type Product struct{
	gorm.Model
	ProductName string `json:"product_name" binding:"required" gorm:"type:varchar(50) not null"`
	Price float32 `json:"price" binding:"required" gorm:"not null"`
	Details string `json:"details" gorm:"type:text"`
	ImageUrl string `json:"imageUrl" binding:"required" gorm:"type:text not null"`
	CategoryID uint
}