package usecase

import (
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/gofiber/contrib/websocket"
)

type notificationUsecase struct {
	// repo    *repository.NotificationRepositoryStruct
	repo    domain.NotificationRepository
	Clients map[string]*websocket.Conn
}

func NewNotificationUsecase(repo domain.NotificationRepository) domain.NotificationUsecase {
	return &notificationUsecase{
		repo:    repo,
		Clients: make(map[string]*websocket.Conn),
	}
}

// WebSocket Notification
func (uc *notificationUsecase) Insert(data *dto.NotificationData) (domain.Notification, error) {
	return uc.repo.Insert(data)
}

func (uc *notificationUsecase) FindByUser(userId string) ([]domain.Notification, error) {
	return uc.repo.FindByUser(userId)
}
