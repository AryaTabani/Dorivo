package repository

import (
	"context"
	"time"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (full_name, email, mobile_number, password_hash, date_of_birth) VALUES (?, ?, ?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, user.Full_name, user.Email, user.Mobile_number, user.Password_hash, user.Date_of_birth)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, full_name, email, mobile_number, password_hash, date_of_birth FROM users WHERE email = ?`
	err := db.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Full_name, &user.Email, &user.Mobile_number, &user.Password_hash, &user.Date_of_birth)
	return &user, err
}

func CreatePasswordResetToken(ctx context.Context, userID int64, tokenHash string, expiresAt time.Time) error {
	query := `INSERT INTO password_reset_tokens (user_id, token_hash, expires_at) VALUES (?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, userID, tokenHash, expiresAt)
	return err
}

func UpdateUserPassword(ctx context.Context, userID int64, newPasswordHash string) error {
	query := `UPDATE users SET password_hash = ? WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, newPasswordHash, userID)
	return err
}
