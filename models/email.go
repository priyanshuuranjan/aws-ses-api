package models

type EmailRequest struct {
    From      string   `json:"from"`
    To        []string `json:"to"`
    Subject   string   `json:"subject"`
    Body      string   `json:"body"`
}

type EmailStats struct {
    TotalSent    int `json:"total_sent"`
    InvalidEmails int `json:"invalid_emails"`
}
