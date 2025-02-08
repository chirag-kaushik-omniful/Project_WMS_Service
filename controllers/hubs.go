package controllers

import (
	"net/http"
	"wms/models"
	"wms/utils/dbconn"

	"github.com/gin-gonic/gin"
)

func ViewHubs(c *gin.Context) {
	var hubs []models.Hub
	if err := dbconn.DB_Instance.Find(&hubs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch hubs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hubs})
}

func CreateHub(c *gin.Context) {

}
