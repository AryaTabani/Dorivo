package repository

import (
	"context"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func GetTenantDashboardStats(ctx context.Context, tenantID string) (*models.DashboardStats, error) {
	var stats models.DashboardStats

	revenueQuery := "SELECT COALESCE(SUM(total_price), 0) FROM orders WHERE tenant_id = ? AND status = 'Completed'"
	err := db.DB.QueryRowContext(ctx, revenueQuery, tenantID).Scan(&stats.TotalRevenue)
	if err != nil {
		return nil, err
	}

	ordersTodayQuery := "SELECT COUNT(*) FROM orders WHERE tenant_id = ? AND created_at >= CURDATE()"
	err = db.DB.QueryRowContext(ctx, ordersTodayQuery, tenantID).Scan(&stats.OrdersToday)
	if err != nil {
		return nil, err
	}

	customersQuery := "SELECT COUNT(*) FROM users WHERE tenant_id = ? AND role = 'CUSTOMER'"
	err = db.DB.QueryRowContext(ctx, customersQuery, tenantID).Scan(&stats.TotalCustomers)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
