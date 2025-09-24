package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func CreatePaymentMethod(ctx context.Context, method *models.PaymentMethod) error {
	query := `INSERT INTO payment_methods (user_id, processor_token, card_brand, last_four, expiry_month, expiry_year) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, method.UserID, method.ProcessorToken, method.CardBrand, method.LastFour, method.ExpiryMonth, method.ExpiryYear)
	return err
}

func GetPaymentMethodsByUserID(ctx context.Context, userID int64) ([]models.PaymentMethod, error) {
	query := `SELECT id, card_brand, last_four, expiry_month, expiry_year FROM payment_methods WHERE user_id = ?`
	rows, err := db.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []models.PaymentMethod
	for rows.Next() {
		var pm models.PaymentMethod
		if err := rows.Scan(&pm.ID, &pm.CardBrand, &pm.LastFour, &pm.ExpiryMonth, &pm.ExpiryYear); err != nil {
			return nil, err
		}
		methods = append(methods, pm)
	}
	return methods, nil
}

func DeletePaymentMethod(ctx context.Context, userID, methodID int64) (int64, error) {
	query := `DELETE FROM payment_methods WHERE id = ? AND user_id = ?`
	result, err := db.DB.ExecContext(ctx, query, methodID, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
