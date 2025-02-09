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
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("Load server successfully")
}

// Connect to SQLite Database
// func InitDB() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("database/notifications.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
// 	return db
// }
