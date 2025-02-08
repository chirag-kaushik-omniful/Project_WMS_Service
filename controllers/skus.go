package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wms/models"
	"wms/utils/dbconn"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ViewSKUs(c *gin.Context) {
	var skus []models.SKU
	if err := dbconn.DB_Instance.Find(&skus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch SKUs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": skus})
}

func CreateSKU(c *gin.Context) {

}

type Order struct {
	SNo      string `json:"sno"`
	SellerID string `json:"seller_id"`
	OrderID  string `json:"order_id"`
	ItemID   string `json:"item_id"`
	Quantity string `json:"quantity"`
	Status   string `json:"status"`
}

type OrderResponse struct {
	ValidOrders   []Order `json:"valid_orders"`
	MissingOrders []Order `json:"missing_orders"`
}

type SKU struct {
	ID string `gorm:"primaryKey"`
}

func VerifySKUs(c *gin.Context) {
	fmt.Println("Going...")
	var orders []Order
	if err := c.BindJSON(&orders); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	db := dbconn.DB_Instance
	var validOrders []Order
	var missingSKUs []Order

	for _, order := range orders {
		var sku SKU
		if err := db.First(&sku, "id = ?", order.ItemID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				missingSKUs = append(missingSKUs, order)
			} else {
				log.Println("DB error:", err)
			}
		} else {
			validOrders = append(validOrders, order)
		}
	}
	fmt.Println(missingSKUs)
	fmt.Println(validOrders)

	orderdata := &OrderResponse{
		ValidOrders:   validOrders,
		MissingOrders: missingSKUs,
	}

	dt, _ := json.Marshal(orderdata)
	fmt.Println(string(dt))
	// if len(missingSKUs) > 0 {
	// 	writeMissingSKUsToCSV(missingSKUs)
	// }

	// c.JSON(200, gin.H{"valid_orders": validOrders})

	c.JSON(200, gin.H{
		"is_success":  true,
		"status_code": 200,
		"data":        gin.H{"orders": orderdata},
		"meta":        gin.H{}})
}

// func writeMissingSKUsToCSV(skus []string) {
// 	file, err := os.OpenFile("missing_skus.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Println("Failed to open CSV file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	for _, sku := range skus {
// 		if err := writer.Write([]string{sku}); err != nil {
// 			log.Println("Failed to write SKU to CSV:", err)
// 		}
// 	}
// }
