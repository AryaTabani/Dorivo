package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func GetSuperAdminByEmail(ctx context.Context, email string) (*models.SuperAdmin, error) {
	var sa models.SuperAdmin
	query := `SELECT id, email, password_hash FROM super_admins WHERE email = ?`
	err := db.DB.QueryRowContext(ctx, query, email).Scan(&sa.ID, &sa.Email, &sa.PasswordHash)
	return &sa, err
}
