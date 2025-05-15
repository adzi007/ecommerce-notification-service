package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST               string
	DB_USERNAME           string
	DB_PASSWORD           string
	DB_PORT               string
	DB_NAME               string
	PORT_APP              string
	API_GATEWAY           string
	URL_PRODUCT_SERVICE   string
	RABBITMQ_HOST_URL     string
	RABBITMQ_PORT         string
	RABBITMQ_USER         string
	RABBITMQ_PASSWORD     string
	RABBITMQ_VIRTUAL_HOST string
}

var (
	ENV        Config
	_, b, _, _ = runtime.Caller(0)

	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

func LoadConfig() {
	// Load .env file only in development
	if os.Getenv("ENV") != "production" {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading .env file: %v", err)
		}
	}

	// Automatically override with environment variables
	viper.AutomaticEnv()

	// Bind environment variables to struct
	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatalf("Failed to unmarshal env vars: %v", err)
	}

	log.Println("Config loaded successfully")
}

// Connect to SQLite Database
// func InitDB() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("database/notifications.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
// 	return db
// }
