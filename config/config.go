package config

import (
	"os"

	"github.com/AlexDeKatz/banking/logging"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURI string
	ServerPort  string
	ServerHost  string
}

var dbURI string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		logging.Error("Error loading .env file")
	}
	dbURI = os.Getenv("DB_URI")

}
func GetConfig() *Config {
	return &Config{
		DatabaseURI: dbURI,
		ServerPort:  os.Getenv("SERVER_PORT"),
		ServerHost:  os.Getenv("SERVER_HOST"),
	}
}
