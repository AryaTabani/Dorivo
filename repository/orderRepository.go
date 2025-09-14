package repository

import (
	"context"
	"database/sql"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func BeginTx(ctx context.Context) (*sql.Tx, error) {
	return db.DB.BeginTx(ctx, nil)
}

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

func GetOrderByIdAndUserID(ctx context.Context, orderID, userID int64) (*models.Order, error) {
	var order models.Order
	query := `SELECT id, user_id, tenant_id, status, total_price, created_at FROM orders WHERE id = ? AND user_id = ?`
	err := db.DB.QueryRowContext(ctx, query, orderID, userID).Scan(
		&order.ID,
		&order.UserID,
		&order.TenantID,
		&order.Status,
		&order.TotalPrice,
		&order.CreatedAt,
	)
	return &order, err
}

func CreateCancellation(ctx context.Context, tx *sql.Tx, userID, orderID int64, reason string) error {
	query := `INSERT INTO cancellations (order_id, user_id, reason) VALUES (?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, orderID, userID, reason)
	return err
}

func UpdateOrderStatus(ctx context.Context, tx *sql.Tx, orderID int64, newStatus string) error {
	query := `UPDATE orders SET status = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, query, newStatus, orderID)
	return err
}

func CheckIfReviewExists(ctx context.Context, orderID int64) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM reviews WHERE order_id = ?)`
	err := db.DB.QueryRowContext(ctx, query, orderID).Scan(&exists)
	return exists, err
}

func CreateReview(ctx context.Context, userID, orderID int64, rating int, comment string) error {
	query := `INSERT INTO reviews (order_id, user_id, rating, comment) VALUES (?, ?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, orderID, userID, rating, comment)
	return err
}
