package models

import "time"

type Order struct {
	ID         int64
	UserID     int64
	TenantID   string
	Status     string
	TotalPrice float64
	CreatedAt  time.Time
}

type OrderSummaryView struct {
	ID              int64     `json:"id"`
	TotalPrice      float64   `json:"total_price"`
	ItemCount       int       `json:"item_count"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	PrimaryItemName string    `json:"primary_item_name"`
	PrimaryItemImg  string    `json:"primary_item_img"`
}

type CancelOrderPayload struct {
	Reason string `json:"reason" binding:"max=255"`
}
