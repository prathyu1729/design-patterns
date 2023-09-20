package main

import (
	"fmt"
)

// LegacyPaymentGateway represents the legacy payment gateway with its own interface.
type LegacyPaymentGateway struct{}

// ProcessPaymentLegacy is a legacy method for processing payments.
func (l *LegacyPaymentGateway) ProcessPaymentLegacy(amount float64) {
	fmt.Printf("Processing payment of $%.2f using the legacy payment gateway\n", amount)
}

// NewFintechPaymentGateway is the new fintech payment gateway interface.
type NewFintechPaymentGateway interface {
	MakePayment(amount float64)
}

// FintechPaymentGatewayAdapter is an adapter that makes the legacy payment gateway compatible with the new fintech interface.
type FintechPaymentGatewayAdapter struct {
	LegacyGateway *LegacyPaymentGateway
}

// MakePayment is a method of the fintech payment gateway interface.
func (a *FintechPaymentGatewayAdapter) MakePayment(amount float64) {
	a.LegacyGateway.ProcessPaymentLegacy(amount)
}

func main() {
	// Create an instance of the legacy payment gateway
	legacyGateway := &LegacyPaymentGateway{}

	// Create an instance of the fintech payment gateway adapter
	adapter := &FintechPaymentGatewayAdapter{LegacyGateway: legacyGateway}

	// Use the fintech payment gateway interface to make a payment
	adapter.MakePayment(200.0)
}
