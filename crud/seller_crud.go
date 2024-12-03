package crud

import (
	// "errors"

	"another-golang-crud/db"
	"another-golang-crud/models"
)

func CreateSeller(seller models.Seller) (models.Seller, error) {
	if err := db.GetDB().Create(&seller).Error; err != nil {
		return models.Seller{}, err
	}
	return seller, nil
}
