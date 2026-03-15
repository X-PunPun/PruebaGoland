package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBConnString string
	RawgAPIKey   string
	RawgBaseURL  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:         getEnv("PORT", "8080"),
		DBConnString: getEnv("DB_CONNECTION", ""),
		RawgAPIKey:   getEnv("RAWG_API_KEY", ""),
		RawgBaseURL:  getEnv("RAWG_BASE_URL", "https://api.rawg.io/api"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
