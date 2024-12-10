package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	APIKey     string
}

// LoadConfig loads environment variables from the `.env` file
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default settings")
	}

	return Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		APIKey:     getEnv("API_KEY", "apitest"),
	}
}

// Helper function to get environment variables with a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
