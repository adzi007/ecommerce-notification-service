package dto

import "time"

type NotificationData struct {
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Link      string    `json:"link"`
	Status    int8      `json:"status"`
	IsRead    int8      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
