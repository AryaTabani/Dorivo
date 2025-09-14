package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var (
	ErrOrderNotFound          = errors.New("order not found")
	ErrForbidden              = errors.New("you do not have permission to access this resource")
	ErrOrderCannotBeCancelled = errors.New("this order cannot be cancelled")
	ErrOrderNotCompleted      = errors.New("a review can only be left for a completed order")
	ErrReviewExists           = errors.New("a review for this order already exists")
)

func GetMyOrders(ctx context.Context, userID int64, status string) ([]models.OrderSummaryView, error) {
	if status == "" {
		status = "Active"
	}
	return repository.GetOrdersByUserID(ctx, userID, status)
}

func CancelOrder(ctx context.Context, userID, orderID int64, reason string) error {
	order, err := repository.GetOrderByIdAndUserID(ctx, orderID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrOrderNotFound
		}
		return err
	}

	if order.Status != "Active" {
		return ErrOrderCannotBeCancelled
	}

	tx, err := repository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := repository.CreateCancellation(ctx, tx, userID, orderID, reason); err != nil {
		return err
	}

	if err := repository.UpdateOrderStatus(ctx, tx, orderID, "Cancelled"); err != nil {
		return err
	}

	return tx.Commit()
}

func LeaveReview(ctx context.Context, userID, orderID int64, payload *models.LeaveReviewPayload) error {
	order, err := repository.GetOrderByIdAndUserID(ctx, orderID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrOrderNotFound
		}
		return err
	}

	if order.Status != "Completed" {
		return ErrOrderNotCompleted
	}

	exists, err := repository.CheckIfReviewExists(ctx, orderID)
	if err != nil {
		return err
	}
	if exists {
		return ErrReviewExists
	}

	return repository.CreateReview(ctx, userID, orderID, payload.Rating, payload.Comment)
}
