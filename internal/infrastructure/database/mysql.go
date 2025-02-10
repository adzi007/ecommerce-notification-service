package database

import (
	"fmt"
	"time"

	"github.com/adzi007/ecommerce-notification-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var dbInstance *postgresDatabase

func NewDatabase() Database {

	// db, err := gorm.Open(sqlite.Open("database/notifications.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	// dbInstance = &postgresDatabase{Db: db}
	// return dbInstance

	dbUsername := config.ENV.DB_USERNAME
	dbPassword := config.ENV.DB_PASSWORD
	dbName := config.ENV.DB_NAME
	dbHost := config.ENV.DB_HOST
	dbPort := config.ENV.DB_PORT

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	fmt.Println("connectionString >>> ", connectionString)

	var db *gorm.DB
	var err error

	// Retry connecting to the database up to 10 times
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
		})
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to the database. Retrying in 2 seconds... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic("failed to connect database after 10 attempts")
	}

	dbInstance = &postgresDatabase{Db: db}
	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
