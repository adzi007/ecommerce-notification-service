package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	// "os"
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
	ENV             Config
	_, b, _, _      = runtime.Caller(0)
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../")
)

func LoadConfig() {
	viper.AutomaticEnv()

	// Explicitly bind environment variables
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_USERNAME")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("PORT_APP")
	viper.BindEnv("API_GATEWAY")
	viper.BindEnv("URL_PRODUCT_SERVICE")
	viper.BindEnv("RABBITMQ_HOST_URL")
	viper.BindEnv("RABBITMQ_PORT")
	viper.BindEnv("RABBITMQ_USER")
	viper.BindEnv("RABBITMQ_PASSWORD")
	viper.BindEnv("RABBITMQ_VIRTUAL_HOST")

	ENV = Config{
		DB_HOST:               viper.GetString("DB_HOST"),
		DB_USERNAME:           viper.GetString("DB_USERNAME"),
		DB_PASSWORD:           viper.GetString("DB_PASSWORD"),
		DB_PORT:               viper.GetString("DB_PORT"),
		DB_NAME:               viper.GetString("DB_NAME"),
		PORT_APP:              viper.GetString("PORT_APP"),
		API_GATEWAY:           viper.GetString("API_GATEWAY"),
		URL_PRODUCT_SERVICE:   viper.GetString("URL_PRODUCT_SERVICE"),
		RABBITMQ_HOST_URL:     viper.GetString("RABBITMQ_HOST_URL"),
		RABBITMQ_PORT:         viper.GetString("RABBITMQ_PORT"),
		RABBITMQ_USER:         viper.GetString("RABBITMQ_USER"),
		RABBITMQ_PASSWORD:     viper.GetString("RABBITMQ_PASSWORD"),
		RABBITMQ_VIRTUAL_HOST: viper.GetString("RABBITMQ_VIRTUAL_HOST"),
	}

	log.Println("os >>> DB_HOST:", os.Getenv("DB_HOST"))
	log.Println("os >>> DB_USERNAME:", os.Getenv("DB_USERNAME"))

	log.Println("✅ Successfully loaded environment variables BindEnv")

}

// Connect to SQLite Database
// func InitDB() *gorm.DB {
// 	db, err := gorm.Open(sqlite.Open("database/notifications.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
// 	return db
// }
