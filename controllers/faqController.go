package controllers

import (
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// GetFAQsHandler godoc
// @Summary      Get FAQs for a tenant
// @Description  Retrieves a list of frequently asked questions for a specific tenant, optionally filtered by category.
// @Tags         Public
// @Produce      json
// @Param        tenantId path     string true "Tenant ID"
// @Param        category query    string false "Filter FAQs by category"
// @Success      200      {object} models.APIResponse[[]models.FAQ]
// @Failure      500      {object} models.APIResponse[any] "Failed to retrieve FAQs"
// @Router       /{tenantId}/faqs [get]
func GetFAQsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		category := c.Query("category")

		faqs, err := services.GetFAQs(c.Request.Context(), tenantID, category)
		if err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Error:   "Failed to retrieve FAQs",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[[]models.FAQ]{
			Success: true,
			Data:    faqs,
		}
		c.JSON(http.StatusOK, response)
	}
}
