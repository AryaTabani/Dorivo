package controllers

import (
	"errors"
	"net/http"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.RegisterPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
			return
		}

		_, err := services.RegisterUser(c.Request.Context(), &payload)
		if err != nil {
			if errors.Is(err, services.ErrUserExists) {
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) // 409 Conflict
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		token, err := services.LoginUser(c.Request.Context(), &payload)
		if err != nil {
			if errors.Is(err, services.ErrInvalidCredentials) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()}) // 401 Unauthorized
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
