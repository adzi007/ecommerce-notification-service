package server

import (
	"log"

	httphandler "github.com/adzi007/ecommerce-notification-service/internal/delivery/http_handler"
	"github.com/adzi007/ecommerce-notification-service/internal/delivery/ws"
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/database"
	"github.com/adzi007/ecommerce-notification-service/internal/repository"
	"github.com/adzi007/ecommerce-notification-service/internal/usecase"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app     *fiber.App
	db      database.Database
	notifWs domain.NotifWebsocket
	// conf *config.Config
}

func NewFiberServer(db database.Database, notifWs domain.NotifWebsocket) Server {
	fiberApp := fiber.New()
	// fiberApp.Logger.SetLevel(log.DEBUG)

	// fiberApp.Get("/docs/*", swagger.HandlerDefault)

	return &fiberServer{
		app:     fiberApp,
		db:      db,
		notifWs: notifWs,
	}
}

func (s *fiberServer) Use(args interface{}) {

	s.app.Use(args)

}

func (s *fiberServer) Start() {

	s.app.Use("/ws", ws.AllowUpgrade)
	s.app.Use("/ws/notification/:userId", websocket.New(s.notifWs.HandleNotificationRoom()))

	s.initializeCartServiceHttpHandler()

	log.Fatal(s.app.Listen(":5000"))

}

func (s *fiberServer) initializeCartServiceHttpHandler() {

	// ctx := context.Background()

	// redisRepo := cachestore.NewRedisCache(ctx, "localhost:6379", "", 0)

	// repository
	notifRepo := repository.NewNotificationRepository(s.db)

	// product service repository

	notifeUsecase := usecase.NewNotificationUsecase(notifRepo)

	// handler
	notifHandler := httphandler.NewCartHttpHandle(notifeUsecase, s.notifWs)

	// lis, err := net.Listen("tcp", ":9001")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()

	// pb.RegisterCartServiceServer(grpcServer, &localGrpc.CartGrpcHandler{})

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }

	// router
	s.app.Post("/send-notification", notifHandler.InsertNewNotifivation)

}
