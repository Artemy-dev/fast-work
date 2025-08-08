// Описываем как работать с переменными окружения

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[-] No .env file.")
		return
	}
	log.Println("[+] .env file loaded.")
}

type DatabaseConfig struct {
	Url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Url: os.Getenv("DATABASE_URL"),
	}
}