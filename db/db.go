package db

import (
	"another-golang-crud/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Inisialisasi koneksi database
var db *gorm.DB

// Fungsi untuk inisialisasi koneksi ke database
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Auto migrate atau setup database schema
	// Misalnya, migrasi model User dan Customer
	new := db.AutoMigrate(&models.Customer{}, &models.Seller{}, &models.Product{}, &models.Order{})
	if new != nil {
		panic(new)
	}
}

// Fungsi untuk mengakses koneksi db di tempat lain
func GetDB() *gorm.DB {
	return db
}
