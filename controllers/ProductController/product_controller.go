package ProductController

import (
	"github.com/alfredoptarigan/go-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	// get all products

	var products []models.Product // create a slice of products (that type is Product)
	models.DB.Find(&products)     // find all products and store them in the slice

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func Show(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error!"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Store(c *gin.Context) {
	var product models.Product

	// get json data from request body
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Destroy(c *gin.Context) {
	var product models.Product

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if models.DB.Where("id = ?", id).Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Product deleted!"})
}
