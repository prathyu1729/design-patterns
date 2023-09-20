package main

// FinancialProduct defines the high-level financial product interface.
type FinancialProduct interface {
	Apply() string
	Pay() string
}

// PaymentMethod defines the low-level payment method interface.
type PaymentMethod interface {
	ProcessPayment(amount float64) (string, error)
}
