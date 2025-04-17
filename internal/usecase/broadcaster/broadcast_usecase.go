package broadcaster

import "github.com/adzi007/ecommerce-notification-service/internal/domain"

type broadcaster struct {
	notifWs domain.NotifWebsocket
}

func NewBroadcaster(notifWs domain.NotifWebsocket) domain.BroadcasterUsecase {
	return &broadcaster{
		notifWs: notifWs,
	}
}

func (h *broadcaster) Broadcast(data domain.Notification) {

	h.notifWs.Broadcast(data)

}
