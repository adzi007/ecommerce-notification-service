package usecase

import (
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/gofiber/contrib/websocket"
)

type notificationUsecase struct {
	// repo    *repository.NotificationRepositoryStruct
	repo        domain.NotificationRepository
	Clients     map[string]*websocket.Conn
	broadcaster domain.BroadcasterUsecase
}

func NewNotificationUsecase(repo domain.NotificationRepository, broadcaster domain.BroadcasterUsecase) domain.NotificationUsecase {
	return &notificationUsecase{
		repo:        repo,
		Clients:     make(map[string]*websocket.Conn),
		broadcaster: broadcaster,
	}
}

// WebSocket Notification
func (uc *notificationUsecase) Insert(data *dto.NotificationData) (domain.Notification, error) {
	// return uc.repo.Insert(data)

	dataNotif, err := uc.repo.Insert(data)

	// if err != nil {
	// }

	uc.broadcaster.Broadcast(dataNotif)

	return dataNotif, err
}

func (uc *notificationUsecase) FindByUser(userId string) ([]domain.Notification, error) {
	return uc.repo.FindByUser(userId)
}
