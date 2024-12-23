package api

import (
    "github.com/gin-gonic/gin"
    "mock-aws-ses/internal/service"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/send-email", service.SendEmail)
    r.GET("/stats", service.GetStats)
}
