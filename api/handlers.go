package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mutembeijoe/smartshop_api/postgres"
	. "github.com/mutembeijoe/smartshop_api/utils"
	"net/http"
)

type Application struct {
	DB *gorm.DB
}

func (app *Application) GetProducts(c *gin.Context) {
	LogInfo("Getting all products...")

	var products []postgres.Product

	if err := app.DB.Set("gorm:auto_preload", true).Find(&products).Error; err != nil {
		LogError("Failed to Fetch products")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch products : %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (app *Application) GetCategories(c *gin.Context) {
	LogInfo("Fetching All categories...")
	var categories []postgres.Category

	if err := app.DB.Find(&categories).Error; err != nil {
		LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

//func (app *Application)AddProduct(c *gin.Context){
//	LogInfo("Attempting to Insert Product into DB")
//	//c.ShouldBindJSON()
//
//}
