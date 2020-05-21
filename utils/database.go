package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mutembeijoe/smartshop_api/postgres"
	"os"
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


func AutoMigrate(db *gorm.DB){
	if autoMigrate:= os.Getenv("AUTOMIGRATE"); autoMigrate=="true"{
		LogInfo("Running Auto Migration to create DATABASE TABLES")
		db.AutoMigrate(&postgres.Category{}, &postgres.Product{})
		LogInfo("Finished Running Auto Migration")
	}
}