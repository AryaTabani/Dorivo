package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	db "example.com/m/v2/DB"
	"example.com/m/v2/models"
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