package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	. "github.com/mutembeijoe/smartshop_api/utils"
	"log"
	"net/http"
	"os"
)

type application struct {
	DB *gorm.DB
}

var (
	app  *application
	port string
)

func init() {
	//Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file")
	}

	//	Init Logger
	InitLogger()

	//Get Port Number
	if port = os.Getenv("PORT_NUMBER"); port == "" {
		port = "8001"
	}

	//Init DB
	LogInfo("Contacting DB...")
	db, err := InitDB()

	app = &application{DB:db}

	if err != nil {
		log.Fatalln("Failed to Initiated db", err.Error())
	}
	//AutoMigrate DB
	AutoMigrate(app.DB)
}

func main() {
	defer app.DB.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": "Hello World",
		})
	})

	LogInfo("Listening on Port : ", port)
	log.Fatalln(r.Run(":"+port))

}
