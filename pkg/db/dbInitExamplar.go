package db

import "rest.com/pkg/model"

var initProducts = []model.Product{
	{ID: 1, Name: "Телефон", Description: "Смартфон с 5G", Cost: 60000, Count: 100, ManufacturerId: 1, SupplierId: 1},
	{ID: 2, Name: "Ноутбук", Description: "Ультрабук для работы", Cost: 75000, Count: 50, ManufacturerId: 1, SupplierId: 2},
	{ID: 3, Name: "Планшет", Description: "Компактный и удобный", Cost: 30000, Count: 60, ManufacturerId: 1, SupplierId: 1},
	{ID: 4, Name: "ПК", Description: "Мощный и быстрый", Cost: 80000, Count: 10, ManufacturerId: 1, SupplierId: 2},
	{ID: 5, Name: "Клавиатура", Description: "Механическая", Cost: 5000, Count: 15, ManufacturerId: 1, SupplierId: 2},
	{ID: 6, Name: "Мышка", Description: "Беспроводная", Cost: 3000, Count: 30, ManufacturerId: 1, SupplierId: 1},
}

var initUsers = []model.User{
	{ID: 1, Username: "user1", Password: "12345", AccessLevel: "User"},
	{ID: 2, Username: "admin1", Password: "12345", AccessLevel: "Admin"},
}