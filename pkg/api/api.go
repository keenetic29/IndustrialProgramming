package api

import (
	"github.com/gin-gonic/gin"
	"rest.com/pkg/handler"
)

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}

func fillEndpoints(router *gin.Engine) {

	router.GET("/products", handler.GetProducts)

	router.GET("/products/:id", handler.GetProductByID)

	router.POST("/products", handler.CreateProduct)

	router.PUT("/products/:id", handler.UpdateProduct)

	router.DELETE("/products/:id", handler.DeleteProduct)

}