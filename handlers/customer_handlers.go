package handlers

import (
	"another-golang-crud/crud"
	"another-golang-crud/models"
	"fmt"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.Customer true "Customer data"
// @Success 201 {object} models.Customer
// @Router /customer [post]
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCustomer, err := crud.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newCustomer)
}

// @Summary Get All Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Success 201 {array} models.Customer
// @Router /customer-all/ [get]
func GetAllCustomer(c *gin.Context) {
	customers, err := crud.GetAllCustomer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// @Summary Get Customer all Orders
// @Tags Customer
// @Accept json
// @Produce json
// @param id path int true "customer id"
// @Success 201 {object} models.Customer
// @Router /customer-all-product/{id} [get]
func GetCustomerAllOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	customers, err := crud.GetCustomerAllOrder(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}
