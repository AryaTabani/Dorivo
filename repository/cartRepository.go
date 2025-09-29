package repository

import (
	"context"
	"database/sql"
	"strconv"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func FindOrCreateCartByUserID(ctx context.Context, tx *sql.Tx, userID int64) (int64, error) {
	var cartID int64
	query := `SELECT id FROM carts WHERE user_id = ?`
	err := tx.QueryRowContext(ctx, query, userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		createQuery := `INSERT INTO carts (user_id) VALUES (?)`
		res, err := tx.ExecContext(ctx, createQuery, userID)
		if err != nil {
			return 0, err
		}
		cartID, err = res.LastInsertId()
		if err != nil {
			return 0, err
		}
		return cartID, nil
	}
	return cartID, err
}

func AddItem(ctx context.Context, tx *sql.Tx, cartID int64, payload *models.AddToCartPayload) error {

	itemQuery := `INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?)`
	res, err := tx.ExecContext(ctx, itemQuery, cartID, payload.ProductID, payload.Quantity)
	if err != nil {
		return err
	}
	cartItemID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if len(payload.OptionIDs) > 0 {
		optionsQuery := `INSERT INTO cart_item_options (cart_item_id, option_id) VALUES `
		var args []interface{}
		for _, optionID := range payload.OptionIDs {
			optionsQuery += `(?, ?),`
			args = append(args, cartItemID, optionID)
		}
		optionsQuery = optionsQuery[:len(optionsQuery)-1]
		_, err = tx.ExecContext(ctx, optionsQuery, args...)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetCartContentsByUserID(ctx context.Context, userID int64) ([]models.CartItem, error) {
	query := `
		SELECT ci.id, ci.product_id, ci.quantity, p.name, p.image_url, p.price, o.name, o.price_modifier
		FROM carts c
		JOIN cart_items ci ON c.id = ci.cart_id
		JOIN products p ON ci.product_id = p.id
		LEFT JOIN cart_item_options cio ON ci.id = cio.cart_item_id
		LEFT JOIN options o ON cio.option_id = o.id
		WHERE c.user_id = ?
		ORDER BY ci.id
	`
	rows, err := db.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cartItemsMap := make(map[int64]*models.CartItem)
	for rows.Next() {
		var itemID, productID int64
		var quantity int
		var productName, productImageURL string
		var basePrice float64
		var optionName sql.NullString
		var optionPrice sql.NullFloat64

		if err := rows.Scan(&itemID, &productID, &quantity, &productName, &productImageURL, &basePrice, &optionName, &optionPrice); err != nil {
			return nil, err
		}
		if _, ok := cartItemsMap[itemID]; !ok {
			cartItemsMap[itemID] = &models.CartItem{
				ID:        strconv.FormatInt(itemID, 10),
				ProductID: strconv.FormatInt(productID, 10),
				Quantity:  quantity,
				Name:      productName,
				ImageURL:  productImageURL,
				BasePrice: basePrice,
				Options:   make([]models.CartItemOption, 0),
			}
		}

		if optionName.Valid {
			cartItemsMap[itemID].Options = append(cartItemsMap[itemID].Options, models.CartItemOption{
				Name:          optionName.String,
				PriceModifier: optionPrice.Float64,
			})
		}
	}

	var items []models.CartItem
	for _, item := range cartItemsMap {
		items = append(items, *item)
	}

	return items, nil
}
