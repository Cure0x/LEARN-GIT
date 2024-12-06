package handlers

import (
	"another-golang-crud/crud"
	"another-golang-crud/models"
	"fmt"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {object} models.Product
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct, err := crud.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// @Summary Create a new Batch Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body []models.Product true "Product data"
// @Success 201 {object} []models.Product
// @Router /product-batch [post]
func CreateBatchProduct(c *gin.Context) {
	var product []models.Product
	fmt.Println(product)
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(product)
	newProduct, err := crud.CreateBatchProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// @Summary Get product by category
// @Tags Product
// @Accept json
// @Produce json
// @Param category path string true "Category name"
// @Success 201 {array} models.Product
// @Router /product-bycategory/{category} [get]
func GetProductByCategory(c *gin.Context) {
	category := c.Param("category")
	ProductByCategory, err := crud.GetProductByCategory(category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ProductByCategory)
}

// @Summary Get product by name
// @Tags Product
// @Accept json
// @Produce json
// @Param name path string true "Product name"
// @Success 201 {array} models.Product
// @Router /product-byname/{name} [get]
func GetProductByName(c *gin.Context) {
	name := c.Param("name")
	ProductByCategory, err := crud.GetProductByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ProductByCategory)
}

// @Summary Get All products
// @Tags Product
// @Accept json
// @Produce json
// @Success 201 {array} models.Product
// @Router /product-all/ [get]
func GetAllProduct(c *gin.Context) {
	products, err := crud.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
