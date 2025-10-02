package models

type DashboardStats struct {
	TotalRevenue   float64 `json:"total_revenue"`
	OrdersToday    int     `json:"orders_today"`
	TotalCustomers int     `json:"total_customers"`
}
