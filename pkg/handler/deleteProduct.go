package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/repository"
)

// Удаление продукта
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	products, err := repository.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := repository.DeleteProduct(products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product deleted"})
}