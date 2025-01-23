package api

import (
	"github.com/gin-gonic/gin"
	"rest.com/pkg/auth"
	"rest.com/pkg/handler"
	"rest.com/pkg/handler/products"
	taskhandler "rest.com/pkg/handler/taskHandler"
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
	protected.GET("/products", products.GetProducts)
    protected.GET("/products/:id", products.GetProductByID)
    protected.POST("/products",  auth.AdminCheck(), products.CreateProduct)
    protected.PUT("/products/:id", auth.AdminCheck(), products.UpdateProduct)
    protected.DELETE("/products/:id", auth.AdminCheck(), products.DeleteProduct)

    protected.POST("/tasks", auth.AuthMiddleware(), taskhandler.CreateTask)
	protected.GET("/tasks/:id", auth.AuthMiddleware(), taskhandler.GetTask)
	protected.DELETE("/tasks/:id", auth.AuthMiddleware(), taskhandler.CancelTask)
}
