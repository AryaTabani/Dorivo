package controllers

import (
	"net/http"
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
