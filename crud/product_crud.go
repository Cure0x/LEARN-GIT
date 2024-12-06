package crud

import (
	// "errors"

	"another-golang-crud/db"
	"another-golang-crud/models"
	"fmt"
)

func CreateProduct(product models.Product) (models.Product, error) {
	if err := db.GetDB().Create(&product).Error; err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func CreateBatchProduct(product []models.Product) ([]models.Product, error) {
	if err := db.GetDB().Create(&product).Error; err != nil {
		return []models.Product{}, err
	}
	return product, nil
}

func GetProductByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	if err := db.GetDB().Where("category = ?", category).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByName(name string) ([]models.Product, error) {
	var products []models.Product
	fmt.Println("%" + name + "%")
	if err := db.GetDB().Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetAllProduct() ([]models.Product, error) {
	var products []models.Product
	if err := db.GetDB().Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
