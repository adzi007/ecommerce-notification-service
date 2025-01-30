package ws

import (
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase"
	"github.com/gofiber/websocket/v2"
)

type WebSocketHandler struct {
	uc *usecase.NotificationUsecase
}

type hub struct {
	Clients               map[*websocket.Conn]bool
	ClientRegisterChanel  chan *websocket.Conn
	ClientRemovalChanel   chan *websocket.Conn
	BroadcastNotification chan domain.Notification
}

// func NewWebSocketHandler(uc *usecase.NotificationUsecase) *WebSocketHandler {
// 	return &WebSocketHandler{uc}
// }

func NewNotificationHub() domain.NotifWebsocket {
	return &hub{
		Clients:               make(map[*websocket.Conn]bool),
		ClientRegisterChanel:  make(chan *websocket.Conn),
		BroadcastNotification: make(chan domain.Notification),
		ClientRemovalChanel:   make(chan *websocket.Conn),
		// ChatService:          cs,
	}
}

func (h *hub) Run() {

	for {
		select {
		case conn := <-h.ClientRegisterChanel:
			h.Clients[conn] = true

		case conn := <-h.ClientRemovalChanel:
			delete(h.Clients, conn)

		case notif := <-h.BroadcastNotification:

			for conn := range h.Clients {

				_ = conn.WriteJSON(notif)
			}

		}
	}

}

// func (h *WebSocketHandler) HandleConnection() func(*websocket.Conn) {

// 	return func(c *websocket.Conn) {
// 		defer c.Close()
// 		userID := c.Query("user_id")
// 		h.uc.Clients[userID] = c
// 	}

// }
