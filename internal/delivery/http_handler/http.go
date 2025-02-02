package httphandler

import "github.com/gofiber/fiber/v2"

type NotificationHttpHandler interface {
	InsertNewNotifivation(ctx *fiber.Ctx) error
	GetNotificationByUser(ctx *fiber.Ctx) error
}
