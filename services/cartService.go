package services

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func AddToCart(ctx context.Context, userID int64, payload *models.AddToCartPayload) error {
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	cartID, err := repository.FindOrCreateCartByUserID(ctx, tx, userID)
	if err != nil {
		return err
	}

	if err := repository.AddItem(ctx, tx, cartID, payload); err != nil {
		return err
	}

	return tx.Commit()
}

func GetCart(ctx context.Context, userID int64) (*models.Cart, error) {
	items, err := repository.GetCartContentsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var grandTotal float64
	for i, item := range items {
		var itemOptionsTotal float64
		for _, opt := range item.Options {
			itemOptionsTotal += opt.PriceModifier
		}
		singleItemPrice := item.BasePrice + itemOptionsTotal
		items[i].TotalPrice = singleItemPrice * float64(item.Quantity)
		grandTotal += items[i].TotalPrice
	}

	cart := &models.Cart{
		Items:      items,
		GrandTotal: grandTotal,
	}

	return cart, nil
}
