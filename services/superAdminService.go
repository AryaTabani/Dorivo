package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidSuperAdminCredentials = errors.New("invalid email or password for super admin")

func LoginSuperAdmin(ctx context.Context, payload *models.SuperAdminLoginPayload) (string, error) {
	sa, err := repository.GetSuperAdminByEmail(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrInvalidSuperAdminCredentials
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(sa.PasswordHash), []byte(payload.Password))
	if err != nil {
		return "", ErrInvalidSuperAdminCredentials
	}

	return generateSuperAdminToken(sa.ID)
}

func generateSuperAdminToken(superAdminID int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": superAdminID,
		"rol": "SUPER_ADMIN",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func CreateTenant(ctx context.Context, name string, config *models.TenantConfig) error {
	return repository.CreateTenant(ctx, name, config)
}

func GetAllTenants(ctx context.Context) ([]models.Tenant, error) {
	return repository.GetAllTenants(ctx)
}

func DeleteTenant(ctx context.Context, tenantID string) error {
	return repository.DeleteTenant(ctx, tenantID)
}
