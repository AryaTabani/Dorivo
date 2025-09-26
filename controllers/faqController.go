package controllers

import (
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

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
