package services

import (
	"context"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func AddToFavorites(ctx context.Context, userID, productID int64) error {
	return repository.AddToFavorites(ctx, userID, productID)
}

func RemoveFromFavorites(ctx context.Context, userID, productID int64) error {
	return repository.RemoveFromFavorites(ctx, userID, productID)
}

func GetFavorites(ctx context.Context, userID int64) ([]models.Product, error) {
	return repository.GetFavorites(ctx, userID)
}
