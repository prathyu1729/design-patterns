package singleton

import "fmt"

func main() {
	// Access the transaction logger singleton instance
	logger := GetTransactionLogger()

	// Log some transactions
	logger.LogTransaction("Transaction 1: Deposit $1000")
	logger.LogTransaction("Transaction 2: Withdraw $500")

	// Access the same instance from another part of the application
	anotherLogger := GetTransactionLogger()

	// Log another transaction
	anotherLogger.LogTransaction("Transaction 3: Transfer $200")

	// Display the transaction log
	log := logger.GetLog()
	fmt.Println("Transaction Log:")
	for _, entry := range log {
		fmt.Println(entry)
	}
}
