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

// SearchProductsHandler godoc
// @Summary      Search and filter products
// @Description  Retrieves a list of products for a tenant, with optional filters for category, tags, price, and sorting.
// @Tags         Public - Products
// @Produce      json
// @Param        tenantId  path     string false "Tenant ID"
// @Param        category  query    string false "Filter by main category (e.g., Meal, Drink)"
// @Param        tags      query    string false "Filter by comma-separated tags (e.g., Pizza,Cheese)"
// @Param        min_price query    number false "Minimum price filter"
// @Param        max_price query    number false "Maximum price filter"
// @Param        sort_by   query    string false "Sort order (e.g., rating_desc)"
// @Success      200       {object} models.APIResponse[[]models.Product]
// @Failure      500       {object} models.APIResponse[any] "Failed to search for products"
// @Router       /{tenantId}/products [get]
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

// GetTagsHandler godoc
// @Summary      Get all filter tags
// @Description  Retrieves a list of all available tags for a specific tenant, used for building filter UIs.
// @Tags         Public - Products
// @Produce      json
// @Param        tenantId path     string true "Tenant ID"
// @Success      200      {object} models.APIResponse[[]models.Tag]
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve tags"
// @Router       /{tenantId}/tags [get]
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

// GetProductDetailsHandler godoc
// @Summary      Get product details
// @Description  Retrieves detailed information for a single product, including its customization options.
// @Tags         Public - Products
// @Produce      json
// @Param        tenantId  path     string true "Tenant ID"
// @Param        productId path     int    true "Product ID"
// @Success      200       {object} models.APIResponse[models.Product]
// @Failure      400       {object} models.APIResponse[any] "Invalid product ID"
// @Failure      404       {object} models.APIResponse[any] "Product not found"
// @Failure      500       {object} models.APIResponse[any] "Failed to retrieve product details"
// @Router       /{tenantId}/products/{productId} [get]
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

// GetBestSellersHandler godoc
// @Summary      Get best-selling products
// @Description  Retrieves a list of the most popular products for a tenant, based on sales volume.
// @Tags         Public - Products
// @Produce      json
// @Param        tenantId path     string true "Tenant ID"
// @Success      200      {object} models.APIResponse[[]models.Product]
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve best sellers"
// @Router       /{tenantId}/products/bestsellers [get]
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

// GetFeaturedProductHandler godoc
// @Summary      Get the featured product
// @Description  Retrieves the main promotional product for a tenant, often used for a large banner.
// @Tags         Public - Products
// @Produce      json
// @Param        tenantId path     string true "Tenant ID"
// @Success      200      {object} models.APIResponse[models.Product]
// @Failure      404      {object} models.APIResponse[any] "No featured product found"
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve featured product"
// @Router       /{tenantId}/products/featured [get]
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

// GetRecommendedProductsHandler godoc
// @Summary      Get recommended products
// @Description  Retrieves a list of products marked as 'recommended' by the tenant admin.
// @Tags         Public - Products
// @Produce      json
// @Param        tenantId path     string true "Tenant ID"
// @Success      200      {object} models.APIResponse[[]models.Product]
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve recommended products"
// @Router       /{tenantId}/products/recommended [get]
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
