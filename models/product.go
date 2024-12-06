package models

// "gorm.io/gorm"

type Product struct {
	// gorm.Model `swaggerignore:"true"`
	ProductId int64   `gorm:"unique;primaryKey;autoIncrement" swaggerignore:"true"`
	Category  string  `json:"category"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Order     Order   `gorm:"foreignKey:ProductId" swaggerignore:"true"`
}
