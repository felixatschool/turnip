package main

import (
  "github.com/gin-gonic/gin"
  "turnip/config"
  "turnip/routes"
)

func main() {
  config.LoadEnv()

  r := gin.Default()

  routes.RegisterRoutes(r)

  port := config.GetPort()
  r.Run(":" + port)
}
