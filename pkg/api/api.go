package api

import (
	"github.com/gin-gonic/gin"
	"rest.com/pkg/auth"
	"rest.com/pkg/handler"
)

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}

func fillEndpoints(router *gin.Engine) {
    router.POST("/register", handler.Registration)
    router.POST("/login", handler.Login)

    protected := router.Group("/")
    protected.Use(auth.AuthMiddleware())
	protected.GET("/products", handler.GetProducts)
    protected.GET("/products/:id", handler.GetProductByID)
    protected.POST("/products",  auth.AdminCheck(), handler.CreateProduct)
    protected.PUT("/products/:id", auth.AdminCheck(), handler.UpdateProduct)
    protected.DELETE("/products/:id", auth.AdminCheck(), handler.DeleteProduct)
}
