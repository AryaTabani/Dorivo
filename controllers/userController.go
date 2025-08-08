package controllers

import (
	"errors"
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID := c.GetHeader("tenantid")

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
		tenantID := c.GetHeader("tenantId")

		var payload models.LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			response := models.APIResponse[any]{
					Success: false,
					Error:  "Invalid request body: " + err.Error(),
				}
			c.JSON(http.StatusBadRequest,response)
			return
		}

		token, err := services.LoginUser(c.Request.Context(), tenantID, &payload)
		if err != nil {
			if errors.Is(err, services.ErrInvalidCredentials) {
				response := models.APIResponse[any]{
					Success: false,
					Error:  err.Error(),
				}
				c.JSON(http.StatusUnauthorized,response)
				return
			}
			response := models.APIResponse[any]{
					Success: false,
					Error:  "Login failed",
				}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
        response := models.APIResponse[any]{
			Success: true,
			Message: "token: "+token,
		}
		c.JSON(http.StatusOK,response)
	}
}
