package api

import (
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
		LogError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch Products",
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
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func (app *Application)AddCategory(c *gin.Context){
	LogInfo("Attempting to Insert Category into DB...")
	var cj categoryJson
	var category postgres.Category
	err:= c.ShouldBindJSON(&cj)
	if err!=nil{
		LogError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	category.CategoryName = cj.Name
	category.CategorySlug = GenerateSlug(cj.Name)


	if err= app.DB.Create(&category).Error; err!=nil{
		LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"payload":category,
		"success":"OK",
	})
}

//func (app *Application)AddProduct(c *gin.Context){
//	LogInfo("Attempting to Insert Product into DB")
//
//
//}
