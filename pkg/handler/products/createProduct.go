package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
	"rest.com/pkg/repository"
)


type ReqProduct struct {
	Name        	string  `json:"name"`
	Description 	string  `json:"description"`
	Cost       		float64 `json:"cost"`
	Count			int     `json:"count"`
	ManufacturerId 	int 	`json:"manufacturerId"`
	SupplierId		int 	`json:"supplierId"`
}


// CreateProduct godoc
// @Summary      Create a product
// @Description  Creating product
// @Tags         products
// @Accept       json
// @Produce      json
// @Security JWT
// @Param        product body   ReqProduct  true  "Add product"
// @Success      201  {object}  model.Product
// @Failure      400  {object}  map[string]string  "Invalid request or error message"
// @Failure      401  {object}  map[string]string  "Invalid request or error message"
// @Failure      403  {object}  map[string]string  "Invalid request or error message"
// @Router       /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct model.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Автоматическая генерация ID
	var maxProductID int = 0
	products, _ := repository.GetAllProducts()
	for _, product := range products {
		if product.ID > maxProductID {
			maxProductID = product.ID
		}
	}
	newProduct.ID = maxProductID + 1

	if err := repository.AddProducts([]model.Product{newProduct}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}