package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Updates an existing product in the repository
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        product  body      model.Product  true  "Updated product data"
// @Success      200      {object}  map[string]string  "Product updated successfully"
// @Failure      400      {object}  map[string]string  "Invalid request or update failed"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      403  {object}  map[string]string  "Invalid request or error message"
// @Router       /products [put]
func UpdateProduct(c *gin.Context) {

	var updatedProd model.Product

	if err := c.BindJSON(&updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if !repository.HandleProductFieldsError(c, updatedProd) {
		return
	}

	if err := repository.UpdateProduct(updatedProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}