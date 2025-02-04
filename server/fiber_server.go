package server

import (
	"log"

	httphandler "github.com/adzi007/ecommerce-notification-service/internal/delivery/http_handler"
	"github.com/adzi007/ecommerce-notification-service/internal/delivery/ws"
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/rabbitmq"
	"github.com/adzi007/ecommerce-notification-service/internal/repository"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app      *fiber.App
	db       database.Database
	notifWs  domain.NotifWebsocket
	rabbitMQ *rabbitmq.RabbitMQ
}

func NewFiberServer(db database.Database, notifWs domain.NotifWebsocket) Server {

	// Initialize RabbitMQ
	rabbitMQ, err := rabbitmq.NewRabbitMQ("amqp://guest:guest@localhost:5672/ecommerce_development")

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	fiberApp := fiber.New()
	return &fiberServer{
		app:      fiberApp,
		db:       db,
		notifWs:  notifWs,
		rabbitMQ: rabbitMQ,
	}
}

func (s *fiberServer) Use(args interface{}) {
	s.app.Use(args)
}

func (s *fiberServer) Start() {

	s.app.Use("/ws", ws.AllowUpgrade)
	s.app.Use("/ws/notification/:userId", websocket.New(s.notifWs.HandleNotificationRoom()))
	s.initializNotificationServiceHandler()

	// Start consuming messages
	// go s.startRabbitMQConsumer()

	log.Fatal(s.app.Listen(":5002"))

}

func (s *fiberServer) initializNotificationServiceHandler() {

	// repository
	notifRepo := repository.NewNotificationRepository(s.db)

	// product service repository
	notifeUsecase := usecase.NewNotificationUsecase(notifRepo)

	// handler
	notifHandler := httphandler.NewCartHttpHandle(notifeUsecase, s.notifWs)

	// router
	s.app.Post("/send-notification", notifHandler.InsertNewNotifivation)
	s.app.Get("/send-notification/:userId", notifHandler.GetNotificationByUser)

	go s.startRabbitMQConsumer(notifeUsecase)

}

// Gracefully close RabbitMQ
func (s *fiberServer) Close() {
	if s.rabbitMQ != nil {
		s.rabbitMQ.Close()
	}
}

// RabbitMQ consumer logic
func (s *fiberServer) startRabbitMQConsumer(notifUsecase domain.NotificationUsecase) {

	queueName := "realtime_notif"

	// Inject event handler
	// eventHandler := ws.NewEventHandler(s.notifWs) // This will handle incoming messages

	// err := s.rabbitMQ.ConsumeOrderStatus(queueName, eventHandler)
	err := s.rabbitMQ.ConsumeOrderStatus(queueName, notifUsecase)

	if err != nil {
		log.Fatalf("Failed to start RabbitMQ consumer: %v", err)
	}
}
