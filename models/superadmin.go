package models

type SuperAdmin struct {
	ID           int64
	Email        string
	PasswordHash string
}

type SuperAdminLoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
