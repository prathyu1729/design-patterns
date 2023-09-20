package main

import "fmt"

// PaymentProcessor represents the receiver that knows how to process payments.
type PaymentProcessor struct {
	accountBalance map[string]float64
}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		accountBalance: make(map[string]float64),
	}
}

func (p *PaymentProcessor) ExecutePayment(accountID string, amount float64) {
	if balance, ok := p.accountBalance[accountID]; ok && balance >= amount {
		p.accountBalance[accountID] -= amount
		fmt.Printf("Payment of $%.2f processed for account %s. New balance: $%.2f\n", amount, accountID, p.accountBalance[accountID])
	} else {
		fmt.Printf("Payment of $%.2f for account %s failed. Insufficient balance.\n", amount, accountID)
	}
}
