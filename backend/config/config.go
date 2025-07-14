package config

import (
  "log"
  "os"
  "strings"
  "github.com/joho/godotenv"
)

func LoadEnv() {
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file found, using default values")
  }
}

func GetPort() string {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  return port
}

func GetAllowedOrigins() []string {
  origins := os.Getenv("ALLOWED_ORIGINS")
  if origins == "" {
    origins = "http://localhost:5173"
  }
  return strings.Split(origins, ",")
}
