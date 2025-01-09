package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// Получение всех продуктов
func GetProducts(c *gin.Context) {
	var products []model.Product

	products, err := repository.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}