package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type EmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type APIStats struct {
	TotalEmailsSent int            `json:"total_emails_sent"`
	EmailUsage      map[string]int `json:"email_usage"`
	Mutex           sync.Mutex
}

var stats = APIStats{
	EmailUsage: make(map[string]int),
}

func sendEmail(c *gin.Context) {
	var emailRequest EmailRequest
	if err := c.ShouldBindJSON(&emailRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email request payload"})
		return
	}

	stats.Mutex.Lock()
	stats.TotalEmailsSent++
	for _, recipient := range emailRequest.To {
		stats.EmailUsage[recipient]++
	}
	stats.Mutex.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"message": "email request processed successfully",
	})
}

func getStats(c *gin.Context) {
	stats.Mutex.Lock()
	response := APIStats{
		TotalEmailsSent: stats.TotalEmailsSent,
		EmailUsage:      stats.EmailUsage,
	}
	stats.Mutex.Unlock()

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.POST("/send-email", sendEmail)
	r.GET("/stats", getStats)
	r.Run(":8080")
}
