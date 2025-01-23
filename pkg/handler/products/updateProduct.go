package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// Обновление существующего продукта
func UpdateProduct(c *gin.Context) {

	var updatedProd model.Product
	

	if err := c.ShouldBindJSON(&updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if err := repository.UpdateProduct(updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}