package routes

import (
	"wms/controllers"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
)

func GetRouter(Router *http.Server) {
	api := Router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello"})
		})

		hub := api.Group("/hub")
		{
			hub.GET("/view", controllers.ViewHubs)
			hub.POST("/create", controllers.CreateHub)
		}

		sku := api.Group("/sku")
		{
			sku.GET("/view", controllers.ViewSKUs)
			sku.POST("/create", controllers.CreateSKU)
			sku.POST("/verify", controllers.VerifySKUs)
		}

		inventory := api.Group("/inventory")
		{
			inventory.GET("/view", controllers.ViewInventory)
			inventory.POST("/edit", controllers.EditInventory)
		}
	}
}
