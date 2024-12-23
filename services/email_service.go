package service

import (
    "net/http"
    "mock-aws-ses/internal/models"
    "sync"

    "github.com/gin-gonic/gin"
)

var emailStats = models.EmailStats{}
var mu sync.Mutex

func SendEmail(c *gin.Context) {
    var emailRequest models.EmailRequest
    if err := c.ShouldBindJSON(&emailRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    mu.Lock()
    emailStats.TotalSent++
    mu.Unlock()

    c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

func GetStats(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"stats": emailStats})
}
