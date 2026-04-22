package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	DB DBConfig
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	return &Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_USER", ""),
			DBName:   getEnv("DB_NAME", ""),
			SSLMode:  getEnv("SSLMODE", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
