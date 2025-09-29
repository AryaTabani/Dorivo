package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func SearchProducts(ctx context.Context, tenantID string, filters map[string][]string) ([]models.Product, error) {
	var args []interface{}
	var whereClauses []string

	baseQuery := `
		SELECT p.id, p.name, p.description, p.price, p.rating, p.image_url, p.main_category, GROUP_CONCAT(t.name) as tags
		FROM products p
		LEFT JOIN product_tags pt ON p.id = pt.product_id
		LEFT JOIN tags t ON pt.tag_id = t.id
	`
	whereClauses = append(whereClauses, "p.tenant_id = ?")
	args = append(args, tenantID)

	for key, values := range filters {
		if len(values) == 0 {
			continue
		}
		switch key {
		case "category":
			whereClauses = append(whereClauses, "p.main_category = ?")
			args = append(args, values[0])
		case "min_price":
			whereClauses = append(whereClauses, "p.price >= ?")
			args = append(args, values[0])
		case "max_price":
			whereClauses = append(whereClauses, "p.price <= ?")
			args = append(args, values[0])
		}
	}

	query := baseQuery
	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	query += " GROUP BY p.id"

	if tagValues, ok := filters["tags"]; ok && len(tagValues) > 0 {
		query += fmt.Sprintf(" HAVING SUM(CASE WHEN t.name IN (?%s) THEN 1 ELSE 0 END) = ?", strings.Repeat(",?", len(tagValues)-1))
		for _, tag := range tagValues {
			args = append(args, tag)
		}
		args = append(args, len(tagValues))
	}

	if sortBy, ok := filters["sort_by"]; ok && len(sortBy) > 0 {
		if sortBy[0] == "rating_desc" {
			query += " ORDER BY p.rating DESC"
		}
	}

	rows, err := db.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		var tags sql.NullString
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Rating, &p.ImageURL, &p.MainCategory, &tags); err != nil {
			return nil, err
		}
		if tags.Valid {
			p.Tags = strings.Split(tags.String, ",")
		}
		products = append(products, p)
	}

	return products, nil
}

func GetTags(ctx context.Context, tenantID string) ([]models.Tag, error) {
	query := "SELECT id, name, main_category FROM tags WHERE tenant_id = ?"
	rows, err := db.DB.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var t models.Tag
		if err := rows.Scan(&t.ID, &t.Name, &t.MainCategory); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, nil
}

func GetProductDetails(ctx context.Context, tenantID string, productID int64) (*models.Product, error) {
	var p models.Product
	productQuery := `SELECT id, name, description, price, rating, image_url, main_category FROM products WHERE id = ? AND tenant_id = ?`
	err := db.DB.QueryRowContext(ctx, productQuery, productID, tenantID).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Rating, &p.ImageURL, &p.MainCategory)
	if err != nil {
		return nil, err
	}

	optionsQuery := `
		SELECT og.id, og.name, og.selection_type, o.id, o.name, o.price_modifier
		FROM option_groups og
		JOIN options o ON og.id = o.option_group_id
		WHERE og.product_id = ?
		ORDER BY og.id, o.id
	`
	rows, err := db.DB.QueryContext(ctx, optionsQuery, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	optionGroupsMap := make(map[int64]*models.OptionGroup)
	for rows.Next() {
		var ogID int64
		var ogName, ogSelectionType string
		var opt models.Option
		if err := rows.Scan(&ogID, &ogName, &ogSelectionType, &opt.ID, &opt.Name, &opt.PriceModifier); err != nil {
			return nil, err
		}

		if _, ok := optionGroupsMap[ogID]; !ok {
			optionGroupsMap[ogID] = &models.OptionGroup{
				ID:            ogID,
				Name:          ogName,
				SelectionType: ogSelectionType,
				Options:       make([]models.Option, 0),
			}
		}
		optionGroupsMap[ogID].Options = append(optionGroupsMap[ogID].Options, opt)
	}

	p.OptionGroups = make([]models.OptionGroup, 0, len(optionGroupsMap))
	for _, group := range optionGroupsMap {
		p.OptionGroups = append(p.OptionGroups, *group)
	}

	return &p, nil
}
