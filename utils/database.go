package utils

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	postgres_db "github.com/mutembeijoe/smartshop_api/postgres"
	"os"
	"strings"
)

//InitDB connects to db and returns *gorm.DB or an error
func InitDB() (*gorm.DB, error) {

	//"host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SSL"))

	db, err := gorm.Open("postgres", uri)

	if err != nil {
		LogError("Connection to db failed")
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		LogWarning("Failed to Ping DB")
		return nil, err
	}

	LogInfo("Connected to DB....")

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	LogInfo("Running Auto Migration to create DATABASE TABLES")
	if err := db.AutoMigrate(&postgres_db.Category{}, &postgres_db.Product{}).Error; err != nil {
		LogError(err.Error())
		LogError("Failed to successfully Auto migrate all Tables...")

		return
	}

	LogInfo("Finished Running Auto Migration")
}

func SeedDB(db *gorm.DB) {
	if seedDB := strings.ToLower(os.Getenv("SEED_DB")); seedDB == "true" {
		LogInfo("Seeding DB.....")
		categories := []postgres_db.Category{
			{
				CategoryName: "T-shirts",
				Options:      postgres.Jsonb{RawMessage: json.RawMessage(`{"sizes": "['xl', 'lg', 'md', 'sm']"}`)},
			},
			{
				CategoryName: "Hoodies",
				Options:      postgres.Jsonb{RawMessage: json.RawMessage(`{"sizes": "['xl', 'lg', 'md', 'sm']"}`)},
			},
		}
		LogInfo("Starting to seed Categories ...")

		for _, category := range categories {
			LogInfo(fmt.Sprintf("INSERTING %s category into DB", category.CategoryName))
			err := db.Create(&category).Error

			if err != nil {
				LogError(fmt.Sprintf("Failed to insert %s category into DB : %s", category.CategoryName, err.Error()))
				return
			}
		}

		LogInfo("Finished seeding categories")

		products := []postgres_db.Product{
			{
				ProductName: "Girl Power Hoodie Black",
				ProductSlug: "girl-power-hoodie -black",
				Price:       1150,
				Description: "Girl Power Hoodie with Rosie the Riveter Poster, Color Black",
				ImageUrl:    "https://res.cloudinary.com/myloxyloto/image/upload/v1589988178/smartshop/Highcompressed_1416322864_h6xxmh.png",
				CategoryID:  2,
			},
			{
				ProductName: "Star Wars Valentines T-shirt",
				ProductSlug: "star-wars-valentines-t-shirt",
				Price:       650,
				Description: "Darth Vader Star Wars Valentine Red T-shirt",
				ImageUrl:    "https://res.cloudinary.com/myloxyloto/image/upload/v1589988198/smartshop/Highcompressed_2114311319_nm4amj.png",
				CategoryID:  1,
			},
		}
		LogInfo("Starting to seed products...")
		for index, product := range products {
			LogInfo(fmt.Sprintf("Insering product %d into db", index+1))
			err := db.Create(&product).Error
			if err != nil {
				LogError(fmt.Sprintf("Failed to Insert product %d into db", index+1))
				return
			}
		}
		LogInfo("Finished seeding Products...")
	}
}
