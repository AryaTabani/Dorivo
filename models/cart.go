package models

type AddToCartPayload struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	OptionIDs []int64 `json:"option_ids"`
}

type CartItemOption struct {
	Name          string  `json:"name"`
	PriceModifier float64 `json:"price_modifier"`
}
type CartItem struct {
	ID         string           `json:"id"`
	ProductID  string           `json:"product_id"`
	Name       string           `json:"name"`
	ImageURL   string           `json:"image_url"`
	Quantity   int              `json:"quantity"`
	BasePrice  float64          `json:"base_price"`
	TotalPrice float64          `json:"total_price"`
	Options    []CartItemOption `json:"options"`
}
type Cart struct {
	Items      []CartItem `json:"items"`
	GrandTotal float64    `json:"grand_total"`
}
