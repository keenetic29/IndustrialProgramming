package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// Получение всех продуктов
func GetProducts(c *gin.Context) {
	var products []model.Product
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	name := c.Query("name")
	sort := c.Query("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	products, err = repository.GetProducts(limitInt, pageInt, name, sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}