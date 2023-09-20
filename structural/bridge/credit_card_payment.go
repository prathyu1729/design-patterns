package main

import "fmt"

// CreditCardPayment is a concrete implementation of PaymentMethod.
type CreditCardPayment struct{}

func (p *CreditCardPayment) ProcessPayment(amount float64) (string, error) {
	return "Credit Card payment processed for $" + fmt.Sprintf("%.2f", amount), nil
}

// BankTransferPayment is another concrete implementation of PaymentMethod.
type BankTransferPayment struct{}

func (p *BankTransferPayment) ProcessPayment(amount float64) (string, error) {
	return "Bank Transfer payment processed for $" + fmt.Sprintf("%.2f", amount), nil
}

// DigitalWalletPayment is yet another concrete implementation of PaymentMethod.
type DigitalWalletPayment struct{}

func (p *DigitalWalletPayment) ProcessPayment(amount float64) (string, error) {
	return "Digital Wallet payment processed for $" + fmt.Sprintf("%.2f", amount), nil
}
