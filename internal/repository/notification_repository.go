package repository

import (
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/k0kubun/pp/v3"
	"gorm.io/gorm"
)

// Notification Model
// type Notification struct {
// 	ID        uint   `gorm:"primaryKey"`
// 	OrderID   string `gorm:"not null"`
// 	UserID    string `gorm:"not null"`
// 	Email     string `gorm:"not null"`
// 	Message   string `gorm:"not null"`
// 	Status    string `gorm:"default:pending"`
// 	CreatedAt string
// }

// Repository
type NotificationRepositoryStruct struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) domain.NotificationRepository {
	return &NotificationRepositoryStruct{
		db: db,
	}
}

func (repo *NotificationRepositoryStruct) FindByUser(userId string) ([]domain.Notification, error) {

	var notifications []domain.Notification
	err := repo.db.Where("user_id = ?", userId).Find(&notifications).Error
	return notifications, err

}

func (repo *NotificationRepositoryStruct) Insert(notification dto.NotificationData) error {

	insertNotif := &domain.Notification{
		UserID: notification.UserID,
		Title:  notification.Title,
		Body:   notification.Body,
		Link:   notification.Link,
		Status: notification.Status,
		IsRead: notification.IsRead,
	}

	result := repo.db.Create(&insertNotif)

	pp.Println("result >>> ", insertNotif)

	return result.Error
}

func (repo *NotificationRepositoryStruct) Update(data domain.Notification) error {
	return repo.db.Model(&domain.Notification{}).Where("id = ?", data.ID).Update("IsRead", data.IsRead).Error
}
