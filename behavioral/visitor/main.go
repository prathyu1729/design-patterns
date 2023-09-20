package main

import "fmt"

func main() {
	// Create payment methods (credit card and PayPal).
	creditCardPayment := &CreditCard{Amount: 100.0}
	paypalPayment := &PayPal{Amount: 50.0}

	// Create a fee calculator (visitor).
	feeCalculator := &FeeCalculator{}

	// Accept the visitor on payment methods to calculate fees.
	creditCardPayment.Accept(feeCalculator)
	paypalPayment.Accept(feeCalculator)

	// Print the total fees.
	fmt.Printf("Total fees: $%.2f\n", feeCalculator.TotalFees)
}
