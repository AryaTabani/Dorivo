package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func CreateAddress(ctx context.Context, userID int64, payload *models.AddAddressPayload) error {
	query := `INSERT INTO user_addresses (user_id, name, address) VALUES (?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, userID, payload.Name, payload.Address)
	return err
}

func GetAddressesByUserID(ctx context.Context, userID int64) ([]models.Address, error) {
	query := `SELECT id, name, address FROM user_addresses WHERE user_id = ? ORDER BY id DESC`
	rows, err := db.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var addr models.Address
		if err := rows.Scan(&addr.ID, &addr.Name, &addr.Address); err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}
	return addresses, nil
}

func DeleteAddress(ctx context.Context, userID int64, addressID int64) (int64, error) {
	query := `DELETE FROM user_addresses WHERE id = ? AND user_id = ?`
	result, err := db.DB.ExecContext(ctx, query, addressID, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
