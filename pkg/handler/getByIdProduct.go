package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/repository"
)

// Получение продукта по ID
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	for _, product := range *repository.GetProducts() {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}