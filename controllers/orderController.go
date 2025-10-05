package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// GetMyOrdersHandler godoc
// @Summary      Get user's order history
// @Description  Retrieves a list of orders for the authenticated user, filterable by status.
// @Tags         Orders
// @Produce      json
// @Security     BearerAuth
// @Param        status query    string false "Filter orders by status (e.g., Active, Completed, Cancelled)"
// @Success      200    {object} models.APIResponse[[]models.OrderSummaryView]
// @Failure      500    {object} models.APIResponse[any] "Failed to retrieve orders"
// @Router       /orders [get]
func GetMyOrdersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		status := c.Query("status")

		orders, err := services.GetMyOrders(c.Request.Context(), userID, status)
		if err != nil {
			return
		}

		response := models.APIResponse[[]models.OrderSummaryView]{
			Success: true,
			Data:    orders,
		}
		c.JSON(http.StatusOK, response)
	}
}

// CancelOrderHandler godoc
// @Summary      Cancel an active order
// @Description  Allows an authenticated user to cancel one of their own active orders.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        orderId path     int                      true "Order ID"
// @Param        reason  body     models.CancelOrderPayload true "Reason for cancellation"
// @Success      200     {object} models.APIResponse[any] "Order cancelled successfully"
// @Failure      400     {object} models.APIResponse[any] "Invalid order ID or request body"
// @Failure      404     {object} models.APIResponse[any] "Order not found"
// @Failure      409     {object} models.APIResponse[any] "Order cannot be cancelled"
// @Failure      500     {object} models.APIResponse[any] "Failed to cancel order"
// @Router       /orders/{orderId}/cancel [post]
func CancelOrderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		orderID, err := strconv.ParseInt(c.Param("orderId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid order ID"})
			return
		}

		var payload models.CancelOrderPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err = services.CancelOrder(c.Request.Context(), userID, orderID, payload.Reason)
		if err != nil {
			if errors.Is(err, services.ErrOrderNotFound) || errors.Is(err, services.ErrForbidden) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			if errors.Is(err, services.ErrOrderCannotBeCancelled) {
				c.JSON(http.StatusConflict, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to cancel order"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Order cancelled successfully"})
	}
}

// LeaveReviewHandler godoc
// @Summary      Leave a review for an order
// @Description  Allows an authenticated user to leave a rating and comment for one of their own completed orders.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        orderId path     int                      true "Order ID"
// @Param        review  body     models.LeaveReviewPayload true "Rating and comment for the order"
// @Success      201     {object} models.APIResponse[any] "Review submitted successfully"
// @Failure      400     {object} models.APIResponse[any] "Invalid order ID or request body"
// @Failure      404     {object} models.APIResponse[any] "Order not found"
// @Failure      409     {object} models.APIResponse[any] "Order is not completed or a review already exists"
// @Failure      500     {object} models.APIResponse[any] "Failed to leave review"
// @Router       /orders/{orderId}/review [post]
func LeaveReviewHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		orderID, err := strconv.ParseInt(c.Param("orderId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid order ID"})
			return
		}

		var payload models.LeaveReviewPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err = services.LeaveReview(c.Request.Context(), userID, orderID, &payload)
		if err != nil {
			if errors.Is(err, services.ErrOrderNotFound) || errors.Is(err, services.ErrForbidden) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			if errors.Is(err, services.ErrOrderNotCompleted) || errors.Is(err, services.ErrReviewExists) {
				c.JSON(http.StatusConflict, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to leave review"})
			return
		}

		c.JSON(http.StatusCreated, models.APIResponse[any]{Success: true, Message: "Review submitted successfully"})
	}
}
