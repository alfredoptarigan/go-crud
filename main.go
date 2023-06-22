package main

import (
	"github.com/alfredoptarigan/go-crud/controllers/ProductController"
	"github.com/alfredoptarigan/go-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, the server is running!",
		})
	})

	r.GET("/api/products", ProductController.Index)
	r.GET("/api/products/:id", ProductController.Show)
	r.POST("/api/products", ProductController.Store)
	r.PUT("/api/products/:id", ProductController.Update)
	r.DELETE("/api/products/:id", ProductController.Destroy)

	r.Run()

}
