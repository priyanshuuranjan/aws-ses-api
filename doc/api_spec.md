# Mock AWS SES API Specification

## Endpoint: POST /send-email

### Request Body
```json
{
  "from": "sender@example.com",
  "to": ["recipient1@example.com", "recipient2@example.com"],
  "subject": "Test Email",
  "body": "This is a test email"
}
