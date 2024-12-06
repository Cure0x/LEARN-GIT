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
		api.POST("/product-batch", handlers.CreateBatchProduct)
		api.GET("/product-bycategory/:category", handlers.GetProductByCategory)
		api.GET("/product-byname/:name", handlers.GetProductByName)
		api.GET("/product-all/", handlers.GetAllProduct)

		//Order
		api.POST("/order", handlers.CreateOrder)
	}
}
