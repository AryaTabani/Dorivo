package models

type ProductPayload struct {
	Name          string   `json:"name" binding:"required"`
	Description   string   `json:"description"`
	Price         float64  `json:"price" binding:"required"`
	ImageURL      string   `json:"image_url"`
	MainCategory  string   `json:"main_category" binding:"required"`
	DiscountPrice *float64 `json:"discount_price"`
	IsFeatured    bool     `json:"is_featured"`
	IsRecommended bool     `json:"is_recommended"`
}
type UpdateOrderStatusPayload struct {
	Status string `json:"status" binding:"required"`
}
