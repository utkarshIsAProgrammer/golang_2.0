package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// structure for the database config
type Config struct {
	DatabseURL string
	Port       string
}

// function to load the configuration from the .env file
func Load() (*Config, error) {
	var err = godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables.")
	}

	// get data from the .env file
	var config *Config = &Config{
		DatabseURL: os.Getenv("DATABASE_URL"),
		Port:       os.Getenv("PORT"),
	}

	return config, nil
}
