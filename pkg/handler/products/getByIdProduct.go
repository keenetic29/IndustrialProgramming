package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// Получение продукта по ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product model.Product
	product, err := repository.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}