package controllers

import (
	"errors"
	"fmt"

	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

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
		response := models.APIResponse[any]{
			Success: true,
			Message: "token: " + token,
		}
		c.JSON(http.StatusOK, response)
	}
}
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
