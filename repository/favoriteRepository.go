package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func AddToFavorites(ctx context.Context, userID, productID int64) error {
	query := `INSERT OR IGNORE INTO user_favorites (user_id, product_id) VALUES (?, ?)`
	_, err := db.DB.ExecContext(ctx, query, userID, productID)
	return err
}

func RemoveFromFavorites(ctx context.Context, userID, productID int64) error {
	query := `DELETE FROM user_favorites WHERE user_id = ? AND product_id = ?`
	_, err := db.DB.ExecContext(ctx, query, userID, productID)
	return err
}

func GetFavorites(ctx context.Context, userID int64) ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.rating, p.image_url, p.main_category, p.discount_price, p.is_featured, p.is_recommended
		FROM products p
		JOIN user_favorites uf ON p.id = uf.product_id
		WHERE uf.user_id = ?
	`
	rows, err := db.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Rating, &p.ImageURL, &p.MainCategory, &p.DiscountPrice, &p.IsFeatured, &p.IsRecommended); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
