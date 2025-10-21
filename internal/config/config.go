package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	DBSSLMode string
	ServerPort string
	DB DbConfig
}

type DbConfig struct {
	MaxOpenConns int
	MaxIdleConns int
}

func Load() (*Config, error){
	if err := godotenv.Load(); err != nil {
		log.Println(".env was nowhtere to be found");
	}

	config := &Config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5433"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName: getEnv("DB_NAME", "employee_db"),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DB: DbConfig{
			MaxOpenConns: 25,
			MaxIdleConns: 5,
		 },
	}

	return config, nil
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, 
		c.DBPort, 
		c.DBUser, 
		c.DBPassword, 
		c.DBName, 
		c.DBSSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}