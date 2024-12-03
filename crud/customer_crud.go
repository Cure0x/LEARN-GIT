package crud

import (
	// "errors"

	"another-golang-crud/db"
	"another-golang-crud/models"
)

func CreateCustomer(customer models.Customer) (models.Customer, error) {
	if err := db.GetDB().Create(&customer).Error; err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

func GetAllCustomer() ([]models.Customer, error) {
	var customer []models.Customer
	if err := db.GetDB().Find(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func GetCustomerAllOrder(customer_id uint) (models.Customer, error) {
	var customer models.Customer
	err := db.GetDB().Preload("Order").Where("customer_id =?", customer_id).First(&customer).Error
	if err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}
