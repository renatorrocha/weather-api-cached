package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	OPENWEATHER_API_KEY string
}

var appConfig *Config

func Load() {
	err := godotenv.Load(filepath.Join("..", ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig = &Config{
		OPENWEATHER_API_KEY: GetEnv("OPENWEATHER_API_KEY", ""),
	}
}

// return env value or default value if not found
func GetEnv(key string, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}

	return defaultValue
}
