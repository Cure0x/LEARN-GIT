package routes

import (
	"another-golang-crud/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		//Customer
		api.POST("/customer", handlers.CreateCustomer)
		api.GET("/customer-all/", handlers.GetAllCustomer)
		api.GET("/customer-all-product/:id", handlers.GetCustomerAllOrder)

		//Seller
		api.POST("/seller", handlers.CreateSeller)

		//Product
		api.POST("/product", handlers.CreateProduct)
		api.GET("/product-bycategory/:name", handlers.GetProductByCategory)
		api.GET("/product-all/", handlers.GetAllProduct)

		//Order
		api.POST("/order", handlers.CreateOrder)
	}
}
