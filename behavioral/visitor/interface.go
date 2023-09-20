package main

// PaymentMethod is the interface for payment methods (elements) that accept visitors.
type PaymentMethod interface {
	Accept(visitor PaymentCalculator)
}

// PaymentCalculator is the interface for visitors calculating payment fees.
type PaymentCalculator interface {
	VisitCreditCard(creditCard *CreditCard)
	VisitPayPal(paypal *PayPal)
}
