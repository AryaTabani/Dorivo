package controllers

import (
	"net/http"

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
