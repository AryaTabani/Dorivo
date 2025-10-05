package services

import (
	"context"
	"errors"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var ErrCartItemNotFound = errors.New("cart item not found or you do not have permission to modify it")

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

func UpdateCartItemQuantity(ctx context.Context, userID, itemID int64, quantity int) error {
	rowsAffected, err := repository.UpdateCartItemQuantity(ctx, userID, itemID, quantity)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrCartItemNotFound
	}
	return nil
}

func RemoveCartItem(ctx context.Context, userID, itemID int64) error {
	rowsAffected, err := repository.RemoveCartItem(ctx, userID, itemID)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrCartItemNotFound
	}
	return nil
}
