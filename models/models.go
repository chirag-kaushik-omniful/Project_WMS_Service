package models

import (
	"time"
)

// Tenant Model
type Tenant struct {
	ID        string    `json:"_id,omitempty" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Address Model
type Address struct {
	ID      string `json:"_id,omitempty" gorm:"primaryKey"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

// Hub Model
type Hub struct {
	ID        string    `json:"_id,omitempty" gorm:"primaryKey"`
	TenantID  string    `json:"tenant_id" gorm:"not null"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Foreign Key Constraint
	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID;constraint:OnDelete:CASCADE"`
}

// Seller Model
type Seller struct {
	ID        string    `json:"_id,omitempty" gorm:"primaryKey"`
	TenantID  string    `json:"tenant_id" gorm:"not null"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Status    string    `json:"status"` // "Active", "Inactive", "Pending"
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Foreign Key Constraint
	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID;constraint:OnDelete:CASCADE"`
}

// Product Model
type Product struct {
	ID          string    `json:"_id,omitempty" gorm:"primaryKey"`
	TenantID    string    `json:"tenant_id" gorm:"not null"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Foreign Key Constraint
	Tenant Tenant `gorm:"foreignKey:TenantID;references:ID;constraint:OnDelete:CASCADE"`
}

// SKU Model
type SKU struct {
	ID          string    `json:"_id,omitempty" gorm:"primaryKey"`
	ProductID   string    `json:"product_id" gorm:"not null"`
	SKUCode     string    `json:"sku_code"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Foreign Key Constraint
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

// Inventory Model
type Inventory struct {
	ID        string    `json:"_id,omitempty" gorm:"primaryKey"`
	HubID     string    `json:"hub_id" gorm:"not null"`
	SKUID     string    `json:"sku_id" gorm:"not null"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Foreign Key Constraints
	Hub Hub `gorm:"foreignKey:HubID;references:ID;constraint:OnDelete:CASCADE"`
	SKU SKU `gorm:"foreignKey:SKUID;references:ID;constraint:OnDelete:CASCADE"`
}
