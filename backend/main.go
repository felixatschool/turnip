package main

import (
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "turnip/config"
  "turnip/routes"
)

func main() {
  config.LoadEnv()

  r := gin.Default()

  r.Use(cors.New(cors.Config{
    AllowOrigins:     config.GetAllowedOrigins(),
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type"},
    AllowCredentials: true,
  }))

  routes.RegisterRoutes(r)

  port := config.GetPort()
  r.Run(":" + port)
}

