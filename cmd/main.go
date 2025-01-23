package main

import (
	"rest.com/pkg/api"
	"rest.com/pkg/db"
)


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
func main() {
	db.InitDB()
	api.StartServer(":8080")
}
