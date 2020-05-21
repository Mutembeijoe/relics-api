package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/mutembeijoe/smartshop_api/utils"
	"log"
)

type application struct {
	DB *gorm.DB
}

var app application

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file")
	}
	db, err := utils.InitDB()

	app.DB = db

	if err != nil {
		log.Fatalln("Failed to Initiated db", err.Error())
	}
}

func main() {
	defer app.DB.Close()
}
