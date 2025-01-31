package ws

import (
	"fmt"

	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/k0kubun/pp/v3"
)

type hub struct {
	Clients               map[*websocket.Conn]bool
	ClientRegisterChanel  chan *websocket.Conn
	ClientRemovalChanel   chan *websocket.Conn
	BroadcastNotification chan domain.Notification
	NotificationUc        domain.NotificationUsecase
}

func NewNotificationHub(notifUc domain.NotificationUsecase) domain.NotifWebsocket {
	return &hub{
		Clients:               make(map[*websocket.Conn]bool),
		ClientRegisterChanel:  make(chan *websocket.Conn),
		BroadcastNotification: make(chan domain.Notification),
		ClientRemovalChanel:   make(chan *websocket.Conn),
		NotificationUc:        notifUc,
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

func AllowUpgrade(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {

		token := ctx.Get("token")

		if token != "" {

			// _, err := utils.DecodeToken(token)

			// if err != nil {
			// 	return fiber.ErrUnauthorized
			// }

			return ctx.Next()

		} else {

			return fiber.ErrUnauthorized
		}

	}
	return fiber.ErrUpgradeRequired
}

func (h *hub) HandleNotificationRoom() func(*websocket.Conn) {

	return func(conn *websocket.Conn) {

		userId := conn.Params("userId")

		pp.Println("userId param >>> ", userId)

		defer func() {
			h.ClientRemovalChanel <- conn
			_ = conn.Close()
		}()

		for {

			var notificationMessage domain.NotifMessageRequest
			errReadJson := conn.ReadJSON(&notificationMessage)

			if errReadJson != nil {
				// Handle error
				return
			}

			insertNotification := dto.NotificationData{
				UserID: notificationMessage.UserID,
				Title:  notificationMessage.Title,
				Body:   notificationMessage.Body,
				Link:   notificationMessage.Link,
				Status: notificationMessage.Status,
				IsRead: notificationMessage.IsRead,
			}

			err := h.NotificationUc.Insert(insertNotification)

			if err != nil {
				fmt.Println("errInserNewChat", err.Error())
				return
			}

			h.BroadcastNotification <- domain.Notification{
				ID:     1,
				UserID: "f6gd7fgdff876g8fd",
				Title:  "Test Brodcast",
				Body:   "Lorem ipsum dolor sit amet",
				Link:   "http://www.example.com/lorem/7dfd8fg6df88gf7",
				Status: 0,
				IsRead: 0,
			}

		}

	}
}
