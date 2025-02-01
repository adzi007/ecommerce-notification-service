package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var dbInstance *postgresDatabase

func NewDatabase() Database {

	db, err := gorm.Open(sqlite.Open("database/notifications.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	dbInstance = &postgresDatabase{Db: db}
	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
