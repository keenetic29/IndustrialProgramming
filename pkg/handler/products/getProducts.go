package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// GetProducts godoc
// @Summary      Get products with filters
// @Description  Retrieves a list of products with optional filters and pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        page   query     int     false  "Page number (default: 1)"
// @Param        limit  query     int     false  "Number of products per page (default: 10)"
// @Param        name   query     string  false  "Filter by product name"
// @Param        sort   query     string  false  "Sort column"
// @Success      200    {array}   model.Product
// @Failure      400    {object}  map[string]string  "Invalid request"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      500    {object}  map[string]string  "Internal server error"
// @Router       /products [get]
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