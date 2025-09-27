package services

import (
	"context"
	"encoding/json"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func CreateOrderStatusNotification(ctx context.Context, userID, orderID int64, title string) error {
	metadata, _ := json.Marshal(map[string]int64{"order_id": orderID})

	notification := &models.Notification{
		UserID:   userID,
		Title:    title,
		Type:     "order_status",
		Metadata: metadata,
	}
	return repository.CreateNotification(ctx, notification)
}

func GetMyNotifications(ctx context.Context, userID int64) ([]models.Notification, error) {
	notificationsPtr, err := repository.GetNotificationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	notifications := make([]models.Notification, len(notificationsPtr))
	for i, n := range notificationsPtr {
		if n != nil {
			notifications[i] = *n
		}
	}
	return notifications, nil
}

func MarkAsRead(ctx context.Context, userID int64, notificationIDs []int64) error {
	if len(notificationIDs) == 0 {
		return nil
	}
	return repository.MarkNotificationsAsRead(ctx, userID, notificationIDs)
}
