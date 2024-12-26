package repository

import "rest.com/pkg/model"

var products = []model.Product{
	{ID: 1, Name: "Телефон", Description: "Смартфон с 5G", Cost: 50000, Count: 100, ManufacturerId: 1, SupplierId: 1},
	{ID: 2, Name: "Ноутбук", Description: "Ультрабук для работы", Cost: 75000, Count: 50, ManufacturerId: 1, SupplierId: 1},
	{ID: 3, Name: "Планшет", Description: "Компактный и удобный", Cost: 30000, Count: 60, ManufacturerId: 1, SupplierId: 1},
}


func GetProducts() *[]model.Product {
	return &products;
}