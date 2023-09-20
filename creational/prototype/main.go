package prototype

import "fmt"

func main() {
	// Create a prototype financial transaction
	prototypeTransaction := &FinancialTransaction{
		TransactionType: "Deposit",
		Amount:          1000.0,
		AccountID:       "12345",
		Description:     "Initial deposit",
	}

	// Clone the prototype to create new transactions
	transaction1 := prototypeTransaction.Clone().(*FinancialTransaction)
	transaction1.TransactionType = "Withdrawal"
	transaction1.Amount = 500.0
	transaction1.Description = "ATM withdrawal"

	transaction2 := prototypeTransaction.Clone().(*FinancialTransaction)
	transaction2.Amount = 200.0
	transaction2.Description = "Online payment"

	// Display the transactions
	displayTransaction(prototypeTransaction)
	displayTransaction(transaction1)
	displayTransaction(transaction2)
}

func displayTransaction(transaction *FinancialTransaction) {
	fmt.Printf("Transaction Type: %s\n", transaction.TransactionType)
	fmt.Printf("Amount: $%.2f\n", transaction.Amount)
	fmt.Printf("Account ID: %s\n", transaction.AccountID)
	fmt.Printf("Description: %s\n", transaction.Description)
	fmt.Println()
}
