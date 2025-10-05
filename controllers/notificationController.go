package controllers

import (
	"net/http"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/services"
	"github.com/gin-gonic/gin"
)

// GetNotificationsHandler godoc
// @Summary      Get user's notifications
// @Description  Retrieves a list of all notifications for the currently authenticated user.
// @Tags         Notifications & Settings
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} models.APIResponse[[]models.Notification]
// @Failure      500 {object} models.APIResponse[any] "Failed to retrieve notifications"
// @Router       /notifications [get]
func GetNotificationsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		notifications, err := services.GetMyNotifications(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to retrieve notifications"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[[]models.Notification]{Success: true, Data: notifications})
	}
}

// MarkReadHandler godoc
// @Summary      Mark notifications as read
// @Description  Marks one or more notifications as read for the authenticated user.
// @Tags         Notifications & Settings
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        ids body     models.MarkReadPayload true "Notification IDs to mark as read"
// @Success      200 {object} models.APIResponse[any] "Notifications marked as read"
// @Failure      400 {object} models.APIResponse[any] "Invalid request body"
// @Failure      500 {object} models.APIResponse[any] "Failed to mark notifications as read"
// @Router       /notifications/read [put]
func MarkReadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("userID")

		var payload models.MarkReadPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, models.APIResponse[any]{Success: false, Error: "Invalid request body: " + err.Error()})
			return
		}

		err := services.MarkAsRead(c.Request.Context(), userID, payload.NotificationIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.APIResponse[any]{Success: false, Error: "Failed to mark notifications as read"})
			return
		}

		c.JSON(http.StatusOK, models.APIResponse[any]{Success: true, Message: "Notifications marked as read"})
	}
}
