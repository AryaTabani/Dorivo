package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

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
