package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func InitDB() {

	dsn := "host=localhost user=postgres password=1 dbname=rest port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	DB.AutoMigrate(&initProducts)
	DB.AutoMigrate(&initUsers)

	insertTestData()
}

func insertTestData() {
	DB.Exec("TRUNCATE products, users")
	DB.Create(initProducts)
	DB.Create(initUsers)
}