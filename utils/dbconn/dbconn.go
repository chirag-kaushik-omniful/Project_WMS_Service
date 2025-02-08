package dbconn

import (
	"context"
	"fmt"
	"log"
	"wms/models"

	"github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
)

var DB_Instance *gorm.DB
var DB_Ctx context.Context

func Connect(masterConfig postgres.DBConfig, slavesConfig *[]postgres.DBConfig) {
	conn := postgres.InitializeDBInstance(masterConfig, slavesConfig)
	ctx := context.Background()

	DB_Instance = conn.GetMasterDB(ctx)
	DB_Ctx = ctx

	if DB_Instance != nil {
		fmt.Println("Database Connected Successfully!")
		MigrateDB(DB_Instance)
	}
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Tenant{}, &models.Address{}, &models.Hub{}, &models.Seller{}, &models.Product{}, &models.SKU{}, &models.Inventory{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Database migration completed successfully!")
}
