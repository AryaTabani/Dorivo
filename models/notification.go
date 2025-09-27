package models

import (
	"encoding/json"
	"time"
)

type Notification struct {
	ID        int64           `json:"id"`
	UserID    int64           `json:"-"`
	Title     string          `json:"title"`
	Type      string          `json:"type"`
	Content   string          `json:"content"`
	IsRead    bool            `json:"is_read"`
	Metadata  json.RawMessage `json:"metadata,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
}

type MarkReadPayload struct {
	NotificationIDs []int64 `json:"notification_ids" binding:"required"`
}
