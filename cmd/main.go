package main

import (
	"rest.com/pkg/api"
	"rest.com/pkg/db"
)

func main() {
	db.InitDB()
	api.StartServer(":8080")
}
