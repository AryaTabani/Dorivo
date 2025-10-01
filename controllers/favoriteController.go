package controllers

import (
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

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
