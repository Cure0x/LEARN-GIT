package models

// "gorm.io/gorm"

type Order struct {
	// gorm.Model `swaggerignore:"true"`
	OrderId    int64 `gorm:"unique;primaryKey;autoIncrement" swaggerignore:"true"`
	CustomerId int64 `json:"customer_id"`
	SellerId   int64 `json:"seller_id"`
	ProductId  int64 `json:"product_id"`
	OrderTotal int64 `json:"order_total"`
}
