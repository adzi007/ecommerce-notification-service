package main

import (
	"github.com/adzi007/ecommerce-notification-service/config"
	"github.com/adzi007/ecommerce-notification-service/internal/domain"
	"github.com/adzi007/ecommerce-notification-service/internal/infrastructure/database"
)

func main() {
	config.LoadConfig()
	db := database.NewDatabase()
	appDbMigrate(db)
}

func appDbMigrate(db database.Database) {

	// err := db.GetDb().Migrator().CreateTable(&domain.Notification{})
	err := db.GetDb().Migrator().AutoMigrate(&domain.Notification{})

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
