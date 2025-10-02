package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (tenant_id, full_name, email, mobile_number, password_hash, date_of_birth) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, user.TenantID, user.Full_name, user.Email, user.Mobile_number, user.Password_hash, user.Date_of_birth)
	return err
}

func GetUserByEmailAndTenant(ctx context.Context, email string, tenantID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, full_name, email, mobile_number, password_hash, date_of_birth, tenant_id,role FROM users WHERE email = ? AND tenant_id = ?`
	err := db.DB.QueryRowContext(ctx, query, email, tenantID).Scan(
		&user.ID,
		&user.Full_name,
		&user.Email,
		&user.Mobile_number,
		&user.Password_hash,
		&user.Date_of_birth,
		&user.TenantID,
		&user.Role,
	)
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
func UpdateUser(ctx context.Context, userID int64, payload *models.UpdateProfilePayload) error {
	query := `UPDATE users SET full_name = ?, mobile_number = ?, date_of_birth = ?, avatar_url = ? WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, payload.Full_name, payload.Mobile_number, payload.Date_of_birth, payload.Avatar_url, userID)
	return err
}

func GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	var user models.User
	var prefsJSON sql.NullString

	query := `SELECT id, full_name, email, mobile_number, date_of_birth, avatar_url, tenant_id, password_hash, notification_preferences FROM users WHERE id = ?`
	err := db.DB.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Full_name,
		&user.Email,
		&user.Mobile_number,
		&user.Date_of_birth,
		&user.Avatar_url,
		&user.TenantID,
		&user.Password_hash,
		&prefsJSON,
	)
	if err != nil {
		return nil, err
	}
	if prefsJSON.Valid {
		err = json.Unmarshal([]byte(prefsJSON.String), &user.NotificationPreferences)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
func UpdateNotificationPreferences(ctx context.Context, userID int64, prefs *models.NotificationPreferences) error {
	prefsJSON, err := json.Marshal(prefs)
	if err != nil {
		return err
	}
	query := `UPDATE users SET notification_preferences = ? WHERE id = ?`
	_, err = db.DB.ExecContext(ctx, query, prefsJSON, userID)
	return err
}
func DeleteUserByID(ctx context.Context, userID int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := db.DB.ExecContext(ctx, query, userID)
	return err
}
