package controllers

import (
	"errors"
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// GetTenantConfigHandler godoc
// @Summary      Get tenant configuration
// @Description  Retrieves the public configuration for a specific tenant, such as name, logo, theme, and contact info.
// @Tags         Public
// @Produce      json
// @Param        tenantId path     string true "The unique name of the tenant"
// @Success      200      {object} models.APIResponse[models.TenantConfig] "Configuration fetched successfully"
// @Failure      404      {object} models.APIResponse[any] "Tenant not found"
// @Failure      500      {object} models.APIResponse[any] "Could not retrieve tenant configuration"
// @Router       /tenant/{tenantId} [get]
func GetTenantConfigHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("tenantId")
		config, err := services.GetTenantConfig(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, services.ErrTenantNotFound) {
				response := models.APIResponse[any]{
					Success: false,
					Error:   err.Error(),
				}
				c.JSON(http.StatusNotFound, response)
				return
			}

			response := models.APIResponse[any]{
				Success: false,
				Error:   "Could not retrieve tenant configuration.",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		response := models.APIResponse[models.TenantConfig]{
			Success: true,
			Message: "Configuration fetched successfully.",
			Data:    *config,
		}
		c.JSON(http.StatusOK, response)
	}
}
