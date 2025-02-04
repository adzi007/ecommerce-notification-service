package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	hub := ws.NewNotificationHub(uc)
	go hub.Run()

	servernya := server.NewFiberServer(db, hub)

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	go servernya.Start()

	// Graceful Shutdown Handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down the server...")
	servernya.Close()
	log.Println("Server stopped.")
}
