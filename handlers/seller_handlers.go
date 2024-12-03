package handlers

import (
	"another-golang-crud/crud"
	"another-golang-crud/models"
	"net/http"

	//strconv

	"github.com/gin-gonic/gin"
)

// @Summary Create a new Seller
// @Tags Seller
// @Accept json
// @Produce json
// @param seller body models.Seller true "Seller data"
// @Success 201 {object} models.Seller
// @router /seller [post]
func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSeller, err := crud.CreateSeller(seller)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newSeller)
}
