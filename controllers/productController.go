package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

func SearchProductsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetString("tenantID")
		filters := make(map[string][]string)

		for key, values := range c.Request.URL.Query() {
			if key == "tags" && len(values) > 0 {
				filters[key] = strings.Split(values[0], ",")
			} else {
				filters[key] = values
			}
		}
		products, err := services.SearchProducts(c.Request.Context(), tenantID, filters)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to search for products"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[[]models.Product]{Success: true, Data: products})
	}
}

func GetTagsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetString("tenantID")
		tags, err := services.GetTags(c.Request.Context(), tenantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve tags"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[[]models.Tag]{Success: true, Data: tags})
	}
}

func GetProductDetailsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		productID, err := strconv.ParseInt(c.Param("productId"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid product ID"})
			return
		}

		product, err := services.GetProductDetails(c.Request.Context(), tenantID, productID)
		if err != nil {
			if errors.Is(err, services.ErrProductNotFound) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve product details"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[*models.Product]{Success: true, Data: product})
	}
}
func GetBestSellersHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		products, err := services.GetBestSellers(c.Request.Context(), tenantID)

		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Failed to retrieve best sellers",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[[]models.Product]{
			Success: true,
			Data:    products,
		}
		c.JSON(http.StatusOK, response)
	}
}
func GetFeaturedProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		product, err := services.GetFeaturedProduct(c.Request.Context(), tenantID)

		if err != nil {
			if errors.Is(err, services.ErrProductNotFound) {
				response := models.APIResponse[any]{
					Success: false,
					Message: "No featured product found",
				}
				c.JSON(http.StatusNotFound, response)
				return
			}
			response := models.APIResponse[any]{
				Success: false,
				Message: "Failed to retrieve featured product",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[*models.Product]{
			Success: true,
			Data:    product,
		}
		c.JSON(http.StatusOK, response)
	}
}
func GetRecommendedProductsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		products, err := services.GetRecommendedProducts(c.Request.Context(), tenantID)
		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Message: "Failed to retrieve recommended products",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[[]models.Product]{
			Success: true,
			Data:    products,
		}
		c.JSON(http.StatusOK, response)
	}
}
