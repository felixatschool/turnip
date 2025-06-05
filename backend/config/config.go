package config

import (
  "log"
  "os"
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
