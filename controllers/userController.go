package controllers

import (
	"errors"
	"fmt"

	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// RegisterHandler godoc
// @Summary      Register a new user
// @Description  Creates a new user account for a specific tenant.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        tenantId path     string                 true "Tenant ID"
// @Param        user     body     models.RegisterPayload true "User Registration Info"
// @Success      201      {object} models.APIResponse[any] "User created successfully"
// @Failure      400      {object} models.APIResponse[any] "Invalid request body"
// @Failure      409      {object} models.APIResponse[any] "User with this email already exists"
// @Failure      500      {object} models.APIResponse[any] "Failed to create user"
// @Router       /{tenantId}/register [post]
func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")
		fmt.Println("Tenant ID from para:", tenantID)
		var payload models.RegisterPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Error:   "Invalid request body: " + err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		_, err := services.RegisterUser(c.Request.Context(), tenantID, &payload)
		if err != nil {
			if errors.Is(err, services.ErrUserExists) {
				response := models.APIResponse[any]{
					Success: false,
					Error:   err.Error(),
				}
				c.JSON(http.StatusConflict, response)
				return
			}
			response := models.APIResponse[any]{
				Success: false,
				Error:   "Failed to create user",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[any]{
			Success: true,
			Message: "User created successfully",
		}
		c.JSON(http.StatusCreated, response)
	}
}

// LoginHandler godoc
// @Summary      Log in a user
// @Description  Authenticates a user for a specific tenant and returns a JWT token.
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        tenantId    path     string              true "Tenant ID"
// @Param        credentials body     models.LoginPayload true "User Login Credentials"
// @Success      200         {object} models.APIResponse[models.LoginResponse] "Login successful"
// @Failure      400         {object} models.APIResponse[any] "Invalid request body"
// @Failure      401         {object} models.APIResponse[any] "Invalid credentials"
// @Failure      500         {object} models.APIResponse[any] "Login failed"
// @Router       /{tenantId}/login [post]
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.Param("tenantId")

		var payload models.LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			response := models.APIResponse[any]{
				Success: false,
				Error:   "Invalid request body: " + err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token, err := services.LoginUser(c.Request.Context(), tenantID, &payload)
		if err != nil {
			if errors.Is(err, services.ErrInvalidCredentials) {
				response := models.APIResponse[any]{
					Success: false,
					Error:   err.Error(),
				}
				c.JSON(http.StatusUnauthorized, response)
				return
			}
			response := models.APIResponse[any]{
				Success: false,
				Error:   "Login failed",
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := models.APIResponse[models.LoginResponse]{
			Success: true,
			Data:    models.LoginResponse{Token: token},
		}
		c.JSON(http.StatusOK, response)
	}
}

// GetProfileHandler godoc
// @Summary      Get user profile
// @Description  Retrieves the profile information for the currently authenticated user.
// @Tags         User & Profile
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[models.User]
// @Failure      404 {object} models.APIResponse[any] "User not found"
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve profile"
// @Router       /profile [get]
func GetProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		user, err := services.GetProfile(c.Request.Context(), userID)
		if err != nil {
			if errors.Is(err, services.ErrUserNotFound) {
				c.JSON(http.StatusNotFound, models.APIResponse[any]{Success: false, Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve profile"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[*models.User]{Success: true, Data: user})
	}
}

// UpdateProfileHandler godoc
// @Summary      Update user profile
// @Description  Updates the profile information for the currently authenticated user.
// @Tags         User & Profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        profile body     models.UpdateProfilePayload true "Updated profile information"
// @Success      200     {object} models.APIResponse[any] "Profile updated successfully"
// @Failure      400     {object} models.APIResponse[any] "Invalid request body"
// @Failure      500     {object} models.APIResponse[any] "Failed to update profile"
// @Router       /profile [put]
func UpdateProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		var payload models.UpdateProfilePayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err := services.UpdateProfile(c.Request.Context(), userID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to update profile"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Profile updated successfully"})
	}
}

// GetNotificationsSettingHandler godoc
// @Summary      Get notification settings
// @Description  Retrieves the notification preferences for the currently authenticated user.
// @Tags         Notifications & Settings
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[models.NotificationPreferences]
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve settings"
// @Router       /profile/notification-settings [get]
func GetNotificationsSettingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		user, err := services.GetProfile(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve profile"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[models.NotificationPreferences]{Success: true, Data: user.NotificationPreferences})
	}
}

// UpdateNotificationSettingsHandler godoc
// @Summary      Update notification settings
// @Description  Updates the notification preferences for the currently authenticated user.
// @Tags         Notifications & Settings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        settings body     models.NotificationPreferences true "New notification settings"
// @Success      200      {object} models.APIResponse[any] "Settings updated successfully"
// @Failure      400      {object} models.APIResponse[any] "Invalid request body"
// @Failure      500      {object} models.APIResponse[any] "Failed to update settings"
// @Router       /profile/notification-settings [put]
func UpdateNotificationSettingsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		var payload models.NotificationPreferences
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}
		err := services.UpdateNotificationPreferences(c.Request.Context(), userID, &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to update notification settings"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Notification settings updated successfully"})
	}
}

// ChangePasswordHandler godoc
// @Summary      Change user password
// @Description  Allows an authenticated user to change their password by providing their current password.
// @Tags         User & Profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        passwords body     models.ChangePasswordPayload true "Current and new password"
// @Success      200       {object} models.APIResponse[any] "Password changed successfully"
// @Failure      400       {object} models.APIResponse[any] "Invalid request body"
// @Failure      401       {object} models.APIResponse[any] "The current password is incorrect"
// @Failure      500       {object} models.APIResponse[any] "Failed to change password"
// @Router       /profile/change-password [put]
func ChangePasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		var payload models.ChangePasswordPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}
		err := services.ChangePassword(c.Request.Context(), userID, &payload)
		if err != nil {
			if errors.Is(err, services.ErrInvalidCredentials) {
				c.JSON(http.StatusUnauthorized, models.APIResponse[any]{Success: false, Error: "The current password is incorrect"})
				return
			}
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to change password"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Password changed successfully"})
	}
}

// DeleteAccountHandler godoc
// @Summary      Delete user account
// @Description  Permanently deletes the account of the currently authenticated user. This action is irreversible.
// @Tags         User & Profile
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[any] "Account deleted successfully"
// @Failure      500 {object} models.APIResponse[any] "Failed to delete account"
// @Router       /profile [delete]
func DeleteAccountHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")
		err := services.DeleteAccount(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to delete account"})
			return
		}
		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Account deleted successfully"})
	}
}
