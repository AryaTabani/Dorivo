package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

func AddAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		var payload models.AddAddressPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err := services.AddAddress(c.Request.Context(), userID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to add new address"})
			return
		}

		c.JSON(http.StatusCreated, models.APIResponse[any]{Success: true, Message: "Address added successfully"})
	}
}

func GetAddressesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		addresses, err := services.GetMyAddresses(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve addresses"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[[]models.Address]{Success: true, Data: addresses})
	}
}

func DeleteAddressHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		addressID, err := strconv.ParseInt(c.Param("addressId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid address ID"})
			return
		}

		err = services.DeleteAddress(c.Request.Context(), userID, addressID)
		if err != nil {
			if errors.Is(err, services.ErrAddressNotFound) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to delete address"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Address deleted successfully"})
	}
}
