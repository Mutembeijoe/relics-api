package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func initDB(){

	//"host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	uri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))

	db, err:= gorm.Open("postgres", uri)

	if err!=nil{
		log.Fatalln("Connection to db failed: ",err)
	}
	defer db.Close();

	if err = db.DB().Ping(); err!=nil{
		log.Fatalln("Failed to Ping DB", err)
	}

	log.Println("Connected to DB....")
}