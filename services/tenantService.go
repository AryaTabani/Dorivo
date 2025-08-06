package services

import (
	"context"
	"errors"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
)

var ErrTenantNotFound = errors.New("tenant not found")

func GetTenantConfig(ctx context.Context, id int64) (*models.TenantConfig, error) {
	tenant, err := repository.GetTenantByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	return &tenant.Config, nil
}