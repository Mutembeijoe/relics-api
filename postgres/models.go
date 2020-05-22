package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Category struct {
	gorm.Model
	CategoryName string         `json:"category_name" gorm:"type:varchar(50)"`
	CategorySlug string         `json:"category_slug" gorm:"not null"`
	Options      postgres.Jsonb `json:"options"`
}

type Product struct {
	gorm.Model
	ProductName string   `json:"product_name" gorm:"type:varchar(50) not null"`
	ProductSlug string   `json:"product_slug" gorm:"not null"`
	Price       float32  `json:"price" gorm:"not null"`
	Description string   `json:"details" gorm:"type:text"`
	ImageUrl    string   `json:"imageUrl" gorm:"type:text not null"`
	Category    Category `json:"category"`
	CategoryID  uint     `json:"-"`
}

//Verify Existence of Category Before saving Product
func (p *Product) BeforeSave(db *gorm.DB) (err error) {
	var c Category
	c.ID = p.CategoryID
	if err = db.First(&c).Error; gorm.IsRecordNotFoundError(err){
		return err
	}
	return
}

// Preload Category After Creating Product
func (p *Product) AfterCreate(db *gorm.DB) (err error) {
	var c Category
	db.First(&c, p.CategoryID)
	p.Category = c
	return
}
