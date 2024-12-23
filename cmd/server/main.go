package main

import (
    "github.com/gin-gonic/gin"
    "mock-aws-ses/internal/api"
)

func main() {
    r := gin.Default()
    api.SetupRoutes(r)
    r.Run(":8080")
}
