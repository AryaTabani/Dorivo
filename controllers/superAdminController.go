package controllers

import (
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

func SuperAdminLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.SuperAdminLoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}
		token, err := services.LoginSuperAdmin(c.Request.Context(), &payload)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Data: gin.H{"token": token}})
	}
}

func CreateTenantHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload struct {
			Name   string              `json:"name" binding:"required"`
			Config models.TenantConfig `json:"config" binding:"required"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: err.Error()})
			return
		}
		err := services.CreateTenant(c.Request.Context(), payload.Name, &payload.Config)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to create tenant"})
			return
		}
		c.JSON(http.StatusCreated, models.APIResponse[any]{Success: true, Message: "Tenant created successfully"})
	}
}

func GetAllTenantsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenants, err := services.GetAllTenants(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve tenants"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[[]models.Tenant]{Success: true, Data: tenants})
	}
}

func DeleteTenantHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		err := services.DeleteTenant(c.Request.Context(), tenantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to delete tenant"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Tenant deleted successfully"})
	}
}
