package controllers

import (
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// GetFavoritesHandler godoc
// @Summary      Get user's favorite products
// @Description  Retrieves a list of all products that the authenticated user has marked as a favorite.
// @Tags         Favorites
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[[]models.Product]
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve favorites"
// @Router       /favorites [get]
func GetFavoritesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		favorites, err := services.GetFavorites(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve favorites"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[[]models.Product]{Success: true, Data: favorites})
	}
}

// AddToFavoritesHandler godoc
// @Summary      Add a product to favorites
// @Description  Adds a specific product to the authenticated user's list of favorites.
// @Tags         Favorites
// @Produce      json
// @Security     BearerAuth
// @Param        productId path     int true "Product ID"
// @Success      201       {object} models.APIResponse[any] "Product added to favorites"
// @Failure      400       {object} models.APIResponse[any] "Invalid product ID"
// @Failure      500       {object} models.APIResponse[any] "Failed to add to favorites"
// @Router       /products/{productId}/favorite [post]
func AddToFavoritesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		productID, err := strconv.ParseInt(c.Param("productId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid product ID"})
			return
		}

		err = services.AddToFavorites(c.Request.Context(), userID, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to add to favorites"})
			return
		}

		c.JSON(http.StatusCreated, models.APIResponse[any]{Success: true, Message: "Product added to favorites"})
	}
}

// RemoveFromFavoritesHandler godoc
// @Summary      Remove a product from favorites
// @Description  Removes a specific product from the authenticated user's list of favorites.
// @Tags         Favorites
// @Produce      json
// @Security     BearerAuth
// @Param        productId path     int true "Product ID"
// @Success      200       {object} models.APIResponse[any] "Product removed from favorites"
// @Failure      400       {object} models.APIResponse[any] "Invalid product ID"
// @Failure      500       {object} models.APIResponse[any] "Failed to remove from favorites"
// @Router       /products/{productId}/favorite [delete]
func RemoveFromFavoritesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		productID, err := strconv.ParseInt(c.Param("productId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid product ID"})
			return
		}

		err = services.RemoveFromFavorites(c.Request.Context(), userID, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to remove from favorites"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Product removed from favorites"})
	}
}
