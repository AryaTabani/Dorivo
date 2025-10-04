package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func GetTenantByID(ctx context.Context, id string) (*models.Tenant, error) {
	var tenant models.Tenant
	var configJSON string

	query := "SELECT id, config FROM tenants WHERE name = ?"
	err := db.DB.QueryRowContext(ctx, query, id).Scan(&tenant.ID, &configJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal([]byte(configJSON), &tenant.Config)
	if err != nil {
		return nil, err
	}

	return &tenant, nil
}
func UpdateTenantConfig(ctx context.Context, tenantID string, config *models.TenantConfig) error {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return err
	}

	query := `UPDATE tenants SET config = ? WHERE name = ?`
	_, err = db.DB.ExecContext(ctx, query, string(configJSON), tenantID)
	return err
}

func CreateTenant(ctx context.Context, name string, config *models.TenantConfig) error {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return err
	}
	query := `INSERT INTO tenants (name, config) VALUES (?, ?)`
	_, err = db.DB.ExecContext(ctx, query, name, string(configJSON))
	return err
}

func GetAllTenants(ctx context.Context) ([]models.Tenant, error) {
	rows, err := db.DB.QueryContext(ctx, "SELECT name, config FROM tenants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []models.Tenant
	for rows.Next() {
		var t models.Tenant
		var configJSON string
		if err := rows.Scan(&t.Name, &configJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(configJSON), &t.Config); err != nil {
			return nil, err
		}
		tenants = append(tenants, t)
	}
	return tenants, nil
}

func DeleteTenant(ctx context.Context, tenantID string) error {
	_, err := db.DB.ExecContext(ctx, "DELETE FROM tenants WHERE name = ?", tenantID)
	return err
}
