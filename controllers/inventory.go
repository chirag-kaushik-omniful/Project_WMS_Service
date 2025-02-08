package controllers

import (
	"net/http"
	"wms/models"
	"wms/utils/dbconn"

	"github.com/gin-gonic/gin"
)

func ViewInventory(c *gin.Context) {
	var inventory []models.Inventory
	if err := dbconn.DB_Instance.Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch inventory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": inventory})
}

func EditInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := dbconn.DB_Instance.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to edit inventory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}
