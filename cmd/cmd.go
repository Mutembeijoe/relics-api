package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	. "github.com/mutembeijoe/smartshop_api/utils"
	"log"
	"os"
)

type application struct {
	DB *gorm.DB
}

var (
	app  *application
	port string
	r *gin.Engine
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

	app = &application{DB: db}

	if err != nil {
		log.Fatalln("Failed to Initiated db", err.Error())
	}
	//AutoMigrate DB
	AutoMigrate(app.DB)

	//SEED_DB
	SeedDB(app.DB)
}

func Run() {
	defer app.DB.Close()

	r = gin.Default()

	registerRoutes()

	LogInfo("Listening on Port : ", port)
	log.Fatalln(r.Run(":" + port))

}
