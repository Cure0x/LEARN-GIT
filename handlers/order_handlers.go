package handlers

import (
	"another-golang-crud/crud"
	"another-golang-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new Order
// @Tags Order
// @Accept json
// @Produce json
// @param order body models.Order true "order data"
// @Success 201 {object} models.Order
// @Router /order [post]
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder, err := crud.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newOrder)
}
