package main

import "fmt"

// PaymentProcessor represents a payment processor (observer) that receives payment updates.
type PaymentProcessor struct {
	name string
}

func NewPaymentProcessor(name string) *PaymentProcessor {
	return &PaymentProcessor{name: name}
}

func (pp *PaymentProcessor) Update(paymentID string, status string) {
	fmt.Printf("Payment processor %s received an update for payment %s: %s\n", pp.name, paymentID, status)
}
