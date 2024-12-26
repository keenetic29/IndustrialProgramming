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
	newProduct.ID = len(*repository.GetProducts()) + 1
	*repository.GetProducts() = append(*repository.GetProducts(), newProduct)

	c.JSON(http.StatusCreated, newProduct)
}