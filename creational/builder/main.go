package builder

import "fmt"

func main() {
	// Create a simple deposit transaction
	simpleDeposit := NewSimpleTransactionBuilder().
		SetTransactionType("Deposit").
		SetAmount(1000.0).
		SetAccountID("12345").
		SetDescription("Initial deposit").
		Build()

	// Create a simple withdrawal transaction
	simpleWithdrawal := NewSimpleTransactionBuilder().
		SetTransactionType("Withdrawal").
		SetAmount(500.0).
		SetAccountID("12345").
		SetDescription("ATM withdrawal").
		Build()

	// Display the transactions
	displayTransaction(simpleDeposit)
	displayTransaction(simpleWithdrawal)
}

func displayTransaction(transaction *Transaction) {
	fmt.Printf("Transaction Type: %s\n", transaction.TransactionType)
	fmt.Printf("Amount: $%.2f\n", transaction.Amount)
	fmt.Printf("Account ID: %s\n", transaction.AccountID)
	fmt.Printf("Description: %s\n", transaction.Description)
	fmt.Println()
}
