package controllers

import (
	"net/http"
	"strconv"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

func CreateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantId := c.Param("tenantId")

		var payload models.ProductPayload
		err := c.ShouldBindJSON(&payload)

		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Invalid request body",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		_, err = services.CreateProduct(c.Request.Context(), tenantId, &payload)
		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Could not create Product",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[any]{
			Success: true,
			Message: "Product created",
		}
		c.JSON(http.StatusOK, response)
	}
}

func UpdateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)

		var payload models.ProductPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}

		err := services.UpdateProduct(c.Request.Context(), tenantID, productID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Product updated successfully"})
	}
}

func DeleteProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		productID, _ := strconv.ParseInt(c.Param("productId"), 10, 64)

		err := services.DeleteProduct(c.Request.Context(), tenantID, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Product deleted successfully"})
	}
}
