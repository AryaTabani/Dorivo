package controllers

import (
	"errors"
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

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
