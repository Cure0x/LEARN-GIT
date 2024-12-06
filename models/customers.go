package models

// "gorm.io/gorm"

type Customer struct {
	// gorm.Model `swaggerignore:"true"`
	CustomerId int64   `gorm:"unique;primaryKey;autoIncrement" swaggerignore:"true"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	BirthDate  string  `json:"birth_date"`
	Address    string  `json:"address"`
	MoneySpent int64   `json:"money_spent" swaggerignore:"true"`
	Order      []Order `gorm:"foreignKey:CustomerId" swaggerignore:"true"`
}
