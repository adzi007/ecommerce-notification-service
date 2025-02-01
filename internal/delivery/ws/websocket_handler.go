package ws

import (
	"fmt"

	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type notificationChanel struct {
	Client *websocket.Conn
	UserID string
}

type hub struct {
	Clients                 map[*websocket.Conn]bool
	ClientRegisterChanel    chan *websocket.Conn
	ClientRemovalChanel     chan *websocket.Conn
	BroadcastNotification   chan domain.Notification
	NotificationUc          domain.NotificationUsecase
	NotifiChanelConnections []notificationChanel
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

			for i := range h.NotifiChanelConnections {

				if h.NotifiChanelConnections[i].UserID == notif.UserID {
					_ = h.NotifiChanelConnections[i].Client.WriteJSON(notif)
				}

			}
			// for conn := range h.Clients {
			// 	_ = conn.WriteJSON(notif)
			// }
		}
	}
}

func (h *hub) Join(client *websocket.Conn, userId string) {

	h.ClientRegisterChanel <- client

	h.NotifiChanelConnections = append(h.NotifiChanelConnections, notificationChanel{
		Client: client,
		UserID: userId,
	})
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

		// pp.Println("userId param >>> ", userId)

		defer func() {
			h.ClientRemovalChanel <- conn
			_ = conn.Close()
		}()

		// h.ClientRegisterChanel <- conn

		h.Join(conn, userId)

		for {

			var notificationMessage domain.NotifMessageRequest
			errReadJson := conn.ReadJSON(&notificationMessage)

			if errReadJson != nil {
				// Handle error
				return
			}

			insertNotification := &dto.NotificationData{
				UserID: notificationMessage.UserID,
				Title:  notificationMessage.Title,
				Body:   notificationMessage.Body,
				Link:   notificationMessage.Link,
				Status: notificationMessage.Status,
				IsRead: notificationMessage.IsRead,
			}

			data, err := h.NotificationUc.Insert(insertNotification)

			if err != nil {
				fmt.Println("errInserNewChat", err.Error())
				return
			}

			h.BroadcastNotification <- data
		}
	}
}

func (h *hub) Broadcast(data domain.Notification) {

	h.BroadcastNotification <- data
}
