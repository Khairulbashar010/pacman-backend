package services

import (
    "fmt"
    "log"
    "net/smtp"
    "strings"
    "packman-backend/internal/config"
)

// SendEmail sends a basic email to a recipient
func SendEmail(to []string, subject, body string) error {
    cfg := config.GetConfig()

    // SMTP server configuration.
    smtpHost := cfg.SMTPHost
    smtpPort := cfg.SMTPPort
    smtpUser := cfg.SMTPUser
    smtpPass := cfg.SMTPPassword
    smtpFrom := cfg.SMTPFrom

    // Setup the message
    msg := buildEmailMessage(smtpFrom, to, subject, body)

    // Connect to the SMTP server
    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

    // Send the email
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, to, []byte(msg))
    if err != nil {
        log.Printf("Error sending email: %v", err)
        return err
    }

    log.Printf("Email sent successfully to: %v", to)
    return nil
}

// buildEmailMessage creates the email message format
func buildEmailMessage(from string, to []string, subject, body string) string {
    // The email headers and body
    headers := make(map[string]string)
    headers["From"] = from
    headers["To"] = strings.Join(to, ",")
    headers["Subject"] = subject
    headers["MIME-version"] = "1.0"
    headers["Content-Type"] = "text/html; charset=\"UTF-8\""

    // Combine the headers and body into a single message
    message := ""
    for k, v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    return message
}

// SendWelcomeEmail sends a welcome email to a new user
func SendWelcomeEmail(to string) error {
    subject := "Welcome to Pac-Man Multiplayer!"
    body := `<html>
                <body>
                    <h1>Welcome to the Pac-Man Game!</h1>
                    <p>Thank you for signing up. Enjoy playing and good luck!</p>
                </body>
             </html>`

    return SendEmail([]string{to}, subject, body)
}
