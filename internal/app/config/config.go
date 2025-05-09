package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	Port       string
	JWTSecret  string
	JWTExpire  int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := &Config{}
	return &Config{
		DBUser:     c.Get("DB_USER", ""),
		DBPassword: c.Get("DB_PASSWORD", ""),
		DBHost:     c.Get("DB_HOST", ""),
		DBPort:     c.Get("DB_PORT", ""),
		DBName:     c.Get("DB_NAME", ""),
		Port:       c.Get("PORT", ""),
		JWTSecret:  c.Get("JWT_SECRET", ""),
		JWTExpire:  c.GetInt("JWT_EXPIRE", 1),
	}
}

func (c *Config) Get(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Printf("Failed to get %s, using default value: %s", key, def)
	return def
}

func (c *Config) GetInt(key string, def int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Failed to parse %s to int, using default value: %d", key, def)
		return def
	}

	return value
}