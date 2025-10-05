package controllers

import (
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// SuperAdminLoginHandler godoc
// @Summary      Super Admin Login
// @Description  Authenticates a super admin and returns a special JWT for platform management.
// @Tags         Super Admin - Authentication
// @Accept       json
// @Produce      json
// @Param        credentials body     models.SuperAdminLoginPayload true "Super Admin Credentials"
// @Success      200         {object} models.APIResponse[models.LoginResponse] "Login successful"
// @Failure      400         {object} models.APIResponse[any] "Invalid request body"
// @Failure      401         {object} models.APIResponse[any] "Invalid credentials"
// @Router       /superadmin/login [post]
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
		c.JSON(http.StatusOK, models.APIResponse[models.LoginResponse]{Success: true, Data: models.LoginResponse{Token: token}})
	}
}

// CreateTenantHandler godoc
// @Summary      Create a new tenant
// @Description  Allows a super admin to create a new tenant on the platform.
// @Tags         Super Admin - Tenant Management
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        tenant body     models.CreateTenantPayload true "New Tenant Information"
// @Success      201    {object} models.APIResponse[any] "Tenant created successfully"
// @Failure      400    {object} models.APIResponse[any] "Invalid request body"
// @Failure      403    {object} models.APIResponse[any] "Forbidden"
// @Failure      500    {object} models.APIResponse[any] "Failed to create tenant"
// @Router       /superadmin/tenants [post]
func CreateTenantHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.CreateTenantPayload
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

// GetAllTenantsHandler godoc
// @Summary      Get all tenants
// @Description  Retrieves a list of all tenants on the platform.
// @Tags         Super Admin - Tenant Management
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[[]models.Tenant]
// @Failure      403 {object} models.APIResponse[any] "Forbidden"
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve tenants"
// @Router       /superadmin/tenants [get]
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

// DeleteTenantHandler godoc
// @Summary      Delete a tenant
// @Description  Permanently deletes a tenant and all of their associated data from the platform.
// @Tags         Super Admin - Tenant Management
// @Produce      json
// @Security     BearerAuth
// @Param        tenantId path     string true "Tenant ID to delete"
// @Success      200      {object} models.APIResponse[any] "Tenant deleted successfully"
// @Failure      403      {object} models.APIResponse[any] "Forbidden"
// @Failure      500      {object} models.APIResponse[any] "Failed to delete tenant"
// @Router       /superadmin/tenants/{tenantId} [delete]
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
