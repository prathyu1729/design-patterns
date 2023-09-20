package builder

// Transaction represents a financial transaction.
type Transaction struct {
	TransactionType string
	Amount          float64
	AccountID       string
	Description     string
}

// TransactionBuilder is the builder interface for constructing transactions.
type TransactionBuilder interface {
	SetTransactionType(transactionType string) TransactionBuilder
	SetAmount(amount float64) TransactionBuilder
	SetAccountID(accountID string) TransactionBuilder
	SetDescription(description string) TransactionBuilder
	Build() *Transaction
}
