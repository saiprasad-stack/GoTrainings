package main

import (
	"fmt"
)

// Notifier is a basic interface for sending notifications
// Pattern inspired by Docker's logging drivers
type Notifier interface {
	Notify(message string) error
	GetType() string
}

// EmailNotifier implements the Notifier interface
type EmailNotifier struct {
	EmailAddress string
}

func (e *EmailNotifier) Notify(message string) error {
	fmt.Printf("[Email] Sending to %s: %s\n", e.EmailAddress, message)
	return nil
}

func (e *EmailNotifier) GetType() string {
	return "email"
}

// SlackNotifier implements the Notifier interface
type SlackNotifier struct {
	WebhookURL string
	Channel    string
}

func (s *SlackNotifier) Notify(message string) error {
	fmt.Printf("[Slack] Sending to #%s: %s\n", s.Channel, message)
	return nil
}

func (s *SlackNotifier) GetType() string {
	return "slack"
}

// SMSNotifier implements the Notifier interface
type SMSNotifier struct {
	PhoneNumber string
}

func (sms *SMSNotifier) Notify(message string) error {
	fmt.Printf("[SMS] Sending to %s: %s\n", sms.PhoneNumber, message)
	return nil
}

func (sms *SMSNotifier) GetType() string {
	return "sms"
}

// AlertSystem accepts any Notifier
func AlertSystem(notifier Notifier, message string) {
	fmt.Printf("Using %s notifier: ", notifier.GetType())
	notifier.Notify(message)
}

func main() {
	// Create different notifiers
	email := &EmailNotifier{EmailAddress: "admin@example.com"}
	slack := &SlackNotifier{WebhookURL: "https://hooks.slack.com/xxx", Channel: "alerts"}
	sms := &SMSNotifier{PhoneNumber: "+1234567890"}

	fmt.Println("=== Basic Interface Example ===")

	// Use different notifiers with the same function
	AlertSystem(email, "Server is running low on memory")
	AlertSystem(slack, "Deployment completed successfully")
	AlertSystem(sms, "Critical: Database is down")

	// Store multiple notifiers in a slice
	notifiers := []Notifier{email, slack, sms}

	fmt.Println("\nBulk notifications:")
	for _, n := range notifiers {
		n.Notify("System health check passed")
	}
}
