package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func GetFAQsByTenant(ctx context.Context, tenantID string, category string) ([]models.FAQ, error) {
	args := []interface{}{tenantID}
	query := "SELECT id, category, question, answer FROM faqs WHERE tenant_id = ?"
	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}
	rows, err := db.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var faqs []models.FAQ
	for rows.Next() {
		var faq models.FAQ
		if err := rows.Scan(&faq.ID, &faq.Category, &faq.Question, &faq.Answer); err != nil {
			return nil, err
		}
		faqs = append(faqs, faq)
	}
	return faqs, nil
}
