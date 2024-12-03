package main

import (
	_ "another-golang-crud/docs" // Import Swagger docs
	"another-golang-crud/routes"

	// "another-golang-crud/crud"
	"another-golang-crud/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go CRUD API
// @version         1.0
// @description     Simple CRUD API with Swagger documentation.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	db.InitDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://127.0.0.1:8080", "http://127.0.0.1:8080", "http://192.168.1.6:8080", "https://192.168.1.6:8080", "http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	routes.RegisterRoutes(r)
	// crud.DebugRelasi()
	r.Run("0.0.0.0:8080")

}
