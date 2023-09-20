package main

import (
	"fmt"
	"time"
)

// PaymentProcessor represents the interface for payment processing (real subject and proxy).
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// RealPaymentProcessor represents the real payment processor.
type RealPaymentProcessor struct {
	Name string
}

func (rp *RealPaymentProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of $%.2f with %s\n", amount, rp.Name)
	// Simulate the actual payment processing logic here.
}

// PaymentProxy represents a proxy for payment processing, adding validation and logging.
type PaymentProxy struct {
	realProcessor RealPaymentProcessor
}

func (pp *PaymentProxy) ProcessPayment(amount float64) {
	if pp.isValidAmount(amount) {
		// Perform additional logging.
		logEntry := fmt.Sprintf("Payment of $%.2f processed by %s at %s",
			amount, pp.realProcessor.Name, time.Now().Format(time.RFC3339))
		fmt.Println(logEntry)

		// Delegate the payment processing to the real processor.
		pp.realProcessor.ProcessPayment(amount)
	} else {
		fmt.Println("Invalid payment amount.")
	}
}

func (pp *PaymentProxy) isValidAmount(amount float64) bool {
	// Simulate a validation check.
	return amount > 0
}

func main() {
	// Create a real payment processor.
	realProcessor := RealPaymentProcessor{Name: "PaymentGatewayXYZ"}

	// Create a payment proxy for additional processing.
	paymentProxy := PaymentProxy{realProcessor: realProcessor}

	// Process payments through the proxy.
	paymentProxy.ProcessPayment(100.0)
	paymentProxy.ProcessPayment(-50.0)
	paymentProxy.ProcessPayment(250.0)
}
