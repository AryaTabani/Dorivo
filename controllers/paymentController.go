package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// AddPaymentMethodHandler godoc
// @Summary      Add a new payment method
// @Description  Adds a new payment method to the user's profile using a token from a payment processor.
// @Tags         User & Profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        token body     models.AddPaymentMethodPayload true "Payment Processor Token"
// @Success      201   {object} models.APIResponse[any] "Payment method added successfully"
// @Failure      400   {object} models.APIResponse[any] "Invalid request body"
// @Failure      500   {object} models.APIResponse[any] "Failed to add payment method"
// @Router       /payment-methods [post]
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

// GetPaymentMethodsHandler godoc
// @Summary      Get user's payment methods
// @Description  Retrieves a list of all saved payment methods for the authenticated user.
// @Tags         User & Profile
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[[]models.PaymentMethod]
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve payment methods"
// @Router       /payment-methods [get]
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

// DeletePaymentMethodHandler godoc
// @Summary      Delete a payment method
// @Description  Deletes a specific payment method belonging to the authenticated user.
// @Tags         User & Profile
// @Produce      json
// @Security     BearerAuth
// @Param        methodId path     int true "Payment Method ID"
// @Success      200      {object} models.APIResponse[any] "Payment method deleted successfully"
// @Failure      400      {object} models.APIResponse[any] "Invalid payment method ID"
// @Failure      404      {object} models.APIResponse[any] "Payment method not found"
// @Failure      500      {object} models.APIResponse[any] "Failed to delete payment method"
// @Router       /payment-methods/{methodId} [delete]
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
