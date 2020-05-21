package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
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

	fmt.Println(uri)
	db, err := gorm.Open("postgres", uri)

	if err != nil {
		log.Println("Connection to db failed")
		return nil, err
	}

	if err = db.DB().Ping(); err != nil {
		log.Println("Failed to Ping DB")
		return nil, err
	}

	log.Println("Connected to DB....")

	return db, nil
}
