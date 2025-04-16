package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	Env            string
	DBHost         string
	DBPort         int
	DBUser         string
	DBPassword     string
	DBName         string
	BaseCurrency   string
	CurrencyAPIURL string
	JWTSecret      string
}

func LoadConfig() *Config {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using system env")
		}
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	return &Config{
		Port:           getEnv("PORT", "8080"),
		Env:            getEnv("ENV", "development"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         dbPort,
		DBUser:         getEnv("DB_USER", ""),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", ""),
		BaseCurrency:   getEnv("BASE_CURRENCY", "USD"),
		CurrencyAPIURL: getEnv("CURRENCY_API_URL", "https://api.exchangerate.host"),
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
