package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port     int
	Database struct {
		Driver   string
		Name     string
		Address  string
		Port     int
		Username string
		Password string
	}
	GoogleClientID     string
	GoogleClientSecret string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	err := godotenv.Load("local.env")

	if err != nil {
		log.Info(err)
	}
	portDB, errParse := strconv.Atoi(os.Getenv("DB_PORT"))
	if errParse != nil {
		log.Warn(errParse)
	}

	var defaultConfig AppConfig
	port, errPort := strconv.Atoi(os.Getenv("APP_PORT"))
	if errPort != nil {
		log.Warn(errParse)
	}

	defaultConfig.Port = port
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = "localhost"
	defaultConfig.Database.Port = portDB
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

	defaultConfig.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	defaultConfig.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	log.Info(defaultConfig)

	return &defaultConfig
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "user" {
		fmt.Println(value)
		return value
	}

	return fallback
}
