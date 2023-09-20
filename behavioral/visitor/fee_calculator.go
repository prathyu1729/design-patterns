package main

// FeeCalculator is a concrete visitor that calculates fees for payment methods.
type FeeCalculator struct {
	TotalFees float64
}

func (fc *FeeCalculator) VisitCreditCard(creditCard *CreditCard) {
	// Calculate credit card processing fee (e.g., 2%).
	fee := creditCard.Amount * 0.02
	fc.TotalFees += fee
}

func (fc *FeeCalculator) VisitPayPal(paypal *PayPal) {
	// Calculate PayPal processing fee (e.g., $0.30 + 3% of the amount).
	fee := 0.30 + paypal.Amount*0.03
	fc.TotalFees += fee
}
