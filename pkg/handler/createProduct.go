package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

func CreateProduct(c *gin.Context) {
	var newProduct model.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Автоматическая генерация ID
	var maxProductID int = 0
	products, _ := repository.GetAllProducts()
	for _, product := range products {
		if product.ID > maxProductID {
			maxProductID = product.ID
		}
	}
	newProduct.ID = maxProductID + 1

	if err := repository.AddProducts([]model.Product{newProduct}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}