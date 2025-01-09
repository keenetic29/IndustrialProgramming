package repository

import (
	"rest.com/pkg/db"
	"rest.com/pkg/model"
)



func GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := db.DB.Find(&products).Error
	return products, err
}

func AddProducts(products []model.Product) error {
	return db.DB.Create(products).Error
}

func GetProductById(id string) (model.Product, error) {
	var product model.Product
	err := db.DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func DeleteProduct(prod model.Product) error {
	return db.DB.Delete(&prod).Error
}

func UpdateProduct(prod model.Product) error {
	return db.DB.Save(prod).Error
}