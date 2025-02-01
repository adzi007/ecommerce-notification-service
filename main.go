package main

import (
	"github.com/adzi007/ecommerce-notification-service/config"
	"github.com/adzi007/ecommerce-notification-service/internal/delivery/ws"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/logger"
	"github.com/adzi007/ecommerce-notification-service/internal/repository"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase"
	"github.com/adzi007/ecommerce-notification-service/server"
	"github.com/gofiber/contrib/fiberzerolog"
)

func main() {
	config.LoadConfig()

	mylog := logger.NewLogger()

	db := database.NewDatabase()

	repo := repository.NewNotificationRepository(db)
	uc := usecase.NewNotificationUsecase(repo)

	// // go delivery.ConsumeOrderUpdates(uc)

	// app := fiber.New()
	// app.Use(fiberzerolog.New(fiberzerolog.Config{
	// 	Logger: &mylog,
	// }))

	hub := ws.NewNotificationHub(uc)
	go hub.Run()

	servernya := server.NewFiberServer(db, hub)

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	// app.Use("/ws", ws.AllowUpgrade)
	// app.Use("/ws/notification/:userId", websocket.New(hub.HandleNotificationRoom()))
	// log.Fatal(app.Listen(":8080"))

	servernya.Start()

}
