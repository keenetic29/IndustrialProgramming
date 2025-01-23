package api

import (
	_ "rest.com/docs"
	"github.com/gin-gonic/gin"
	"rest.com/pkg/auth"
	"rest.com/pkg/handler"
	"rest.com/pkg/handler/products"
	taskhandler "rest.com/pkg/handler/taskHandler"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

)

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}


// @title           Swagger Example API
// @version         1.0
// @description     This is tutorial project on creating rest api.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

// @termsOfService  http://swagger.io/terms/

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
