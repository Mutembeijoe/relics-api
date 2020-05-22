package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	pg "github.com/mutembeijoe/smartshop_api/postgres"
	. "github.com/mutembeijoe/smartshop_api/utils"
	"net/http"
	"strconv"
)

type Application struct {
	DB *gorm.DB
}

func (app *Application) GetProducts(c *gin.Context) {
	LogInfo("Getting all products...")

	var products []pg.Product

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

	LogInfo("Successfully Fetched Products")
}

func (app *Application) GetCategories(c *gin.Context) {
	LogInfo("Fetching All categories...")
	var categories []pg.Category

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

func (app *Application) AddCategory(c *gin.Context) {
	LogInfo("Attempting to Insert Category into DB...")
	var cj categoryJson
	var category pg.Category
	err := c.ShouldBindJSON(&cj)
	if err != nil {
		LogError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	category.CategoryName = cj.Name
	category.CategorySlug = GenerateSlug(cj.Name)
	category.Options = postgres.Jsonb{RawMessage: cj.Options}

	if err = app.DB.Create(&category).Error; err != nil {
		LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "500 - Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"payload": category,
		"success": "OK",
	})

	LogInfo("Category Insert Successful")
}

func (app *Application) AddProduct(c *gin.Context) {
	LogInfo("Attempting to Insert Product into DB")
	var pj productJson
	var product pg.Product

	if err := c.ShouldBindJSON(&pj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product.ProductName = pj.Name
	product.ProductSlug = GenerateSlug(pj.Name)
	product.Price = pj.Price
	product.Description = pj.Description
	product.CategoryID = pj.CategoryID
	product.ImageUrl = pj.ImageUrl

	if err := app.DB.Create(&product).Set("gorm:auto_preload", true).Error; err != nil {
		LogError("Failed to Insert Product into DB : ", err.Error())
		if gorm.IsRecordNotFoundError(err){
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"Invalid CategoryID",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": " 500 - Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"payload": product,
		"success": "OK",
	})

	LogInfo("Product Insert Successful")
}

func (app *Application) GetProductByID(c *gin.Context) {
	LogInfo("Fetching a product by ID...")
	paramId := c.Param("id")
	var product pg.Product

	id, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	if err = app.DB.First(&product, id).Error; err != nil {
		LogInfo(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404 - Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payload": product,
		"success": "OK",
	})

	LogInfo("Successfully Fetched product")
}
