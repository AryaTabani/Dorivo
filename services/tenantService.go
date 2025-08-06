package services

import (
	"context"
	"errors"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var ErrTenantNotFound = errors.New("tenant not found")

func GetTenantConfig(ctx context.Context, id string) (*models.TenantConfig, error) {
	tenant, err := repository.GetTenantByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}

	return &tenant.Config, nil
}
