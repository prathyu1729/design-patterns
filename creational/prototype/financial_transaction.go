package prototype

// FinancialTransaction represents a financial transaction.
type FinancialTransaction struct {
	TransactionType string
	Amount          float64
	AccountID       string
	Description     string
}

// Clone creates a deep copy of the FinancialTransaction.
func (t *FinancialTransaction) Clone() Transaction {
	return &FinancialTransaction{
		TransactionType: t.TransactionType,
		Amount:          t.Amount,
		AccountID:       t.AccountID,
		Description:     t.Description,
	}
}
