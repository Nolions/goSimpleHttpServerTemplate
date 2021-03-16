package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	App App
	DB  Database
}

// Application 相關設定
type App struct {
	Debug bool
	Port  string
}

// Database 相關設定
type Database struct {
	Driver    string
	Host      string
	Port      string
	Database  string
	User      string
	Password  string
	Charset   string
	Collation string
}

var Conf AppConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// 讀取環境變數中設定檔
func Load() {
	Conf.App = App{
		Port:  os.Getenv("APP_PORT"),
		Debug: getEnvBool("APP_DEBUG"),
	}
	Conf.DB = Database{
		Driver:    os.Getenv("DB_DRIVER"),
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Database:  os.Getenv("DB_DATABASE"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		Charset:   os.Getenv("DB_CHARSET"),
		Collation: os.Getenv("DB_COLLATION"),
	}
}

// 環境變數中bool值
func getEnvBool(key string) bool {
	s := os.Getenv(key)

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return v
}
