package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/repository"
)

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Deletes a product from the repository by its ID. Requires authentication.
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  map[string]string  "Product deleted successfully"
// @Failure      404  {object}  map[string]string  "Product not found"
// @Failure      400  {object}  map[string]string  "Failed to delete product"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      403  {object}  map[string]string  "Invalid request or error message"
// @Router       /products/{id} [delete]
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