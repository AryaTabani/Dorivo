package models

type User struct {
	ID                      int64                   `json:"id"`
	TenantID                string                  `json:"-"`
	Role                    string                  `json:"role"`
	Full_name               string                  `json:"full_name"`
	Email                   string                  `json:"email"`
	Mobile_number           string                  `json:"mobile_number"`
	Date_of_birth           string                  `json:"date_of_birth"`
	Avatar_url              string                  `json:"avatar_url"`
	NotificationPreferences NotificationPreferences `json:"notification_preferences"`
	Password_hash           string                  `json:"-"`
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
type UpdateProfilePayload struct {
	Full_name     string `json:"full_name" binding:"required"`
	Mobile_number string `json:"mobile_number"`
	Date_of_birth string `json:"date_of_birth"`
	Avatar_url    string `json:"avatar_url"`
}

type NotificationPreferences struct {
	GeneralNotifications bool `json:"general_notifications"`
	Sound                bool `json:"sound"`
	Vibrate              bool `json:"vibrate"`
	SpecialOffers        bool `json:"special_offers"`
	Payments             bool `json:"payments"`
	Cashback             bool `json:"cashback"`
}

type ChangePasswordPayload struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}
