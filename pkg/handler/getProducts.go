package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/repository"
)

// Получение всех продуктов
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetProducts())
}