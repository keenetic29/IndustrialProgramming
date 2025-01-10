package repository

import (
	"rest.com/pkg/db"
	"rest.com/pkg/model"
)



func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := db.DB.Find(&products).Error
	return products, err
}

func GetProducts(limit int, page int, searchString string, sort string) ([]model.Product, error) {
	var products []model.Product

	offset := (page - 1) * limit
	query := db.DB.Limit(limit).Offset(offset)
	if searchString != "" {
		query = query.Where("name ILIKE ?", "%"+searchString+"%")
	}
	if sort != "" {
		query = query.Order(sort + " DESC")
	}

	err := query.Find(&products).Error
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