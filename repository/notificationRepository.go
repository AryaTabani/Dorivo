package repository

import (
	"context"
	"strings"

	db "github.com/AryaTabani/Dorivo/DB"
	"github.com/AryaTabani/Dorivo/models"
)

func CreateNotification(ctx context.Context, n *models.Notification) error {
	query := `INSERT INTO notifications (user_id, title, type, metadata) VALUES (?, ?, ?, ?)`
	_, err := db.DB.ExecContext(ctx, query, n.UserID, n.Title, n.Type, n.Metadata)
	return err
}

func GetNotificationsByUserID(ctx context.Context, userID int64) ([]*models.Notification, error) {
	query := `SELECT id, user_id, title, type, content, is_read, metadata, created_at FROM notifications WHERE user_id = ? ORDER BY created_at DESC`
	rows, err := db.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notifications []*models.Notification
	for rows.Next() {
		var n models.Notification
		if err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Type, &n.Content, &n.IsRead, &n.Metadata, &n.CreatedAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, &n)
	}
	return notifications, nil
}

func MarkNotificationsAsRead(ctx context.Context, userID int64, notificationIDs []int64) error {
	if len(notificationIDs) == 0 {
		return nil
	}
	query := `UPDATE notifications SET is_read = TRUE WHERE user_id = ? AND id IN (?` + strings.Repeat(",?", len(notificationIDs)-1) + `)`
	args := make([]interface{}, 0, len(notificationIDs)+1)
	args[0] = userID
	for i, id := range notificationIDs {
		args[i+1] = id
	}
	_, err := db.DB.ExecContext(ctx, query, args...)
	return err
}
