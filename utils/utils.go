package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()

	if err!=nil{
		log.Fatalln("Failed to load env files")
	}

}