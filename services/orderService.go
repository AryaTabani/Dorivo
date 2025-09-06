package services

import (
	"context"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func GetMyOrders(ctx context.Context, userID int64, status string) ([]models.OrderSummaryView, error) {
	if status == "" {
		status = "active"
	}
	return repository.GetOrdersByUserID(ctx, userID, status)
}
