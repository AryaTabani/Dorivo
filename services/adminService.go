package services

import (
	"context"
	"errors"
	"fmt"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func CreateProduct(ctx context.Context, tenantID string, payload *models.ProductPayload) (int64, error) {
	return repository.CreateProduct(ctx, tenantID, payload)
}

func UpdateProduct(ctx context.Context, tenantID string, productID int64, payload *models.ProductPayload) error {
	return repository.UpdateProduct(ctx, tenantID, productID, payload)
}

func DeleteProduct(ctx context.Context, tenantID string, productID int64) error {
	return repository.DeleteProduct(ctx, tenantID, productID)
}

func UpdateTenantConfig(ctx context.Context, tenantID string, payload *models.TenantConfig) error {
	err := repository.UpdateTenantConfig(ctx, tenantID, payload)
	if err != nil {
		return err
	}
	cacheKey := fmt.Sprintf("tenant_config:%s", tenantID)
	db.Rdb.Del(db.Ctx, cacheKey)

	return nil
}
func GetTenantOrders(ctx context.Context, tenantID string, status string) ([]models.Order, error) {
	if status == "" {
		status = "Active"
	}
	return repository.GetOrdersByTenantID(ctx, tenantID, status)
}
func UpdateOrderStatus(ctx context.Context, tenantID string, orderID int64, newStatus string) error {
	RowsAffected, err := repository.AdminUpdateOrderStatus(ctx, tenantID, orderID, newStatus)
	if err != nil {
		return err
	}
	if RowsAffected == 0 {
		return errors.New("order not found or you do not have permission to update it")
	}
	return nil
}
func GetTenantCustomers(ctx context.Context, tenantID string) ([]models.User, error) {
	return repository.GetUsersByTenantID(ctx, tenantID)
}
func GetDashboardStats(ctx context.Context, tenantID string) (*models.DashboardStats, error) {
	return repository.GetTenantDashboardStats(ctx, tenantID)
}
