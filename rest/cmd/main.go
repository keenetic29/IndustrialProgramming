package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest.com/pkg/model"
)



var products = []model.Product{
	{ID: 1, Name: "Телефон", Description: "Смартфон с 5G", Price: 50000, Quantity: 100},
	{ID: 2, Name: "Ноутбук", Description: "Ультрабук для работы", Price: 75000, Quantity: 50},
	{ID: 3, Name: "Планшет", Description: "Компактный и удобный", Price: 30000, Quantity: 60},
}

// Получение всех продуктов
func getProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// Получение продукта по ID
func getProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Создание нового продукта
func createProduct(c *gin.Context) {
	var newProduct model.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Автоматическая генерация ID
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, newProduct)
}

// Обновление существующего продукта
func updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var updatedProduct model.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			products[i].ID = id // Убедиться, что ID не меняется
			c.JSON(http.StatusOK, products[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Удаление продукта
func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func main() {
	router := gin.Default()

	// CRUD-маршруты
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}
