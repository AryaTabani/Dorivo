package models

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreateProductResponse struct {
	ProductID int64 `json:"product_id"`
}
