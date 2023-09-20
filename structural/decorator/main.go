package main

import "fmt"

// Transaction interface defines the basic behavior of a financial transaction.
type Transaction interface {
	Execute() string
}

// SimpleTransfer represents a basic money transfer.
type SimpleTransfer struct {
	Amount float64
}

func (st SimpleTransfer) Execute() string {
	return fmt.Sprintf("Transfer %.2f dollars", st.Amount)
}

// SecureTransferDecorator is a decorator that adds security to a transaction.
type SecureTransferDecorator struct {
	Transaction Transaction
}

func (std SecureTransferDecorator) Execute() string {
	// Add security logic here, for example, encrypting the transaction.
	return fmt.Sprintf("Secure %s", std.Transaction.Execute())
}

func main() {
	// Create a simple transfer
	simpleTransfer := SimpleTransfer{Amount: 100.00}
	fmt.Println("Simple Transfer:", simpleTransfer.Execute())

	// Wrap the simple transfer with security using the decorator
	secureTransfer := SecureTransferDecorator{Transaction: simpleTransfer}
	fmt.Println("Secure Transfer:", secureTransfer.Execute())
}
