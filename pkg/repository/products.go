package repository

import "rest.com/pkg/model"

var products = []model.Product{
	{ID: 1, Name: "Телефон", Description: "Смартфон с 5G", Price: 50000, Quantity: 100},
	{ID: 2, Name: "Ноутбук", Description: "Ультрабук для работы", Price: 75000, Quantity: 50},
	{ID: 3, Name: "Планшет", Description: "Компактный и удобный", Price: 30000, Quantity: 60},
}


func GetProducts() *[]model.Product {
	return &products;
}