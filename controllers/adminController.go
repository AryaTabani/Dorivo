package controllers

import (
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// CreateProductHandler godoc
// @Summary      Create a new product
// @Description  Allows a tenant admin to create a new product for their store.
// @Tags         Admin Panel - Product Management
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string                true "Tenant ID"
// @Param        product  body     models.ProductPayload true "New Product Data"
// @Success      201      {object} models.APIResponse[models.CreateProductResponse] "Product created successfully"
// @Failure      400      {object} models.APIResponse[any] "Invalid request body"
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      500      {object} models.APIResponse[any] "Failed to create product"
// @Router       /{tenantId}/admin/products [post]
func CreateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantId := c.Param("tenantId")

		var payload models.ProductPayload
		err := c.ShouldBindJSON(&payload)

		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Invalid request body",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		productID, err := services.CreateProduct(c.Request.Context(), tenantId, &payload)
		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Could not create Product",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[models.CreateProductResponse]{
			Success: true,
			Data:    models.CreateProductResponse{ProductID: productID},
		}
		c.JSON(http.StatusOK, response)
	}
}

// UpdateProductHandler godoc
// @Summary      Update an existing product
// @Description  Allows a tenant admin to update the details of an existing product.
// @Tags         Admin Panel - Product Management
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId  path     string                true "Tenant ID"
// @Param        productId path     int                   true "Product ID"
// @Param        product   body     models.ProductPayload true "Updated Product Data"
// @Success      201      {object} models.APIResponse[models.CreateProductResponse] "Product created successfully"
// @Failure      400       {object} models.APIResponse[any] "Invalid request body"
// @Failure      403       {object} models.APIResponse[any] "Forbidden"
// @Failure      500       {object} models.APIResponse[any] "Failed to update product"
// @Router       /{tenantId}/admin/products/{productId} [put]
func UpdateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)

		var payload models.ProductPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}

		err := services.UpdateProduct(c.Request.Context(), tenantID, productID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Product updated successfully"})
	}
}

// DeleteProductHandler godoc
// @Summary      Delete a product
// @Description  Allows a tenant admin to delete a product from their store.
// @Tags         Admin Panel - Product Management
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId  path     string true "Tenant ID"
// @Param        productId path     int    true "Product ID"
// @Success      200       {object} models.APIResponse[any] "Product deleted successfully"
// @Failure      403       {object} models.APIResponse[any] "Forbidden"
// @Failure      500       {object} models.APIResponse[any] "Failed to delete product"
// @Router       /{tenantId}/admin/products/{productId} [delete]
func DeleteProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)

		err := services.DeleteProduct(c.Request.Context(), tenantID, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Product deleted successfully"})
	}
}

// UpdateTenantConfigHandler godoc
// @Summary      Update tenant configuration
// @Description  Allows a tenant admin to update their own store's configuration (e.g., name, theme, contact info).
// @Tags         Admin Panel - Configuration
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string              true "Tenant ID"
// @Param        config   body     models.TenantConfig true "Updated Tenant Configuration"
// @Success      200      {object} models.APIResponse[any] "Tenant configuration updated successfully"
// @Failure      400      {object} models.APIResponse[any] "Invalid request body"
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      500      {object} models.APIResponse[any] "Failed to update tenant configuration"
// @Router       /{tenantId}/admin/config [put]
func UpdateTenantConfigHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")

		var payload models.TenantConfig
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err := services.UpdateTenantConfig(c.Request.Context(), tenantID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to update tenant configuration"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Tenant configuration updated successfully"})
	}
}

// GetTenantOrdersHandler godoc
// @Summary      Get all orders for the tenant
// @Description  Allows a tenant admin to view all orders placed for their store, filterable by status.
// @Tags         Admin Panel - Order Management
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string true "Tenant ID"
// @Param        status   query    string false "Filter orders by status (e.g., Active, Completed, Cancelled)"
// @Success      200      {object} models.APIResponse[[]models.Order]
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve orders"
// @Router       /{tenantId}/admin/orders [get]
func GetTenantOrdersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		status := c.Query("status")

		orders, err := services.GetTenantOrders(c.Request.Context(), tenantID, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve orders"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[[]models.Order]{Success: true, Data: orders})
	}
}

// UpdateOrderStatusHandler godoc
// @Summary      Update an order's status
// @Description  Allows a tenant admin to update the status of a specific order (e.g., to 'Preparing', 'Completed').
// @Tags         Admin Panel - Order Management
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string                         true "Tenant ID"
// @Param        orderId  path     int                            true "Order ID"
// @Param        status   body     models.UpdateOrderStatusPayload true "New Order Status"
// @Success      200      {object} models.APIResponse[any] "Order status updated successfully"
// @Failure      400      {object} models.APIResponse[any] "Invalid request body"
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      404      {object} models.APIResponse[any] "Order not found"
// @Router       /{tenantId}/admin/orders/{orderId}/status [put]
func UpdateOrderStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		orderID, _ := strconv.ParseInt(c.Param("orderId"), 10, 64)

		var payload models.UpdateOrderStatusPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}

		err := services.UpdateOrderStatus(c.Request.Context(), tenantID, orderID, payload.Status)
		if err != nil {
			c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Order status updated successfully"})
	}
}

// GetTenantCustomersHandler godoc
// @Summary      Get all customers for the tenant
// @Description  Allows a tenant admin to view a list of all customers who have registered with their store.
// @Tags         Admin Panel - Customer Management
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string true "Tenant ID"
// @Success      200      {object} models.APIResponse[[]models.User]
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve customers"
// @Router       /{tenantId}/admin/customers [get]
func GetTenantCustomersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		customers, err := services.GetTenantCustomers(c.Request.Context(), tenantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve customers"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[[]models.User]{Success: true, Data: customers})
	}
}

// GetDashboardStatsHandler godoc
// @Summary      Get dashboard analytics
// @Description  Retrieves key performance statistics for the tenant's store, such as total revenue and orders today.
// @Tags         Admin Panel - Dashboard
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string true "Tenant ID"
// @Success      200      {object} models.APIResponse[models.DashboardStats]
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve dashboard stats"
// @Router       /{tenantId}/admin/dashboard/stats [get]
func GetDashboardStatsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		stats, err := services.GetDashboardStats(c.Request.Context(), tenantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve dashboard stats"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[*models.DashboardStats]{Success: true, Data: stats})
	}
}
