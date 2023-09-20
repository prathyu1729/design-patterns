package main

// CreditCard is a concrete payment method representing a credit card payment.
type CreditCard struct {
	Amount float64
}

func (cc *CreditCard) Accept(visitor PaymentCalculator) {
	visitor.VisitCreditCard(cc)
}
