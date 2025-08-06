package controllers

import (
	"errors"
	"net/http"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetTenantConfigHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("tenantId")
		config, err := services.GetTenantConfig(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, services.ErrTenantNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve tenant configuration"})
			return
		}

		c.JSON(http.StatusOK, config)
	}
}
