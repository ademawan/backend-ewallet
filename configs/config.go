package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port     int
	Database struct {
		Driver   string
		Name     string
		Address  string
		Port     string
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

	var defaultConfig AppConfig
	port, errPort := strconv.Atoi(os.Getenv("PORT"))
	if errPort != nil {
		log.Warn(errPort)
	}

	defaultConfig.Port = port
	defaultConfig.Database.Driver = os.Getenv("DRIVER")
	defaultConfig.Database.Name = os.Getenv("NAME")
	defaultConfig.Database.Address = "ADDRESS"
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("USERNAME")
	defaultConfig.Database.Password = os.Getenv("PASSWORD")

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
