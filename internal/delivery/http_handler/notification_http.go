package httphandler

import (
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/dto"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/logger"
	"github.com/gofiber/fiber/v2"
)

type notifHttpHandler struct {
	notifUc domain.NotificationUsecase
	notifWs domain.NotifWebsocket
}

func NewCartHttpHandle(notifUc domain.NotificationUsecase, notifWs domain.NotifWebsocket) NotificationHttpHandler {
	return &notifHttpHandler{
		notifUc: notifUc,
		notifWs: notifWs,
	}
}

func (h *notifHttpHandler) InsertNewNotifivation(ctx *fiber.Ctx) error {

	reqBody := new(dto.NotificationData)
	if err := ctx.BodyParser(reqBody); err != nil {
		logger.Error().Err(err).Msg("Error binding request body")
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	data, err := h.notifUc.Insert(reqBody)

	h.notifWs.Broadcast(data)

	if err != nil {

		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to save notification",
			"error":   err.Error(),
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "success create a new cart 1",
	})
}
