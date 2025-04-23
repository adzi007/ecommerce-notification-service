package server

import (
	"fmt"
	"log"

	"github.com/adzi007/ecommerce-notification-service/config"
	httphandler "github.com/adzi007/ecommerce-notification-service/internal/delivery/http_handler"
	"github.com/adzi007/ecommerce-notification-service/internal/delivery/ws"
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/monitoring"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/rabbitmq"
	"github.com/adzi007/ecommerce-notification-service/internal/repository"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase/broadcaster"
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

	host := config.ENV.RABBITMQ_HOST_URL
	port := config.ENV.RABBITMQ_PORT
	user := config.ENV.RABBITMQ_USER
	password := config.ENV.RABBITMQ_PASSWORD
	vhost := config.ENV.RABBITMQ_VIRTUAL_HOST

	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, vhost)

	// Initialize RabbitMQ
	// rabbitMQ, err := rabbitmq.NewRabbitMQ("amqp://guest:guest@localhost:5672/ecommerce_development")
	rabbitMQ, err := rabbitmq.NewRabbitMQ(rabbitMQURL)

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

	portApp := config.ENV.PORT_APP

	log.Fatal(s.app.Listen(":" + portApp))

}

func (s *fiberServer) initializNotificationServiceHandler() {

	// repository
	notifRepo := repository.NewNotificationRepository(s.db)

	// Broadcast Usecase
	broadcastUsecase := broadcaster.NewBroadcaster(s.notifWs)

	// product service repository
	notifeUsecase := usecase.NewNotificationUsecase(notifRepo, broadcastUsecase)

	// handler
	notifHandler := httphandler.NewCartHttpHandle(notifeUsecase, s.notifWs)

	// Add metrics route
	s.app.Get("/metrics", monitoring.MetricsHandler())

	// router
	s.app.Post("/send-notification", notifHandler.InsertNewNotifivation)
	s.app.Get("/notification/:userId", notifHandler.GetNotificationByUser)

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
	err := s.rabbitMQ.ConsumeOrderStatus(queueName, notifUsecase)

	if err != nil {
		log.Fatalf("Failed to start RabbitMQ consumer: %v", err)
	}
}
