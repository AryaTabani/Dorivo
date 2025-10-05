package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// AddAddressHandler godoc
// @Summary      Add a new address
// @Description  Adds a new delivery address to the authenticated user's profile.
// @Tags         User & Profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        address body     models.AddAddressPayload true "Address Information"
// @Success      201     {object} models.APIResponse[any] "Address added successfully"
// @Failure      400     {object} models.APIResponse[any] "Invalid request body"
// @Failure      500     {object} models.APIResponse[any] "Failed to add new address"
// @Router       /addresses [post]
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

// GetAddressesHandler godoc
// @Summary      Get user's addresses
// @Description  Retrieves a list of all saved delivery addresses for the authenticated user.
// @Tags         User & Profile
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[[]models.Address]
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve addresses"
// @Router       /addresses [get]
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

// DeleteAddressHandler godoc
// @Summary      Delete an address
// @Description  Deletes a specific address belonging to the authenticated user.
// @Tags         User & Profile
// @Produce      json
// @Security     BearerAuth
// @Param        addressId path     int true "Address ID"
// @Success      200       {object} models.APIResponse[any] "Address deleted successfully"
// @Failure      400       {object} models.APIResponse[any] "Invalid address ID"
// @Failure      404       {object} models.APIResponse[any] "Address not found"
// @Failure      500       {object} models.APIResponse[any] "Failed to delete address"
// @Router       /addresses/{addressId} [delete]
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
