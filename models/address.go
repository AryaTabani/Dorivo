package models

type Address struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"-"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type AddAddressPayload struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
