package routes

import (
  "github.com/gin-gonic/gin"
  "turnip/controllers"
)

func RegisterRoutes(r *gin.Engine) {
  api := r.Group("/api")
  {
    api.GET("/health", controllers.HealthCheck)
  }
}
