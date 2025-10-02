package controllers

import (
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

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
		_, err = services.CreateProduct(c.Request.Context(), tenantId, &payload)
		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Could not create Product",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[any]{
			Success: true,
			Message: "Product created",
		}
		c.JSON(http.StatusOK, response)
	}
}

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
