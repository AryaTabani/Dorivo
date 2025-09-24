package models

type PaymentMethod struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"-"`
	ProcessorToken string `json:"-"`
	CardBrand      string `json:"card_brand"`
	LastFour       string `json:"last_four"`
	ExpiryMonth    int    `json:"expiry_month"`
	ExpiryYear     int    `json:"expiry_year"`
}

type AddPaymentMethodPayload struct {
	ProcessorToken string `json:"processor_token" binding:"required"`
}
