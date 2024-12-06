package models

// "gorm.io/gorm"

type Seller struct {
	// gorm.Model `swaggerignore:"true"`
	SellerId  int64  `gorm:"unique;primaryKey;autoIncrement" swaggerignore:"true"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Order     Order  `gorm:"foreignKey:SellerId" swaggerignore:"true"`
}
