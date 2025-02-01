package repository

import (
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/database"
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
	db database.Database
}

func NewNotificationRepository(db database.Database) domain.NotificationRepository {
	return &NotificationRepositoryStruct{
		db: db,
	}
}

func (repo *NotificationRepositoryStruct) FindByUser(userId string) ([]domain.Notification, error) {

	var notifications []domain.Notification
	err := repo.db.GetDb().Where("user_id = ?", userId).Find(&notifications).Error
	return notifications, err

}

func (repo *NotificationRepositoryStruct) Insert(notification *dto.NotificationData) (domain.Notification, error) {

	insertNotif := domain.Notification{
		UserID: notification.UserID,
		Title:  notification.Title,
		Body:   notification.Body,
		Link:   notification.Link,
		Status: notification.Status,
		IsRead: notification.IsRead,
	}

	result := repo.db.GetDb().Create(&insertNotif)

	// pp.Println("result >>> ", insertNotif)

	return insertNotif, result.Error
}

func (repo *NotificationRepositoryStruct) Update(data domain.Notification) error {
	return repo.db.GetDb().Model(&domain.Notification{}).Where("id = ?", data.ID).Update("IsRead", data.IsRead).Error
}
