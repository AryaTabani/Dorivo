package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

func AddPaymentMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		var payload models.AddPaymentMethodPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err := services.AddPaymentMethod(c.Request.Context(), userID, &payload)
		if err != nil {
			if errors.Is(err, services.ErrProcessorFailed) {
				c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to add payment method"})
			return
		}

		c.JSON(http.StatusCreated, models.APIResponse[any]{Success: true, Message: "Payment method added successfully"})
	}
}

func GetPaymentMethodsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		methods, err := services.GetMyPaymentMethods(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve payment methods"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[[]models.PaymentMethod]{Success: true, Data: methods})
	}
}

func DeletePaymentMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		methodID, err := strconv.ParseInt(c.Param("methodId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid payment method ID"})
			return
		}

		err = services.DeletePaymentMethod(c.Request.Context(), userID, methodID)
		if err != nil {
			if errors.Is(err, services.ErrPaymentMethodNotFound) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to delete payment method"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Payment method deleted successfully"})
	}
}
