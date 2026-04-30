package config

import(
	"log"
	"os"
	"github.com/joho/godotenv"
)
type Config struct {
	Port        string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	JWTSecret   string
	JWTExpiry   string
}

func Load() *Config{
	err := godotenv.Load()
	if(err != nil){
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Port: os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiry: os.Getenv("JWT_EXPIRY"),
	}
}
// getEnv reads a key from environment, returns fallback if not found
func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
func (c *Config) GetJWTSecret() string {
	return c.JWTSecret
}