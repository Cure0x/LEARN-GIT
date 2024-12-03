package crud

import (
	// "errors"

	"another-golang-crud/db"
	"another-golang-crud/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	if err := db.GetDB().Create(&order).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}
