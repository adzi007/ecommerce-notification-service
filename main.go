package main

import (
	"log"

	"github.com/adzi007/ecommerce-notification-service/config"
	"github.com/adzi007/ecommerce-notification-service/internal/delivery"
	"github.com/adzi007/ecommerce-notification-service/internal/delivery/ws"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-notification-service/internal/repository"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	config.LoadConfig()

	mylog := logger.NewLogger()

	db := config.InitDB()

	repo := repository.NewNotificationRepository(db)
	uc := usecase.NewNotificationUsecase(repo)

	go delivery.ConsumeOrderUpdates(uc)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	hub := ws.NewNotificationHub()
	go hub.Run()

	app.Use("/ws", ws.AllowUpgrade)

	app.Use("/ws/notification/:userId", websocket.New(hub.HandleWsChatRoom()))

	// wsCon := ws.NewWebSocketHandler(uc)

	// app.Use("/ws", websocket.New(wsCon.HandleConnection()))

	log.Fatal(app.Listen(":8080"))

}
