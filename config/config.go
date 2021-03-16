package config

import (
"github.com/joho/godotenv"
"log"
"os"
"strconv"
)

type AppConfig struct {
	Debug bool
	Port  string
}

var Conf AppConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Load() {
	Conf.Port = os.Getenv("APP_PORT")
	Conf.Debug = getEnvBool("APP_DEBUG")
}

func getEnvBool(key string) bool {
	s := os.Getenv(key)

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return v
}