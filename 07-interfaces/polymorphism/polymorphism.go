package main

import (
	"fmt"
	"time"
)

// Payment interface - different payment methods implement this
type Payment interface {
	ProcessPayment(amount float64) string
	GetPaymentMethod() string
}

// CreditCard payment method
type CreditCard struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing $%.2f via Credit Card ending in %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

func (cc CreditCard) GetPaymentMethod() string {
	return "Credit Card"
}

// PayPal payment method
type PayPal struct {
	Email string
}

func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing $%.2f via PayPal account: %s", amount, pp.Email)
}

func (pp PayPal) GetPaymentMethod() string {
	return "PayPal"
}

// Cryptocurrency payment method
type Cryptocurrency struct {
	WalletAddress string
	CoinType      string
}

func (crypto Cryptocurrency) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing $%.2f via %s to wallet: %s", amount, crypto.CoinType, crypto.WalletAddress[:10]+"...")
}

func (crypto Cryptocurrency) GetPaymentMethod() string {
	return fmt.Sprintf("Cryptocurrency (%s)", crypto.CoinType)
}

// Notification interface - different notification channels
type Notification interface {
	Send(message string) error
}

// EmailNotification sends notifications via email
type EmailNotification struct {
	To      string
	From    string
	Subject string
}

func (e EmailNotification) Send(message string) error {
	fmt.Printf("[EMAIL] To: %s | Subject: %s | Message: %s\n", e.To, e.Subject, message)
	return nil
}

// SMSNotification sends notifications via SMS
type SMSNotification struct {
	PhoneNumber string
}

func (s SMSNotification) Send(message string) error {
	fmt.Printf("[SMS] To: %s | Message: %s\n", s.PhoneNumber, message)
	return nil
}

// PushNotification sends push notifications
type PushNotification struct {
	DeviceID string
	AppName  string
}

func (p PushNotification) Send(message string) error {
	fmt.Printf("[PUSH] App: %s | Device: %s | Message: %s\n", p.AppName, p.DeviceID, message)
	return nil
}

// Logger interface - different logging implementations
type Logger interface {
	Log(message string)
	LogError(err error)
}

// ConsoleLogger logs to console
type ConsoleLogger struct {
	Prefix string
}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] %s: %s\n", time.Now().Format("15:04:05"), cl.Prefix, message)
}

func (cl ConsoleLogger) LogError(err error) {
	fmt.Printf("[%s] ERROR - %s: %v\n", time.Now().Format("15:04:05"), cl.Prefix, err)
}

// FileLogger simulates logging to a file
type FileLogger struct {
	FileName string
}

func (fl FileLogger) Log(message string) {
	fmt.Printf("[FILE: %s] %s: %s\n", fl.FileName, time.Now().Format("2006-01-02 15:04:05"), message)
}

func (fl FileLogger) LogError(err error) {
	fmt.Printf("[FILE: %s] ERROR %s: %v\n", fl.FileName, time.Now().Format("2006-01-02 15:04:05"), err)
}

// processPayment demonstrates polymorphism with Payment interface
func processPayment(p Payment, amount float64) {
	fmt.Printf("\nPayment Method: %s\n", p.GetPaymentMethod())
	fmt.Println(p.ProcessPayment(amount))
}

// sendNotification demonstrates polymorphism with Notification interface
func sendNotification(n Notification, message string) {
	n.Send(message)
}

// logActivity demonstrates polymorphism with Logger interface
func logActivity(logger Logger, activity string) {
	logger.Log(activity)
}

// OrderProcessor demonstrates using multiple interfaces together
type OrderProcessor struct {
	payment      Payment
	notification Notification
	logger       Logger
}

func (op OrderProcessor) ProcessOrder(orderID string, amount float64) {
	op.logger.Log(fmt.Sprintf("Starting order processing for Order ID: %s", orderID))

	// Process payment
	result := op.payment.ProcessPayment(amount)
	op.logger.Log(result)

	// Send notification
	notifMsg := fmt.Sprintf("Your order %s has been processed successfully!", orderID)
	op.notification.Send(notifMsg)

	op.logger.Log(fmt.Sprintf("Order %s completed", orderID))
}

func main() {
	fmt.Println("=== Polymorphism in Go ===")

	// 1. Payment Polymorphism
	fmt.Println("\n--- Payment Processing ---")
	creditCard := CreditCard{CardNumber: "1234567890123456", CVV: "123", ExpiryDate: "12/25"}
	paypal := PayPal{Email: "user@example.com"}
	crypto := Cryptocurrency{WalletAddress: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", CoinType: "Bitcoin"}

	processPayment(creditCard, 150.00)
	processPayment(paypal, 75.50)
	processPayment(crypto, 200.00)

	// 2. Notification Polymorphism
	fmt.Println("\n--- Notification System ---")
	email := EmailNotification{To: "customer@example.com", From: "noreply@shop.com", Subject: "Order Confirmation"}
	sms := SMSNotification{PhoneNumber: "+1-555-0123"}
	push := PushNotification{DeviceID: "device-12345", AppName: "ShopApp"}

	sendNotification(email, "Your order has been shipped!")
	sendNotification(sms, "Your package will arrive tomorrow")
	sendNotification(push, "Flash sale: 50% off!")

	// 3. Logger Polymorphism
	fmt.Println("\n--- Logging System ---")
	consoleLog := ConsoleLogger{Prefix: "APP"}
	fileLog := FileLogger{FileName: "app.log"}

	logActivity(consoleLog, "User logged in")
	logActivity(fileLog, "Database connection established")

	// 4. Combining Multiple Interfaces
	fmt.Println("\n--- Order Processing System ---")
	processor := OrderProcessor{
		payment:      creditCard,
		notification: email,
		logger:       consoleLog,
	}
	processor.ProcessOrder("ORD-2024-001", 299.99)

	// 5. Slice of different types implementing same interface
	fmt.Println("\n--- Batch Payment Processing ---")
	payments := []Payment{creditCard, paypal, crypto}
	totalAmount := 0.0
	for i, payment := range payments {
		amount := float64((i + 1) * 50)
		totalAmount += amount
		fmt.Printf("\nTransaction %d:\n", i+1)
		processPayment(payment, amount)
	}
	fmt.Printf("\nTotal processed: $%.2f\n", totalAmount)

	// 6. Dynamic behavior based on interface type
	fmt.Println("\n--- Multi-Channel Notifications ---")
	notifications := []Notification{email, sms, push}
	message := "Important system update available"
	for _, notif := range notifications {
		sendNotification(notif, message)
	}
}
