package main

import (
	"fmt"
)

// Component represents a generic payment component.
type Component interface {
	ProcessPayment() string
}

// IndividualPayment represents an individual payment.
type IndividualPayment struct {
	ID     string
	Amount float64
}

func (ip *IndividualPayment) ProcessPayment() string {
	return fmt.Sprintf("Processing payment ID: %s, Amount: $%.2f", ip.ID, ip.Amount)
}

// PaymentBatch represents a batch of payments.
type PaymentBatch struct {
	ID       string
	Payments []Component
}

func (pb *PaymentBatch) ProcessPayment() string {
	var result string
	for _, payment := range pb.Payments {
		result += payment.ProcessPayment() + "\n"
	}
	return result
}

func main() {
	// Create individual payment transactions
	payment1 := &IndividualPayment{ID: "12345", Amount: 100.0}
	payment2 := &IndividualPayment{ID: "67890", Amount: 50.0}

	// Create a batch of payments
	batch := &PaymentBatch{
		ID:       "Batch001",
		Payments: []Component{payment1, payment2},
	}

	// Process individual payments and payment batch
	fmt.Println("Individual Payments:")
	fmt.Println(payment1.ProcessPayment())
	fmt.Println(payment2.ProcessPayment())

	fmt.Println("\nPayment Batch:")
	fmt.Println(batch.ProcessPayment())
}
