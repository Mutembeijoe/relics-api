package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(){
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"products":"All products Here",
		})
	} )
}