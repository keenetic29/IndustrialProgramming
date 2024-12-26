package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// Обновление существующего продукта
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var updatedProduct model.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, product := range *repository.GetProducts() {
		if product.ID == id {
			(*repository.GetProducts())[i] = updatedProduct
			(*repository.GetProducts())[i].ID = id // Убедиться, что ID не меняется
			c.JSON(http.StatusOK, (*repository.GetProducts())[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}