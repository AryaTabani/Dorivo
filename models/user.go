package models

type User struct {
	ID            int64  `json:"id"`
	TenantID      int64  `json:"-"`
	Full_name     string `json:"full_name"`
	Email         string `json:"email"`
	Mobile_number string `json:"mobile_number"`
	Date_of_birth string `json:"date_of_birth"`
	Password_hash string `json:"-"`
}

type RegisterPayload struct {
	Full_name     string `json:"full_name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Mobile_number string `json:"mobile_number"`
	Date_of_birth string `json:"date_of_birth"`
	Password      string `json:"password" binding:"required,min=8"`
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResetPasswordPayload struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}