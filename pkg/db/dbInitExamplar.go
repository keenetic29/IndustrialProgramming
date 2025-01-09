package db

import "rest.com/pkg/model"

var initProducts = []model.Product{
	{ID: 1, Name: "Телефон", Description: "Смартфон с 5G", Cost: 50000, Count: 100, ManufacturerId: 1, SupplierId: 1},
	{ID: 2, Name: "Ноутбук", Description: "Ультрабук для работы", Cost: 75000, Count: 50, ManufacturerId: 1, SupplierId: 1},
	{ID: 3, Name: "Планшет", Description: "Компактный и удобный", Cost: 30000, Count: 60, ManufacturerId: 1, SupplierId: 1},
}

var initUsers = []model.User{
	{ID: 1, Username: "user1", Password: "12345", AccessLevel: "User"},
	{ID: 2, Username: "admin1", Password: "12345", AccessLevel: "Admin"},
}