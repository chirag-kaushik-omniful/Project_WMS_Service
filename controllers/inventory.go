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

func UpdateInventory(c *gin.Context) {
	var req struct {
		ItemID   uint `json:"item_id"`
		Quantity int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var inventory models.Inventory
	if err := dbconn.DB_Instance.Where("sk_uid = ?", req.ItemID).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	if inventory.Quantity < req.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	inventory.Quantity -= req.Quantity
	if err := dbconn.DB_Instance.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_success":  true,
		"status_code": 200,
		"data": gin.H{
			"message": "Inventory updated successfully",
			"status":  string(http.StatusOK)},
		"meta": gin.H{}})
}
