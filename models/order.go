package models

import "time"

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
	Reason string `json:"reason" binding:"required"`
}
