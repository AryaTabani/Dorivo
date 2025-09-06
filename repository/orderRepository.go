package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func GetOrdersByUserID(ctx context.Context, userID int64, status string) ([]models.OrderSummaryView, error) {
	query := `
		SELECT o.id, o.total_price, o.status, o.created_at,
			(SELECT item_name FROM order_items WHERE order_id = o.id LIMIT 1) as primary_item_name,
			(SELECT image_url FROM order_items WHERE order_id = o.id LIMIT 1) as primary_item_img,
			(SELECT COUNT(*) FROM order_items WHERE order_id = o.id) as item_count
		FROM orders o
		WHERE o.user_id = ? AND o.status = ?
		ORDER BY o.created_at DESC;
	`
	rows, err := db.DB.QueryContext(ctx, query, userID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.OrderSummaryView
	for rows.Next() {
		var order models.OrderSummaryView
		err := rows.Scan(&order.ID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.PrimaryItemName, &order.PrimaryItemImg, &order.ItemCount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
