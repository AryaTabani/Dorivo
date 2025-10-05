package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// AddToCartHandler godoc
// @Summary      Add an item to the cart
// @Description  Adds a product with selected options and quantity to the user's shopping cart.
// @Tags         Cart & Checkout
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        item body     models.AddToCartPayload true "Item to Add"
// @Success      200  {object} models.APIResponse[any] "Item added to cart"
// @Failure      400  {object} models.APIResponse[any] "Invalid request body"
// @Failure      500  {object} models.APIResponse[any] "Failed to add item to cart"
// @Router       /cart/items [post]
func AddToCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		var payload models.AddToCartPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err := services.AddToCart(c.Request.Context(), userID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to add item to cart"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Item added to cart"})
	}
}

// GetCartHandler godoc
// @Summary      Get cart contents
// @Description  Retrieves the full contents of the user's shopping cart, with calculated totals.
// @Tags         Cart & Checkout
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[models.Cart]
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve cart"
// @Router       /cart [get]
func GetCartHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		cart, err := services.GetCart(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve cart"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[*models.Cart]{Success: true, Data: cart})
	}
}

// UpdateCartItemHandler godoc
// @Summary      Update item quantity
// @Description  Updates the quantity of a specific item in the user's cart.
// @Tags         Cart & Checkout
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        itemId   path     int                         true "Cart Item ID"
// @Param        quantity body     models.UpdateCartItemPayload true "New Quantity"
// @Success      200      {object} models.APIResponse[any] "Item quantity updated"
// @Failure      400      {object} models.APIResponse[any] "Invalid request body or item ID"
// @Failure      404      {object} models.APIResponse[any] "Cart item not found"
// @Failure      500      {object} models.APIResponse[any] "Failed to update item"
// @Router       /cart/items/{itemId} [put]
func UpdateCartItemHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		itemID, err := strconv.ParseInt(c.Param("itemId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid item ID"})
			return
		}

		var payload models.UpdateCartItemPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err = services.UpdateCartItemQuantity(c.Request.Context(), userID, itemID, payload.Quantity)
		if err != nil {
			if errors.Is(err, services.ErrCartItemNotFound) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to update item"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Item quantity updated"})
	}
}

// RemoveCartItemHandler godoc
// @Summary      Remove item from cart
// @Description  Removes a specific item entirely from the user's shopping cart.
// @Tags         Cart & Checkout
// @Produce      json
// @Security     BearerAuth
// @Param        itemId path     int true "Cart Item ID"
// @Success      200    {object} models.APIResponse[any] "Item removed from cart"
// @Failure      400    {object} models.APIResponse[any] "Invalid item ID"
// @Failure      404    {object} models.APIResponse[any] "Cart item not found"
// @Failure      500    {object} models.APIResponse[any] "Failed to remove item"
// @Router       /cart/items/{itemId} [delete]
func RemoveCartItemHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		itemID, err := strconv.ParseInt(c.Param("itemId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid item ID"})
			return
		}

		err = services.RemoveCartItem(c.Request.Context(), userID, itemID)
		if err != nil {
			if errors.Is(err, services.ErrCartItemNotFound) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to remove item"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Item removed from cart"})
	}
}
