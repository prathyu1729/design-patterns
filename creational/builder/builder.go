package builder

// SimpleTransactionBuilder is a concrete builder for simple transactions.
type SimpleTransactionBuilder struct {
	transaction *Transaction
}

func NewSimpleTransactionBuilder() *SimpleTransactionBuilder {
	return &SimpleTransactionBuilder{
		transaction: &Transaction{},
	}
}

func (builder *SimpleTransactionBuilder) SetTransactionType(transactionType string) TransactionBuilder {
	builder.transaction.TransactionType = transactionType
	return builder
}

func (builder *SimpleTransactionBuilder) SetAmount(amount float64) TransactionBuilder {
	builder.transaction.Amount = amount
	return builder
}

func (builder *SimpleTransactionBuilder) SetAccountID(accountID string) TransactionBuilder {
	builder.transaction.AccountID = accountID
	return builder
}

func (builder *SimpleTransactionBuilder) SetDescription(description string) TransactionBuilder {
	builder.transaction.Description = description
	return builder
}

func (builder *SimpleTransactionBuilder) Build() *Transaction {
	return builder.transaction
}
