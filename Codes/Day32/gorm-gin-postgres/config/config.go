package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "arvinder.pal"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "gorm_gin_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	// Debug prints
	log.Printf("DEBUG: DB_HOST = %q", cfg.DBHost)
	log.Printf("DEBUG: DB_PORT = %q", cfg.DBPort)
	log.Printf("DEBUG: DB_USER = %q", cfg.DBUser)
	log.Printf("DEBUG: DB_PASSWORD = %q", cfg.DBPassword)
	log.Printf("DEBUG: DB_NAME = %q", cfg.DBName)
	log.Printf("DEBUG: DB_SSLMODE = %q", cfg.DBSSLMode)
	log.Printf("DEBUG: SERVER_PORT = %q", cfg.ServerPort)

	return cfg
}
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue //hhhhhh
}
