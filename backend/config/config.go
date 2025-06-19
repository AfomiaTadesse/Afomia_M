package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	MongoURI  string
	JWTSecret string
	Port      string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		JWTSecret: getEnv("JWT_SECRET", "default-secret-key"), 
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}